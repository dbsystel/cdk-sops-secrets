package main

import (
	"context"
	"encoding/json"
	"log"

	runtime "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"go.mozilla.org/sops/v3/decrypt"
)

type SOPSS3File struct {
	Bucket string `json:"Bucket"`
	Key    string `json:"Key"`
	Format string `json:"Format"`
}

type SOPSResourcePropertys struct {
	SecretARN         string     `json:"SecretARN"`
	S3SOPSContentFile SOPSS3File `json:"S3SOPSContentFile"`
}

type crSopsInput struct {
	RequestType           string                 `json:"RequestType"`
	LogicalResourceId     string                 `json:"LogicalResourceId"`
	PhysicalResourceId    string                 `json:"PhysicalResourceId"`
	ResourceProperties    *SOPSResourcePropertys `json:"ResourceProperties"`
	OldResourceProperties *SOPSResourcePropertys `json:"OldResourceProperties,omitempty"`
	ResourceType          string                 `json:"ResourceType"`
	RequestId             string                 `json:"RequestId"`
	StackId               string                 `json:"StackId"`
}

type crSopsOutput struct {
	PhysicalResourceId string                              `json:"PhysicalResourceId"`
	NoEcho             bool                                `json:"NoEcho"`
	Data               secretsmanager.PutSecretValueOutput `json:"Data"`
}

var awsSession = session.New()
var secrets = secretsmanager.New(awsSession)
var s3Manager = s3manager.NewDownloader(awsSession)

func getS3FileContent(file SOPSS3File) []byte {
	log.Printf("Downloading file '%s' from bucket '%s'\n", file.Key, file.Bucket)
	buf := aws.NewWriteAtBuffer([]byte{})
	resp, err := s3Manager.Download(buf, &s3.GetObjectInput{
		Bucket: &file.Bucket,
		Key:    &file.Key,
	})
	if err != nil {
		log.Fatalf("S3 download error:\n%v\n", err)
	}
	log.Printf("Downloaded %d bytes", resp)
	return buf.Bytes()
}

func decryptSopsFileContent(content []byte, file SOPSS3File) []byte {
	log.Printf("Decrypting content with format %s\n", file.Format)
	resp, err := decrypt.Data(content, file.Format)
	if err != nil {
		log.Fatalf("Decryption error:\n%v\n", err)
	}
	log.Println("decrypted")
	return resp
}

func updateSecret(secretArn string, secretContent []byte) *secretsmanager.PutSecretValueOutput {
	secretContentString := string(secretContent)
	input := &secretsmanager.PutSecretValueInput{
		SecretId:     &secretArn,
		SecretString: &secretContentString,
	}
	secretResp, secretErr := secrets.PutSecretValue(input)
	if secretErr != nil {
		log.Fatalf("Failed to store secret value:\n%v\n", secretErr)
	}
	log.Printf("Succesfully stored secret:\n%v\n", secretResp)
	return secretResp
}

func handleRequest(ctx context.Context, event crSopsInput) crSopsOutput {
	// event
	eventJson, _ := json.MarshalIndent(event, "", "  ")
	log.Printf("Function invoked with:\n %s", eventJson)

	sopsFile := &event.ResourceProperties.S3SOPSContentFile

	// This is where the magic happens
	ecnryptedContent := getS3FileContent(*sopsFile)
	decryptedContent := decryptSopsFileContent(ecnryptedContent, *sopsFile)
	updateSecretResp := updateSecret(event.ResourceProperties.SecretARN, decryptedContent)

	log.Println("Successfully finished, returning")
	return crSopsOutput{
		PhysicalResourceId: *updateSecretResp.ARN,
		NoEcho:             false,
		Data:               *updateSecretResp,
	}
}

func main() {
	runtime.Start(handleRequest)
}
