//go:build !integration

package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
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
	t                     *testing.T
	putParameter          putParameterCalls
	putSecretValue        putSecretValueCalls
	getObjectEtag         getObjectEtagCalls
	getObject             getObjectCalls
	ssmGetParameterValues map[string]string
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

func (m *MockAwsClient) SsmGetParameter(parameterName string) (*string, error) {
	if m.ssmGetParameterValues == nil {
		return nil, fmt.Errorf("SsmGetParameter not expected in unit tests")
	}
	v, ok := m.ssmGetParameterValues[parameterName]
	if !ok {
		return nil, fmt.Errorf("SsmGetParameter: unexpected parameter %q", parameterName)
	}
	return &v, nil
}

func (m *MockAwsClient) SchedulerCreateOrUpdateSchedule(name string, groupName string, scheduleExpression string, topicArn string, roleArn string, message string) error {
	return nil
}

func (m *MockAwsClient) SchedulerDeleteSchedule(name string, groupName string) error {
	return nil
}

func (m *MockAwsClient) SchedulerListSchedules(groupName string) ([]string, error) {
	return nil, nil
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

func TestLoadAgeKeysFromSSM(t *testing.T) {
	makeClients := func(params map[string]string) *MockAwsClient {
		return &MockAwsClient{
			t:                     t,
			putParameter:          putParameterCalls{},
			putSecretValue:        putSecretValueCalls{},
			getObjectEtag:         getObjectEtagCalls{},
			getObject:             getObjectCalls{},
			ssmGetParameterValues: params,
		}
	}

	t.Run("noop when SOPS_AGE_KEY_PARAMS is unset", func(t *testing.T) {
		t.Setenv("SOPS_AGE_KEY_PARAMS", "")
		t.Setenv("SOPS_AGE_KEY", "static-key")

		err := loadAgeKeysFromSSM(makeClients(nil))
		assert.NoError(t, err)
		assert.Equal(t, "static-key", os.Getenv("SOPS_AGE_KEY"))
	})

	t.Run("single SSM parameter replaces SOPS_AGE_KEY when no static key", func(t *testing.T) {
		t.Setenv("SOPS_AGE_KEY_PARAMS", "/my/age/key")
		t.Setenv("SOPS_AGE_KEY", "")
		staticAgeKey = ""

		err := loadAgeKeysFromSSM(makeClients(map[string]string{
			"/my/age/key": "AGE-SECRET-KEY-FROM-SSM",
		}))
		assert.NoError(t, err)
		assert.Equal(t, "AGE-SECRET-KEY-FROM-SSM", os.Getenv("SOPS_AGE_KEY"))
	})

	t.Run("multiple SSM parameters are joined with newline", func(t *testing.T) {
		t.Setenv("SOPS_AGE_KEY_PARAMS", "/key/one\n/key/two")
		t.Setenv("SOPS_AGE_KEY", "")
		staticAgeKey = ""

		err := loadAgeKeysFromSSM(makeClients(map[string]string{
			"/key/one": "AGE-KEY-ONE",
			"/key/two": "AGE-KEY-TWO",
		}))
		assert.NoError(t, err)
		got := os.Getenv("SOPS_AGE_KEY")
		parts := strings.Split(got, "\n")
		assert.Equal(t, []string{"AGE-KEY-ONE", "AGE-KEY-TWO"}, parts)
	})

	t.Run("static cold-start key is preserved and prepended", func(t *testing.T) {
		const coldStartKey = "AGE-SECRET-KEY-COLDSTART"
		staticAgeKey = coldStartKey
		t.Setenv("SOPS_AGE_KEY_PARAMS", "/ssm/key")
		t.Setenv("SOPS_AGE_KEY", coldStartKey)

		err := loadAgeKeysFromSSM(makeClients(map[string]string{
			"/ssm/key": "AGE-SECRET-KEY-FROM-SSM",
		}))
		assert.NoError(t, err)
		got := os.Getenv("SOPS_AGE_KEY")
		parts := strings.Split(got, "\n")
		assert.Equal(t, []string{coldStartKey, "AGE-SECRET-KEY-FROM-SSM"}, parts)
	})

	t.Run("warm invocation does not accumulate SSM keys", func(t *testing.T) {
		const coldStartKey = "AGE-SECRET-KEY-COLDSTART"
		staticAgeKey = coldStartKey
		t.Setenv("SOPS_AGE_KEY_PARAMS", "/ssm/key")

		clients := makeClients(map[string]string{"/ssm/key": "AGE-SECRET-KEY-FROM-SSM"})

		assert.NoError(t, loadAgeKeysFromSSM(clients))
		first := os.Getenv("SOPS_AGE_KEY")

		assert.NoError(t, loadAgeKeysFromSSM(clients))
		second := os.Getenv("SOPS_AGE_KEY")

		assert.Equal(t, first, second, "warm invocation must not duplicate keys")
		parts := strings.Split(second, "\n")
		assert.Equal(t, []string{coldStartKey, "AGE-SECRET-KEY-FROM-SSM"}, parts)
	})

	t.Run("blank lines in SOPS_AGE_KEY_PARAMS are skipped", func(t *testing.T) {
		t.Setenv("SOPS_AGE_KEY_PARAMS", "\n/real/key\n\n")
		t.Setenv("SOPS_AGE_KEY", "")
		staticAgeKey = ""

		err := loadAgeKeysFromSSM(makeClients(map[string]string{
			"/real/key": "AGE-REAL-KEY",
		}))
		assert.NoError(t, err)
		assert.Equal(t, "AGE-REAL-KEY", os.Getenv("SOPS_AGE_KEY"))
	})

	t.Run("SSM fetch error is propagated", func(t *testing.T) {
		t.Setenv("SOPS_AGE_KEY_PARAMS", "/bad/param")
		t.Setenv("SOPS_AGE_KEY", "")
		staticAgeKey = ""

		err := loadAgeKeysFromSSM(makeClients(map[string]string{}))
		assert.Error(t, err)
		assert.Contains(t, err.Error(), "/bad/param")
	})
}
