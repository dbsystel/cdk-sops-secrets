package event

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"log"
	"log/slog"
	"os"
	"strings"

	_ "embed"

	"github.com/aws/aws-lambda-go/cfn"
	jsonSchemaGen "github.com/invopop/jsonschema"
	jsonSchemaValidate "github.com/kaptinlin/jsonschema"
	"github.com/markussiebert/cdk-sops-secrets/internal/client"
	"github.com/markussiebert/cdk-sops-secrets/internal/sops"
)

//go:embed config-schema.json
var configSchema []byte

type SopsS3File struct {
	Bucket string `json:"Bucket"`
	Key    string `json:"Key"`
}

func (s *SopsS3File) IsEmpty() bool {
	return s == nil || (s.Bucket == "" && s.Key == "")
}

type SopsInline struct {
	Content string `json:"Content"`
	Hash    string `json:"Hash"`
}

func (s *SopsInline) IsEmpty() bool {
	return s == nil || (s.Content == "" && s.Hash == "")
}

type ResourceType string

const (
	// Secret, flattened and stringified
	SECRET ResourceType = "SECRET"
	// Secret, just the raw value
	SECRET_BINARY ResourceType = "SECRET_BINARY"
	// The JSON object is flattened into multiple parameters
	PARAMETER_MULTI ResourceType = "PARAMETER_MULTI"
	// The raw value is stored as Parameter
	PARAMETER ResourceType = "PARAMETER"
)

type SopsSyncResourcePropertys struct {
	ResourceType     ResourceType `json:"ResourceType" jsonschema:"enum=SECRET,enum=SECRET_BINARY,enum=PARAMETER_MULTI,enum=PARAMETER"`
	Format           sops.Format  `json:"Format" jsonschema:"enum=json,enum=yaml,enum=dotenv,enum=binary"`
	Target           string       `json:"Target"`
	EncryptionKey    *string      `json:"EncryptionKey,omitempty"`
	SopsS3File       *SopsS3File  `json:"SopsS3File,omitempty"`
	SopsInline       *SopsInline  `json:"SopsInline,omitempty"`
	FlattenSeparator *string      `json:"FlattenSeparator,omitempty"`
	// ServiceToken is the ARN of the service token that was passed to the custom resource
	// Populated by the CloudFormation
	ServiceToken *string `json:"ServiceToken,omitempty"`
}

func FromCfnEvent(event cfn.Event) (*SopsSyncResourcePropertys, error) {
	logger := slog.With("Package", "event", "Function", "FromCfnEvent")
	compiler := jsonSchemaValidate.NewCompiler()
	schema, err := compiler.Compile(configSchema)
	if err != nil {
		logger.Error("Failed to compile schema", "Error", err)
	}

	slog.Debug("FromCfnEvent", "Event", event)

	result := schema.Validate(event.ResourceProperties)
	if !result.IsValid() {
		// Sort the details by evaluationPath
		logger.Info("Invalid ResourceProperties", "EvaluationResult", result)
		return nil, fmt.Errorf("invalid resource properties (more details in log):\n%v", event.ResourceProperties)
	}

	var props SopsSyncResourcePropertys
	data, err := json.Marshal(event.ResourceProperties)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal resource properties: %v", err)
	}

	err = json.Unmarshal(data, &props)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal resource properties: %v", err)
	}

	return &props, nil
}

func (p *SopsSyncResourcePropertys) GeneratePhysicalResourceId() string {
	return fmt.Sprintf("%s:%s:%s", "arn:custom:sopssync:", p.ResourceType, p.Target)
}

func (p *SopsSyncResourcePropertys) sopsInlineToSopsEncryptedSecret() (*sops.EncryptedSopsSecret, error) {
	content, contentErr := base64.StdEncoding.DecodeString(p.SopsInline.Content)
	if contentErr != nil {
		return nil, contentErr
	}
	return sops.CreateEncryptedSopsSecret(content, p.Format, p.SopsInline.Hash)
}

func (p *SopsSyncResourcePropertys) sopsS3FileToSopsEncryptedSecret(clients client.AwsClient) (*sops.EncryptedSopsSecret, error) {
	s3File := client.SopsS3File{
		Bucket: p.SopsS3File.Bucket,
		Key:    p.SopsS3File.Key,
	}
	s3Content, s3ContentErr := clients.S3GetObject(s3File)
	if s3ContentErr != nil {
		return nil, s3ContentErr
	}

	s3Etag, s3EtagErr := clients.S3GetObjectETAG(s3File)
	if s3EtagErr != nil {
		return nil, s3EtagErr
	}

	return sops.CreateEncryptedSopsSecret(s3Content, p.Format, *s3Etag)
}

func (p *SopsSyncResourcePropertys) GetEncryptedSopsSecret(client client.AwsClient) (*sops.EncryptedSopsSecret, error) {
	if !p.SopsInline.IsEmpty() && !p.SopsS3File.IsEmpty() {
		return nil, fmt.Errorf("both inline and S3 secret content found")
	}
	if !p.SopsInline.IsEmpty() {
		return p.sopsInlineToSopsEncryptedSecret()
	}
	if !p.SopsS3File.IsEmpty() {
		return p.sopsS3FileToSopsEncryptedSecret(client)
	}
	return nil, fmt.Errorf("no secret content found")
}

func GenerateSchema() {
	schema := jsonSchemaGen.Reflect(&SopsSyncResourcePropertys{})

	parts := strings.Split(schema.Ref, "/")
	resourceName := parts[len(parts)-1]
	schema.Properties = schema.Definitions[resourceName].Properties
	schema.AdditionalProperties = schema.Definitions[resourceName].AdditionalProperties
	schema.Type = schema.Definitions[resourceName].Type
	schema.Required = schema.Definitions[resourceName].Required
	delete(schema.Definitions, resourceName)
	schema.Ref = ""
	schema.ID = "SopsSyncResourcePropertys"
	schemaJSON, err := schema.MarshalJSON()
	if err != nil {
		log.Fatalf("Error marshaling JSON schema: %v", err)
	}

	var prettyJSON bytes.Buffer

	err = json.Indent(&prettyJSON, schemaJSON, "", "  ")
	if err != nil {
		log.Fatalf("Error indenting JSON schema: %v", err)
	}

	schemaJSON = prettyJSON.Bytes()
	if err != nil {
		log.Fatalf("Error generating JSON schema: %v", err)
	}

	file, err := os.Create("config-schema.json")
	if err != nil {
		log.Fatalf("Error creating schema file: %v", err)
	}
	defer file.Close()

	_, err = file.Write(schemaJSON)
	if err != nil {
		log.Fatalf("Error writing JSON schema to file: %v", err)
	}

	log.Printf("Generated JSON schema")
}
