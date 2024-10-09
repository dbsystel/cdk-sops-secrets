package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"reflect"
	"regexp"
	"sort"
	"strconv"
	"strings"

	runtime "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3iface"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/service/s3/s3manager/s3manageriface"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/secretsmanager/secretsmanageriface"
	"github.com/aws/aws-sdk-go/service/ssm"
	"github.com/aws/aws-sdk-go/service/ssm/ssmiface"
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
	SecretARN          string     `json:"SecretARN,omitempty"`
	ParameterName      string     `json:"ParameterName,omitempty"`
	EncryptionKey      string     `json:"EncryptionKey,omitempty"`
	SopsS3File         SopsS3File `json:"SopsS3File,omitempty"`
	SopsInline         SopsInline `json:"SopsInline,omitempty"`
	Format             string     `json:"Format"`
	ConvertToJSON      string     `json:"ConvertToJSON,omitempty"`
	Flatten            string     `json:"Flatten,omitempty"`
	FlattenSeparator   string     `json:"FlattenSeparator,omitempty"`
	ParameterKeyPrefix string     `json:"ParameterKeyPrefix,omitempty"`
	StringifyValues    string     `json:"StringifyValues,omitempty"`
	CreationType       string     `json:"CreationType,omitempty"`
	ResourceType       string     `json:"ResourceType,omitempty"`
}

type AWS struct {
	secretsmanager secretsmanageriface.SecretsManagerAPI
	ssm            ssmiface.SSMAPI
	s3Downloader   s3manageriface.DownloaderAPI
	s3Api          s3iface.S3API
}

func (a AWS) getS3FileContent(file SopsS3File) (data []byte, err error) {
	log.Printf("Downloading file '%s' from bucket '%s'\n", file.Key, file.Bucket)
	buf := aws.NewWriteAtBuffer([]byte{})
	resp, err := a.s3Downloader.Download(buf, &s3.GetObjectInput{
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

func (a AWS) updateSSMParameter(parameterName string, parameterContent []byte, keyId string) (response *ssm.PutParameterOutput, err error) {
	parameterContentString := string(parameterContent)
	input := &ssm.PutParameterInput{
		Name:      &parameterName,
		Value:     &parameterContentString,
		Type:      aws.String("SecureString"),
		Overwrite: aws.Bool(true),
		KeyId:     &keyId,
	}
	paramResp, paramErr := a.ssm.PutParameter(input)
	if paramErr != nil {
		return nil, errors.New(fmt.Sprintf("Failed to store parameter value:\nparameterName: %s\n%v\n", parameterName, paramErr))
	}
	log.Printf("Successfully stored parameter:\n%v\n", paramResp)
	return paramResp, nil
}

func generatePhysicalResourceId(input string) string {
	re := regexp.MustCompile(`(^arn:.*:secretsmanager:|ssm:)(.*)`)
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

		tempArn := generatePhysicalResourceId(resourceProperties.SecretARN + resourceProperties.ParameterName)

		sopsFile := resourceProperties.SopsS3File
		sopsInline := resourceProperties.SopsInline

		var encryptedContent []byte
		var sopsHash string

		// If SopsS3File is provided, we have to download this file
		if sopsFile != (SopsS3File{}) {
			attr, err := a.s3Api.GetObjectAttributes(&s3.GetObjectAttributesInput{
				Bucket: &sopsFile.Bucket,
				Key:    &sopsFile.Key,
			})
			encryptedContent, err = a.getS3FileContent(sopsFile)
			if err != nil {
				return tempArn, nil, err
			}
			if attr.Checksum == nil {
				return tempArn, nil, errors.New("No checksum found in S3 object")
			}
			if attr.Checksum.ChecksumCRC32 != nil {
				sopsHash = *attr.Checksum.ChecksumCRC32
			} else if attr.Checksum.ChecksumCRC32C != nil {
				sopsHash = *attr.Checksum.ChecksumCRC32C
			} else if attr.Checksum.ChecksumSHA1 != nil {
				sopsHash = *attr.Checksum.ChecksumSHA1
			} else if attr.Checksum.ChecksumSHA256 != nil {
				sopsHash = *attr.Checksum.ChecksumSHA256
			} else {
				return tempArn, nil, errors.New("No checksum found in S3 object")
			}
			log.Println(sopsHash)
		}

		// If SopsInline is provided, we have to base64 decode this content
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
		var decryptedInterface interface{}
		switch resourceProperties.Format {
		case "json":
			{
				err := json.Unmarshal(decryptedContent, &decryptedInterface)
				if err != nil {
					return tempArn, nil, fmt.Errorf("Failed to parse json content: %v", err)
				}
			}
		case "yaml":
			{
				err := yaml.Unmarshal(decryptedContent, &decryptedInterface)
				if err != nil {
					return tempArn, nil, fmt.Errorf("Failed to parse yaml content: %v", err)
				}
			}
		case "dotenv":
			{
				var dotEnvMap = make(map[string]string)
				dotenvLines := strings.Split(string(decryptedContent), "\n")
				for _, line := range dotenvLines {
					if line != "" && !strings.HasPrefix(line, "#") {
						parts := strings.SplitN(line, "=", 2)
						if len(parts) == 2 {
							key := strings.TrimSpace(parts[0])
							value := strings.TrimSpace(parts[1])
							dotEnvMap[key] = value
						}
					}
				}
				decryptedInterface = dotEnvMap
				resourceProperties.Flatten = "false"
				resourceProperties.StringifyValues = "false"
			}
		case "binary":
			{
				resourceProperties.Flatten = "false"
				resourceProperties.StringifyValues = "false"
				resourceProperties.ConvertToJSON = "false"
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
			err := flatten("", decryptedInterface, flattenedInterface, resourceProperties.FlattenSeparator)
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

		if resourceProperties.ResourceType == "SECRET" {
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
		} else if resourceProperties.ResourceType == "PARAMETER" {
			if resourceProperties.CreationType == "MULTI" && resourcePropertiesFlatten {
				log.Printf("Patching multiple string parameters")
				v := reflect.ValueOf(finalInterface)
				returnData := make(map[string]interface{})
				keys := v.MapKeys()
				keysOrder := func(i, j int) bool { return keys[i].Interface().(string) < keys[j].Interface().(string) }
				sort.Slice(keys, keysOrder)
				for _, key := range keys {
					strKey := resourceProperties.ParameterKeyPrefix + key.String()
					log.Printf("Parameter: " + strKey)
					value := v.MapIndex(key).Interface()
					strValue, ok := value.(string)
					if !ok {
						return tempArn, nil, nil
					}

					_, err := a.updateSSMParameter(strKey, []byte(strValue), resourceProperties.EncryptionKey)
					if err != nil {
						return tempArn, nil, err
					}
					// A returnData map for each parameter is not created, because it would limit the number of possible parameters unnecessarily
				}
				returnData["Prefix"] = resourceProperties.ParameterKeyPrefix
				returnData["Count"] = len(keys)
				return tempArn, returnData, nil
			} else {
				log.Printf("Patching single string parameter")
				response, err := a.updateSSMParameter(resourceProperties.ParameterName, decryptedContent, resourceProperties.EncryptionKey)
				if err != nil {
					return tempArn, nil, err
				}
				returnData := make(map[string]interface{})
				returnData["ParameterName"] = resourceProperties.ParameterName
				returnData["Version"] = response.Version
				returnData["Tier"] = response.Tier
				return tempArn, returnData, nil
			}
		} else {
			// Should never happen ...
			return tempArn, nil, errors.New("Neither SecretARN nor ParameterName is provided")
		}
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
		ssm:            ssm.New(awsSession),
		s3Downloader:   s3manager.NewDownloader(awsSession),
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

func flatten(parentkey string, input any, output map[string]interface{}, separator string) error {
	switch child := input.(type) {
	case map[string]interface{}:
		{
			for k, v := range child {
				if parentkey == "" {
					flatten(k, v, output, separator)
				} else {
					flatten(fmt.Sprintf("%s%s%s", parentkey, separator, k), v, output, separator)
				}
			}
		}
	case []interface{}:
		{
			for i, v := range child {
				if parentkey == "" {
					flatten(fmt.Sprintf("[%d]", i), v, output, separator)
				} else {
					flatten(fmt.Sprintf("%s[%d]", parentkey, i), v, output, separator)
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
