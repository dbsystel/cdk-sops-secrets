package main

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/aws/aws-sdk-go/service/s3/s3manager/s3manageriface"
	"github.com/aws/aws-sdk-go/service/secretsmanager"
	"github.com/aws/aws-sdk-go/service/secretsmanager/secretsmanageriface"
)

type SecretsManagerMockClient struct {
	secretsmanageriface.SecretsManagerAPI
}

func (m *SecretsManagerMockClient) PutSecretValue(input *secretsmanager.PutSecretValueInput) (*secretsmanager.PutSecretValueOutput, error) {
	versionId := fmt.Sprintf("%x", sha256.Sum256([]byte(*input.SecretString)))

	name := fmt.Sprintf("%x", sha256.Sum256([]byte(*input.SecretId)))

	return &secretsmanager.PutSecretValueOutput{
		ARN:           input.SecretId,
		Name:          &name,
		VersionStages: []*string{&name},
		VersionId:     &versionId,
	}, nil
}

type S3ManagerMockClient struct {
	s3manageriface.DownloaderAPI
}

func (d S3ManagerMockClient) Download(w io.WriterAt, input *s3.GetObjectInput, options ...func(*s3manager.Downloader)) (n int64, err error) {
	dat, err := os.ReadFile(*input.Key)
	check(err)
	_, err = w.WriteAt(dat, int64(0))
	check(err)
	return int64(len(dat)), err
}
