package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"

	runtime "github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/service/s3/s3manager/s3manageriface"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/secretsmanager/secretsmanageriface"
	"go.mozilla.org/sops/v3/decrypt"

	"github.com/aws/aws-lambda-go/cfn"
)

type SOPSS3File struct {
	Bucket string `json:"Bucket"`
	Key    string `json:"Key"`
	Format string `json:"Format"`
}

type SOPSResourcePropertys struct {
	SecretARN         string     `json:"SecretARN"`
	S3SOPSContentFile SOPSS3File `json:"S3SOPSContentFile"`
	SOPSAgeKey        string     `json:"SOPSAgeKey,omitempty"`
}

type AWS struct {
	secretsmanager secretsmanageriface.SecretsManagerAPI
	s3downlaoder   s3manageriface.DownloaderAPI
}

func (a AWS) getS3FileContent(file SOPSS3File) (data []byte, err error) {
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

func decryptSopsFileContent(content []byte, file SOPSS3File) (data []byte, err error) {
	log.Printf("Decrypting content with format %s\n", file.Format)
	resp, err := decrypt.Data(content, file.Format)
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

		resourceProperties := SOPSResourcePropertys{}
		if err := json.Unmarshal(jsonResourceProps, &resourceProperties); err != nil {
			return "", nil, err
		}

		// Enable AGE Support
		if resourceProperties.SOPSAgeKey != "" {
			os.Setenv("SOPS_AGE_KEY", resourceProperties.SOPSAgeKey)
		}

		sopsFile := resourceProperties.S3SOPSContentFile

		// This is where the magic happens
		ecnryptedContent, err := a.getS3FileContent(sopsFile)
		if err != nil {
			return "", nil, err
		}
		decryptedContent, err := decryptSopsFileContent(ecnryptedContent, sopsFile)
		if err != nil {
			return "", nil, err
		}
		updateSecretResp, err := a.updateSecret(resourceProperties.SecretARN, decryptedContent)
		if err != nil {
			return "", nil, err
		}

		log.Println("Successfully finished, returning")
		returnData := make(map[string]interface{})

		returnData["ARN"] = updateSecretResp.ARN
		returnData["Name"] = updateSecretResp.Name
		returnData["VersionStages"] = updateSecretResp.VersionStages
		returnData["VersionId"] = updateSecretResp.VersionId

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
