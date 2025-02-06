package main

import (
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/aws/aws-lambda-go/cfn"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/markussiebert/cdk-sops-secrets/internal/client"
	"github.com/stretchr/testify/assert"
)

type MockAwsClient struct {
	snapsFileName string
	t             *testing.T
}

func (m *MockAwsClient) S3GetObject(file client.SopsS3File) ([]byte, error) {
	snaps.WithConfig(snaps.Filename(m.snapsFileName)).MatchSnapshot(m.t, file)
	localFile := "../" + file.Key
	content, err := os.ReadFile(localFile)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (m *MockAwsClient) S3GetObjectETAG(file client.SopsS3File) (*string, error) {
	etag := "mock-etag"
	snaps.WithConfig(snaps.Filename(m.snapsFileName)).MatchSnapshot(m.t, file)
	return &etag, nil
}

func (m *MockAwsClient) SecretsManagerPutSecretValue(sopsHash string, secretArn string, secretContent *[]byte) (*secretsmanager.PutSecretValueOutput, error) {
	snaps.WithConfig(snaps.Filename(m.snapsFileName)).MatchSnapshot(m.t, map[string]interface{}{
		"sopsHash":      sopsHash,
		"secretArn":     secretArn,
		"secretContent": string(*secretContent),
	})
	arn := "mock-arn"
	return &secretsmanager.PutSecretValueOutput{
		ARN:           &arn,
		Name:          &secretArn,
		VersionStages: []string{"mock-version-stage"},
		VersionId:     &sopsHash,
	}, nil
}

func (m *MockAwsClient) SsmPutParameter(parameterName string, parameterContent *[]byte, keyId string) (*ssm.PutParameterOutput, error) {
	snaps.WithConfig(snaps.Filename(m.snapsFileName)).MatchSnapshot(m.t, map[string]interface{}{
		"parameterName":    parameterName,
		"parameterContent": string(*parameterContent),
		"keyId":            keyId,
	})
	return &ssm.PutParameterOutput{Version: 1}, nil
}

func TestHandleRequestWithClients(t *testing.T) {
	t.Logf("Running at: %s", time.Now().String()) // Forces re-run
	os.Setenv("AWS_REGION", "eu-central-1")
	os.Setenv("SOPS_AGE_KEY", "AGE-SECRET-KEY-1EFUWJ0G2XJTJFWTAM2DGMA4VCK3R05W58FSMHZP3MZQ0ZTAQEAFQC6T7T3")
	files, err := os.ReadDir("events")
	if err != nil {
		t.Fatalf("failed to read events directory: %v", err)
	}

	var tests []struct {
		name  string
		event cfn.Event
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		content, err := os.ReadFile("events/" + file.Name())
		if err != nil {
			t.Fatalf("failed to read file %s: %v", file.Name(), err)
		}

		var resourceProperties map[string]interface{}
		if err := json.Unmarshal(content, &resourceProperties); err != nil {
			t.Fatalf("failed to unmarshal JSON from file %s: %v", file.Name(), err)
		}

		tests = append(tests, struct {
			name  string
			event cfn.Event
		}{
			name: file.Name(),
			event: cfn.Event{
				RequestType:        cfn.RequestCreate,
				ResourceProperties: resourceProperties,
			},
		})
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clients := &MockAwsClient{
				snapsFileName: tt.name,
				t:             t,
			}

			physicalResourceID, data, err := HandleRequestWithClients(clients, tt.event)

			assert.NoError(t, err)
			assert.NotEmpty(t, physicalResourceID)
			assert.NotNil(t, data)
		})
	}
}
