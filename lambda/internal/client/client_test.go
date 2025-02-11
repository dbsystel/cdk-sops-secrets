package client

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"testing"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/aws/smithy-go/ptr"
	"github.com/stretchr/testify/assert"
)

func TestS3GetObject(t *testing.T) {
	tests := []struct {
		name           string
		mockClient     *MockAwsClient
		expectedData   []byte
		expectedErrMsg string
	}{
		{
			name: "successful get object",
			mockClient: &MockAwsClient{
				t:             t,
				snapsFileName: "s3_get_object",
				GetObjectRet: &s3.GetObjectOutput{
					Body: io.NopCloser(bytes.NewReader([]byte("mock content"))),
				},
				ReturnError: nil,
			},
			expectedData: []byte("mock content"),
		},
		{
			name: "get object error",
			mockClient: &MockAwsClient{
				t:             t,
				snapsFileName: "s3_get_object_error",
				GetObjectRet:  nil,
				ReturnError:   fmt.Errorf("mock error"),
			},
			expectedErrMsg: "S3 get object error:\nmock error",
		},
		{
			name: "reader body error",
			mockClient: &MockAwsClient{
				t:             t,
				snapsFileName: "s3_get_object_reader_error",
				GetObjectRet: &s3.GetObjectOutput{
					Body: io.NopCloser(&mockErrorReader{}),
				},
				ReturnError: nil,
			},
			expectedErrMsg: "read buffer error:\nmock read error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				ctx: context.Background(),
				s3:  tt.mockClient,
			}

			data, err := client.S3GetObject(SopsS3File{Bucket: "test-bucket", Key: "test-key"})
			if tt.expectedErrMsg != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErrMsg)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedData, data)
			}
		})
	}
}

func TestS3GetObjectETAG(t *testing.T) {
	tests := []struct {
		name           string
		mockClient     *MockAwsClient
		expectedETAG   *string
		expectedErrMsg string
	}{
		{
			name: "successful get object etag",
			mockClient: &MockAwsClient{
				t:             t,
				snapsFileName: "s3_get_object_etag",
				GetObjectAttributesRet: &s3.GetObjectAttributesOutput{
					ETag: aws.String("mock-etag"),
				},
				ReturnError: nil,
			},
			expectedETAG: aws.String("mock-etag"),
		},
		{
			name: "get object attributes error",
			mockClient: &MockAwsClient{
				t:                      t,
				snapsFileName:          "s3_get_object_attributes_error",
				GetObjectAttributesRet: nil,
				ReturnError:            fmt.Errorf("mock error"),
			},
			expectedErrMsg: "error while getting S3 object attributes:\nmock error",
		},
		{
			name: "etag not found",
			mockClient: &MockAwsClient{
				t:             t,
				snapsFileName: "s3_get_object_etag_not_found",
				GetObjectAttributesRet: &s3.GetObjectAttributesOutput{
					ETag: nil,
				},
				ReturnError: nil,
			},
			expectedErrMsg: "no ETag checksum found in S3 object:\n<nil>",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				ctx: context.Background(),
				s3:  tt.mockClient,
			}

			etag, err := client.S3GetObjectETAG(SopsS3File{Bucket: "test-bucket", Key: "test-key"})
			if tt.expectedErrMsg != "" {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.expectedErrMsg)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedETAG, etag)
			}
		})
	}
}

func TestSecretsManagerPutSecretValue(t *testing.T) {
	tests := []struct {
		name           string
		mockClient     *MockAwsClient
		expectedARN    *string
		expectedErrMsg string
		binary         *bool
	}{
		{
			name: "successful put secret value",
			mockClient: &MockAwsClient{
				t:             t,
				snapsFileName: "secrets_manager_put_secret_value",
				PutSecretValueRet: &secretsmanager.PutSecretValueOutput{
					ARN: aws.String("mock-arn"),
				},
				ReturnError: nil,
			},
			binary:      ptr.Bool(false),
			expectedARN: aws.String("mock-arn"),
		},
		{
			name: "put secret value error",
			mockClient: &MockAwsClient{
				t:                 t,
				snapsFileName:     "secrets_manager_put_secret_value_error",
				PutSecretValueRet: nil,
				ReturnError:       fmt.Errorf("mock error"),
			},
			binary:         ptr.Bool(false),
			expectedErrMsg: "failed to store secret value:\nsecretArn: test-arn\nClientRequestToken: test-hash\nmock error",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				ctx:            context.Background(),
				secretsManager: tt.mockClient,
			}

			content := []byte("content")
			resp, err := client.SecretsManagerPutSecretValue("hash", "arn", &content, tt.binary)
			assert.Equal(t, tt.mockClient.ReturnError, err)
			assert.Equal(t, tt.mockClient.PutSecretValueRet, resp)
		})
	}
}

func TestSsmPutParameter(t *testing.T) {
	tests := []struct {
		name            string
		mockClient      *MockAwsClient
		expectedVersion *int64
		expectedErrMsg  string
	}{
		{
			name: "successful update parameter",
			mockClient: &MockAwsClient{
				t:             t,
				snapsFileName: "ssm_update_update_parameter",
				PutParameterRet: &ssm.PutParameterOutput{
					Version: *aws.Int64(1),
				},
				ReturnError: nil,
			},
		},
		{
			name: "update parameter error",
			mockClient: &MockAwsClient{
				t:               t,
				snapsFileName:   "ssm_update_update_parameter_error",
				PutParameterRet: nil,
				ReturnError:     fmt.Errorf("mock error"),
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client := &Client{
				ctx: context.Background(),
				ssm: tt.mockClient,
			}
			content := []byte("content")
			resp, err := client.SsmPutParameter("parameter", &content, "key")
			assert.Equal(t, tt.mockClient.ReturnError, err)
			assert.Equal(t, tt.mockClient.PutParameterRet, resp)
		})
	}
}
