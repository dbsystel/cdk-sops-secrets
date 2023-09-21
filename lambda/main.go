package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"

	runtime "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/service/s3/s3manager/s3manageriface"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/secretsmanager/secretsmanageriface"
	"github.com/getsops/sops/v3/decrypt"
	"gopkg.in/yaml.v3"

	"github.com/aws/aws-lambda-go/cfn"
)

type SopsS3File struct {
	Bucket string `json:"Bucket,omitempty"`
	Key    string `json:"Key,omitempty"`
}

type SopsInline struct {
	Content string `json:"Content,omitempty"`
	Hash    string `json:"Hash,omitempty"`
}
type SopsSyncResourcePropertys struct {
	SecretARN       string     `json:"SecretARN"`
	SopsS3File      SopsS3File `json:"SopsS3File,omitempty"`
	SopsInline      SopsInline `json:"SopsInline,omitempty"`
	Format          string     `json:"Format"`
	ConvertToJSON   string     `json:"ConvertToJSON,omitempty"`
	Flatten         string     `json:"Flatten,omitempty"`
	StringifyValues string     `json:"StringifyValues,omitempty"`
}

type AWS struct {
	secretsmanager secretsmanageriface.SecretsManagerAPI
	s3downlaoder   s3manageriface.DownloaderAPI
}

func (a AWS) getS3FileContent(file SopsS3File) (data []byte, err error) {
	log.Printf("Downloading file '%s' from bucket '%s'\n", file.Key, file.Bucket)
	buf := aws.NewWriteAtBuffer([]byte{})
	resp, err := a.s3downlaoder.Download(buf, &s3.GetObjectInput{
		Bucket: &file.Bucket,
		Key:    &file.Key,
	})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("S3 download error:\n%v\n", err))
	}
	log.Printf("Downloaded %d bytes", resp)
	return buf.Bytes(), nil
}

func decryptSopsFileContent(content []byte, format string) (data []byte, err error) {
	log.Printf("Decrypting content with format %s\n", format)
	resp, err := decrypt.Data(content, format)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("Decryption error:\n%v\n", err))
	}
	log.Println("Decrypted")
	return resp, nil
}

func (a AWS) updateSecret(sopsHash string, secretArn string, secretContent []byte) (data *secretsmanager.PutSecretValueOutput, err error) {
	secretContentString := string(secretContent)
	input := &secretsmanager.PutSecretValueInput{
		SecretId:           &secretArn,
		SecretString:       &secretContentString,
		ClientRequestToken: &sopsHash,
	}
	secretResp, secretErr := a.secretsmanager.PutSecretValue(input)
	if secretErr != nil {
		return nil, errors.New(fmt.Sprintf("Failed to store secret value:\nsecretArn: %s\nClientRequestToken: %s\n%v\n", secretArn, sopsHash, secretErr))
	}
	arn := generatePhysicalResourceId(*secretResp.ARN)
	secretResp.ARN = &arn
	log.Printf("Succesfully stored secret:\n%v\n", secretResp)
	return secretResp, nil
}

func generatePhysicalResourceId(input string) string {
	re := regexp.MustCompile(`(^arn:.*:secretsmanager:)(.*)`)
	return re.ReplaceAllString(input, `arn:custom:sopssync:$2`)
}

func (a AWS) syncSopsToSecretsmanager(ctx context.Context, event cfn.Event) (physicalResourceID string, data map[string]interface{}, err error) {
	// event
	// eventJson, _ := json.MarshalIndent(event, "", "  ")
	// log.Printf("Function invoked with:\n %s", eventJson)

	if event.RequestType == cfn.RequestCreate || event.RequestType == cfn.RequestUpdate {

		// some casting

		jsonResourceProps, err := json.Marshal(event.ResourceProperties)
		if err != nil {
			return "error", nil, err
		}

		resourceProperties := SopsSyncResourcePropertys{}
		if err := json.Unmarshal(jsonResourceProps, &resourceProperties); err != nil {
			return "", nil, err
		}

		tempArn := generatePhysicalResourceId(resourceProperties.SecretARN)

		sopsFile := resourceProperties.SopsS3File
		sopsInline := resourceProperties.SopsInline

		var encryptedContent []byte
		var sopsHash string

		// If SopsS3File is provided, we have to download this file
		if sopsFile != (SopsS3File{}) {
			encryptedContent, err = a.getS3FileContent(sopsFile)
			fileNameParts := strings.Split(sopsFile.Key, ".")
			fileNameParts = fileNameParts[:len(fileNameParts)-1]
			fileNameWithoutEnding := strings.Join(fileNameParts, ".")
			sopsHash = fileNameWithoutEnding
			log.Println(sopsHash)
			if err != nil {
				return tempArn, nil, err
			}
		}

		// If SopsInline is provbided, we have to base64 decode this content
		if sopsInline != (SopsInline{}) {
			encryptedContent, err = base64.StdEncoding.DecodeString(sopsInline.Content)
			sopsHash = sopsInline.Hash
			if err != nil {
				return tempArn, nil, err
			}
		}

		if encryptedContent == nil {
			return tempArn, nil, errors.New("No encrypted content found! Did you pass SopsS3File or SopsInline?")
		}
		if sopsHash == "" {
			return tempArn, nil, errors.New("No sopsHash found! Did you pass SopsS3File or SopsInline?")
		}

		decryptedContent, err := decryptSopsFileContent(encryptedContent, resourceProperties.Format)
		if err != nil {
			return tempArn, nil, err
		}
		//log.Println(string(decryptedContent))
		var decryptedInterface interface{}
		switch resourceProperties.Format {
		case "json":
			{
				err := json.Unmarshal(decryptedContent, &decryptedInterface)
				if err != nil {
					return tempArn, nil, err
				}
			}
		case "yaml":
			{
				err := yaml.Unmarshal(decryptedContent, &decryptedInterface)
				if err != nil {
					return tempArn, nil, err
				}
			}
		default:
			return "", nil, errors.New(fmt.Sprintf("Format %s not supported", resourceProperties.Format))
		}

		if resourceProperties.Flatten == "" {
			resourceProperties.Flatten = "true"
		}
		resourcePropertiesFlatten, err := strconv.ParseBool(resourceProperties.Flatten)
		if err != nil {
			return tempArn, nil, err
		}

		var finalInterface interface{}

		if resourcePropertiesFlatten {
			flattenedInterface := make(map[string]interface{})
			err := flatten("", decryptedInterface, flattenedInterface)
			if err != nil {
				return tempArn, nil, err
			}
			finalInterface = flattenedInterface
		} else {
			finalInterface = decryptedInterface
		}

		if resourceProperties.StringifyValues == "" {
			resourceProperties.StringifyValues = "true"
		}
		resourcePropertiesStringifyValues, err := strconv.ParseBool(resourceProperties.StringifyValues)
		if err != nil {
			return tempArn, nil, err
		}
		if resourcePropertiesStringifyValues {
			finalInterface, _, err = stringifyValues(finalInterface)
			if err != nil {
				return tempArn, nil, err
			}
		}

		if resourceProperties.ConvertToJSON == "" {
			resourceProperties.ConvertToJSON = "true"
		}
		resourcePropertieConvertToJSON, err := strconv.ParseBool(resourceProperties.ConvertToJSON)
		if err != nil {
			return tempArn, nil, err
		}
		if resourcePropertieConvertToJSON || resourceProperties.Format == "json" {
			decryptedContent, err = toJSON(finalInterface)
			if err != nil {
				return tempArn, nil, err
			}
		} else if resourceProperties.Format == "yaml" {
			decryptedContent, err = toYAML(finalInterface)
			if err != nil {
				return tempArn, nil, err
			}
		}
		// Write the secret
		updateSecretResp, err := a.updateSecret(sopsHash, resourceProperties.SecretARN, decryptedContent)
		if err != nil {
			return tempArn, nil, err
		}

		returnData := make(map[string]interface{})

		returnData["ARN"] = *updateSecretResp.ARN
		returnData["Name"] = *updateSecretResp.Name
		returnData["VersionStages"] = updateSecretResp.VersionStages
		returnData["VersionId"] = *updateSecretResp.VersionId

		return *updateSecretResp.ARN, returnData, nil
	} else if event.RequestType == cfn.RequestDelete {
		return "", nil, nil
	} else {
		// Should never happen ...
		return "", nil, errors.New(fmt.Sprintf("RequestType '%s' not supported", event.RequestType))
	}
}

func handleRequest(ctx context.Context, event cfn.Event) (physicalResourceID string, data map[string]interface{}, err error) {
	awsSession := session.New()

	a := &AWS{
		secretsmanager: secretsmanager.New(awsSession),
		s3downlaoder:   s3manager.NewDownloader(awsSession),
	}

	return a.syncSopsToSecretsmanager(ctx, event)
}

func main() {
	runtime.Start(cfn.LambdaWrap(handleRequest))
}

func fromYAML(in []byte) (interface{}, error) {
	var ret interface{}
	err := yaml.Unmarshal(in, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func fromJSON(in []byte) (interface{}, error) {
	var ret interface{}
	err := json.Unmarshal(in, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func toJSON(in any) ([]byte, error) {
	ret, err := json.MarshalIndent(in, "", "  ")
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func toYAML(in any) ([]byte, error) {
	ret, err := yaml.Marshal(in)
	if err != nil {
		return nil, err
	}
	return ret, nil
}

func stringifyValues(input any) (interface{}, string, error) {
	switch child := input.(type) {
	case map[string]interface{}:
		{
			output := make(map[string]interface{})
			for k, v := range child {
				object, val, err := stringifyValues(v)
				if err != nil {
					return nil, "", err
				}
				if object != nil {
					output[k] = object
				} else {
					output[k] = val
				}
			}
			return output, "", nil
		}
	case []interface{}:
		{
			output := []interface{}{}
			for _, v := range child {
				object, val, err := stringifyValues(v)
				if err != nil {
					return nil, "", err
				}
				if object != nil {
					output = append(output, object)
				} else {
					output = append(output, val)
				}
			}
			return output, "", nil
		}
	default:
		{
			return nil, fmt.Sprint(input), nil
		}
	}
}

func flatten(parentkey string, input any, output map[string]interface{}) error {

	switch child := input.(type) {
	case map[string]interface{}:
		{
			for k, v := range child {
				if parentkey == "" {
					flatten(k, v, output)
				} else {
					flatten(fmt.Sprintf("%s%s%s", parentkey, ".", k), v, output)
				}

			}
		}
	case []interface{}:
		{
			for i, v := range child {
				if parentkey == "" {
					flatten(fmt.Sprintf("[%d]", i), v, output)
				} else {
					flatten(fmt.Sprintf("%s[%d]", parentkey, i), v, output)
				}
			}
		}
	default:
		{
			output[parentkey] = input
		}
	}

	return nil
}

func toSopsSyncResourcePropertys(input *map[string]interface{}) (*SopsSyncResourcePropertys, error) {
	jsonResourceProps, err := json.Marshal(&input)
	if err != nil {
		return nil, err
	}

	resourceProperties := SopsSyncResourcePropertys{}
	if err := json.Unmarshal(jsonResourceProps, &resourceProperties); err != nil {
		return nil, err
	}
	return &resourceProperties, nil
}
