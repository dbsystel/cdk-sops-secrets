package client

import (
	"context"
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/gkampitakis/go-snaps/snaps"
)

type mockErrorReader struct{}

func (r *mockErrorReader) Read(p []byte) (n int, err error) {
	return 0, fmt.Errorf("mock read error")
}

// MockAwsClient is a mock for the real AWS client
type MockAwsClient struct {
	t                      *testing.T
	snapsFileName          string
	ReturnError            error
	GetObjectRet           *s3.GetObjectOutput
	GetObjectAttributesRet *s3.GetObjectAttributesOutput
	PutParameterRet        *ssm.PutParameterOutput
	PutSecretValueRet      *secretsmanager.PutSecretValueOutput
}

func (s *MockAwsClient) GetObject(ctx context.Context, params *s3.GetObjectInput, optFns ...func(*s3.Options)) (*s3.GetObjectOutput, error) {
	if s.ReturnError == nil {
		snaps.WithConfig(snaps.Filename(s.snapsFileName)).MatchSnapshot(s.t, params)
	}
	return s.GetObjectRet, s.ReturnError
}

func (s *MockAwsClient) GetObjectAttributes(ctx context.Context, params *s3.GetObjectAttributesInput, optFns ...func(*s3.Options)) (*s3.GetObjectAttributesOutput, error) {
	if s.ReturnError == nil {
		snaps.WithConfig(snaps.Filename(s.snapsFileName)).MatchSnapshot(s.t, params)
	}
	return s.GetObjectAttributesRet, s.ReturnError
}

func (s *MockAwsClient) PutParameter(ctx context.Context, params *ssm.PutParameterInput, optFns ...func(*ssm.Options)) (*ssm.PutParameterOutput, error) {
	if s.ReturnError == nil {
		snaps.WithConfig(snaps.Filename(s.snapsFileName)).MatchSnapshot(s.t, params)
	}
	return s.PutParameterRet, s.ReturnError
}

func (s *MockAwsClient) PutSecretValue(ctx context.Context, params *secretsmanager.PutSecretValueInput, optFns ...func(*secretsmanager.Options)) (*secretsmanager.PutSecretValueOutput, error) {
	if s.ReturnError == nil {
		snaps.WithConfig(snaps.Filename(s.snapsFileName)).MatchSnapshot(s.t, params)
	}
	return s.PutSecretValueRet, s.ReturnError
}

// MockClient is a mock for "our" AWS client
type MockClient struct {
	S3Content []byte
	S3Etag    string
	S3Err     error
	S3EtagErr error
}

func (m *MockClient) S3GetObject(file SopsS3File) ([]byte, error) {
	return m.S3Content, m.S3Err
}

func (m *MockClient) S3GetObjectETAG(file SopsS3File) (*string, error) {
	return &m.S3Etag, m.S3EtagErr
}

func (m *MockClient) SecretsManagerPutSecretValue(sopsHash string, secretArn string, secretContent *[]byte) (*secretsmanager.PutSecretValueOutput, error) {
	return nil, nil
}

func (m *MockClient) SsmPutParameter(parameterName string, parameterContent *[]byte, keyId string) (*ssm.PutParameterOutput, error) {
	return nil, nil
}
