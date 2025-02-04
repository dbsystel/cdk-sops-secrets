package client

import (
	"bytes"
	"context"
	"fmt"
	"log"
	"log/slog"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	s3Types "github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	ssmTypes "github.com/aws/aws-sdk-go-v2/service/ssm/types"
)

type S3Client interface {
	GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error)
	GetObjectAttributes(ctx context.Context, params *s3.GetObjectAttributesInput, optFns ...func(*s3.Options)) (*s3.GetObjectAttributesOutput, error)
}

type SecretsManagerClient interface {
	PutSecretValue(ctx context.Context, params *secretsmanager.PutSecretValueInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.PutSecretValueOutput, error)
}

type SsmClient interface {
	PutParameter(ctx context.Context, params *ssm.PutParameterInput, optFns ...func(*ssm.Options)) (*ssm.PutParameterOutput, error)
}

type AwsClient interface {
	S3GetObject(file SopsS3File) (data []byte, err error)
	S3GetObjectETAG(file SopsS3File) (*string, error)
	SecretsManagerPutSecretValue(sopsHash string, secretArn string, secretContent *[]byte) (data *secretsmanager.PutSecretValueOutput, err error)
	SsmPutParameter(parameterName string, parameterContent *[]byte, keyId string) (response *ssm.PutParameterOutput, err error)
}

type Client struct {
	ctx            context.Context
	s3             S3Client
	secretsManager SecretsManagerClient
	ssm            SsmClient
}

func CreateAwsClients(context context.Context) AwsClient {
	cfg, err := config.LoadDefaultConfig(context)
	if err != nil {
		log.Fatalf("unable to load SDK config, %v", err)
	}
	return &Client{
		ctx:            context,
		s3:             s3.NewFromConfig(cfg),
		secretsManager: secretsmanager.NewFromConfig(cfg),
		ssm:            ssm.NewFromConfig(cfg),
	}
}

type SopsS3File struct {
	Bucket string `json:"Bucket,omitempty"`
	Key    string `json:"Key,omitempty"`
}

func (c *Client) S3GetObject(file SopsS3File) (data []byte, err error) {
	logger := slog.With("Package", "client", "Function", "S3GetObject")
	logger.Info("Downloading file", "Bucket", file.Bucket, "Key", file.Key)

	resp, err := c.s3.GetObject(c.ctx, &s3.GetObjectInput{
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
	logger.Info("Downloaded file", "Size", buf.Len())
	return buf.Bytes(), nil
}

func (c *Client) S3GetObjectETAG(file SopsS3File) (*string, error) {
	attr, err := c.s3.GetObjectAttributes(c.ctx, &s3.GetObjectAttributesInput{
		Bucket: &file.Bucket,
		Key:    &file.Key,
		ObjectAttributes: []s3Types.ObjectAttributes{
			"ETag",
		},
	})

	if err != nil {
		return nil, fmt.Errorf("error while getting S3 object attributes:\n%v", err)
	}

	if attr.ETag == nil {
		return nil, fmt.Errorf("no ETag checksum found in S3 object:\n%v", err)
	}

	return attr.ETag, nil
}

func (c *Client) SecretsManagerPutSecretValue(sopsHash string, secretArn string, secretContent *[]byte) (data *secretsmanager.PutSecretValueOutput, err error) {
	secretContentString := string(*secretContent)
	input := &secretsmanager.PutSecretValueInput{
		SecretId:           &secretArn,
		SecretString:       &secretContentString,
		ClientRequestToken: &sopsHash,
	}
	return c.secretsManager.PutSecretValue(c.ctx, input)
}

func (c *Client) SsmPutParameter(parameterName string, parameterContent *[]byte, keyId string) (response *ssm.PutParameterOutput, err error) {
	parameterContentString := string(*parameterContent)
	input := &ssm.PutParameterInput{
		Name:      &parameterName,
		Value:     &parameterContentString,
		Type:      ssmTypes.ParameterTypeSecureString,
		Overwrite: aws.Bool(true),
		KeyId:     &keyId,
	}
	return c.ssm.PutParameter(c.ctx, input)
}
