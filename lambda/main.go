package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"

	runtime "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/service/s3/s3manager/s3manageriface"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/secretsmanager/secretsmanageriface"
	"go.mozilla.org/sops/v3/decrypt"
	"gopkg.in/yaml.v3"

	"github.com/aws/aws-lambda-go/cfn"
)

type SopsS3File struct {
	Bucket string `json:"Bucket"`
	Key    string `json:"Key"`
}

type SopsSyncResourcePropertys struct {
	SecretARN       string     `json:"SecretARN"`
	SopsS3File      SopsS3File `json:"SopsS3File"`
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

func (a AWS) updateSecret(secretArn string, secretContent []byte) (data *secretsmanager.PutSecretValueOutput, err error) {
	secretContentString := string(secretContent)
	input := &secretsmanager.PutSecretValueInput{
		SecretId:     &secretArn,
		SecretString: &secretContentString,
	}
	secretResp, secretErr := a.secretsmanager.PutSecretValue(input)
	if secretErr != nil {
		return nil, errors.New(fmt.Sprintf("Failed to store secret value:\n%v\n", secretErr))
	}
	re := regexp.MustCompile(`(^arn:.*:secretsmanager:)(.*)`)
	arn := re.ReplaceAllString(*secretResp.ARN, `arn:custom:sopssync:$2`)
	secretResp.ARN = &arn
	log.Printf("Succesfully stored secret:\n%v\n", secretResp)
	return secretResp, nil
}

func (a AWS) syncSopsToSecretsmanager(ctx context.Context, event cfn.Event) (physicalResourceID string, data map[string]interface{}, err error) {
	// event
	// eventJson, _ := json.MarshalIndent(event, "", "  ")
	// log.Printf("Function invoked with:\n %s", eventJson)

	if event.RequestType == cfn.RequestCreate || event.RequestType == cfn.RequestUpdate {

		// some casting

		jsonResourceProps, err := json.Marshal(event.ResourceProperties)
		if err != nil {
			return "", nil, err
		}

		resourceProperties := SopsSyncResourcePropertys{}
		if err := json.Unmarshal(jsonResourceProps, &resourceProperties); err != nil {
			return "", nil, err
		}

		sopsFile := resourceProperties.SopsS3File

		// This is where the magic happens
		ecnryptedContent, err := a.getS3FileContent(sopsFile)
		if err != nil {
			return "", nil, err
		}
		decryptedContent, err := decryptSopsFileContent(ecnryptedContent, resourceProperties.Format)
		if err != nil {
			return "", nil, err
		}
		log.Println(string(decryptedContent))
		var decryptedInterface interface{}
		switch resourceProperties.Format {
		case "json":
			{
				err := json.Unmarshal(decryptedContent, &decryptedInterface)
				if err != nil {
					return "", nil, err
				}
			}
		case "yaml":
			{
				err := yaml.Unmarshal(decryptedContent, &decryptedInterface)
				if err != nil {
					return "", nil, err
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
			return "", nil, err
		}

		var finalInterface interface{}

		if resourcePropertiesFlatten {
			flattenedInterface := make(map[string]interface{})
			err := flatten("", decryptedInterface, flattenedInterface)
			if err != nil {
				return "", nil, err
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
			return "", nil, err
		}
		if resourcePropertiesStringifyValues {
			finalInterface, _, err = stringifyValues(finalInterface)
			if err != nil {
				return "", nil, err
			}
		}

		if resourceProperties.ConvertToJSON == "" {
			resourceProperties.ConvertToJSON = "true"
		}
		resourcePropertieConvertToJSON, err := strconv.ParseBool(resourceProperties.ConvertToJSON)
		if err != nil {
			return "", nil, err
		}
		if resourcePropertieConvertToJSON || resourceProperties.Format == "json" {
			decryptedContent, err = toJSON(finalInterface)
			if err != nil {
				return "", nil, err
			}
		} else if resourceProperties.Format == "yaml" {
			decryptedContent, err = toYAML(finalInterface)
			if err != nil {
				return "", nil, err
			}
		}
		// Write the secret
		updateSecretResp, err := a.updateSecret(resourceProperties.SecretARN, decryptedContent)
		if err != nil {
			return "", nil, err
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
