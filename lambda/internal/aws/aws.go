package aws

import (
	"bytes"
	"context"
	"fmt"
	"log"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	ssmTypes "github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

type Client struct {
	ctx            context.Context
	S3             *s3.Client
	SecretsManager *secretsmanager.Client
	Ssm            *ssm.Client
}

func CreateAwsClients() *Client {
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRegion("us-west-2"))
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	return &Client{
		S3:             s3.NewFromConfig(cfg),
		SecretsManager: secretsmanager.NewFromConfig(cfg),
		Ssm:            ssm.NewFromConfig(cfg),
	}
}

type S3GetObjectInput struct {
	Bucket string `json:"Bucket,omitempty"`
	Key    string `json:"Key,omitempty"`
}

func (c *Client) S3GetObject(file S3GetObjectInput) (data []byte, err error) {
	log.Printf("Downloading file '%s' from bucket '%s'\n", file.Key, file.Bucket)
	resp, err := c.S3.GetObject(context.TODO(), &s3.GetObjectInput{
		Bucket: &file.Bucket,
		Key:    &file.Key,
	})
	if err != nil {
		return nil, fmt.Errorf("S3 get object error:\n%v", err)
	}
	defer resp.Body.Close()

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("read buffer error:\n%v", err)
	}
	log.Printf("Downloaded %d bytes", buf.Len())
	return buf.Bytes(), nil
}

func (c *Client) SecretsManagerPutSecretValue(sopsHash string, secretArn string, secretContent []byte) (data *secretsmanager.PutSecretValueOutput, err error) {
	secretContentString := string(secretContent)
	input := &secretsmanager.PutSecretValueInput{
		SecretId:           &secretArn,
		SecretString:       &secretContentString,
		ClientRequestToken: &sopsHash,
	}
	secretResp, secretErr := c.SecretsManager.PutSecretValue(c.ctx, input)
	if secretErr != nil {
		return nil, fmt.Errorf("failed to store secret value:\nsecretArn: %s\nClientRequestToken: %s\n%v", secretArn, sopsHash, secretErr)
	}
	return secretResp, nil
}

func (c *Client) SsmUpdateUpdateParameter(parameterName string, parameterContent []byte, keyId string) (response *ssm.PutParameterOutput, err error) {
	parameterContentString := string(parameterContent)
	input := &ssm.PutParameterInput{
		Name:      &parameterName,
		Value:     &parameterContentString,
		Type:      ssmTypes.ParameterTypeSecureString,
		Overwrite: aws.Bool(true),
		KeyId:     &keyId,
	}
	paramResp, paramErr := c.Ssm.PutParameter(c.ctx, input)
	if paramErr != nil {
		return nil, fmt.Errorf("failed to store parameter value:\nparameterName: %s\n%v", parameterName, paramErr)
	}
	log.Printf("Successfully stored parameter:\n%v\n", paramResp)
	return paramResp, nil
}
