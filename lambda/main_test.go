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

type putParameterCalls map[string]interface{}
type putSecretValueCalls map[string]interface{}
type getObjectEtagCalls map[string]client.SopsS3File
type getObjectCalls map[string]client.SopsS3File
type MockAwsClient struct {
	t              *testing.T
	putParameter   putParameterCalls
	putSecretValue putSecretValueCalls
	getObjectEtag  getObjectEtagCalls
	getObject      getObjectCalls
}

func (m *MockAwsClient) S3GetObject(file client.SopsS3File) ([]byte, error) {
	m.getObject[file.Key] = file
	localFile := "../" + file.Key
	content, err := os.ReadFile(localFile)
	if err != nil {
		return nil, err
	}
	return content, nil
}

func (m *MockAwsClient) S3GetObjectETAG(file client.SopsS3File) (*string, error) {
	etag := "mock-etag"
	m.getObjectEtag[file.Key] = file
	return &etag, nil
}

func (m *MockAwsClient) SecretsManagerPutSecretValue(sopsHash string, secretArn string, secretContent *[]byte, binary *bool) (*secretsmanager.PutSecretValueOutput, error) {
	m.putSecretValue[secretArn] = map[string]interface{}{"sopsHash": sopsHash, "secretContent": string(*secretContent), "binary": *binary}
	arn := "mock-arn"
	return &secretsmanager.PutSecretValueOutput{
		ARN:           &arn,
		Name:          &secretArn,
		VersionStages: []string{"mock-version-stage"},
		VersionId:     &sopsHash,
	}, nil
}

func (m *MockAwsClient) SsmPutParameter(parameterName string, parameterContent *[]byte, keyId string) (*ssm.PutParameterOutput, error) {
	m.putParameter[parameterName] = map[string]interface{}{"parameterContent": string(*parameterContent), "keyId": keyId}
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
				t:              t,
				putParameter:   putParameterCalls{},
				putSecretValue: putSecretValueCalls{},
				getObjectEtag:  getObjectEtagCalls{},
				getObject:      getObjectCalls{},
			}

			physicalResourceID, data, err := HandleRequestWithClients(clients, tt.event)
			snaps.WithConfig(snaps.Filename(t.Name()+"/"+tt.name)).MatchSnapshot(t, clients.getObject, clients.getObjectEtag, clients.putParameter, clients.putSecretValue)

			assert.NoError(t, err)
			assert.NotEmpty(t, physicalResourceID)
			assert.NotNil(t, data)
		})
	}
}
