package main

import (
	"context"
	"encoding/json"
	"os"
	"slices"
	"strings"
	"testing"
	"time"

	"github.com/aws/aws-lambda-go/cfn"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/markussiebert/cdk-sops-secrets/internal/client"
	"github.com/stretchr/testify/assert"
)

func prepareUploadS3(t *testing.T, resourceProps map[string]interface{}) {
	cfg, err := config.LoadDefaultConfig(context.Background())

	if err != nil {
		t.Errorf("unable to load SDK config, %v", err)
	}

	sopsS3File := resourceProps["SopsS3File"].(map[string]interface{})

	s3Key := sopsS3File["Key"].(string)
	s3Bucket := sopsS3File["Bucket"].(string)
	localFile := "../" + s3Key

	file, err := os.Open(localFile)
	if err != nil {
		t.Errorf("failed to open file: %v", err)
	}
	defer file.Close()

	s3Client := s3.NewFromConfig(cfg)
	_, err = s3Client.PutObject(context.Background(), &s3.PutObjectInput{
		Bucket: &s3Bucket,
		Key:    &s3Key,
		Body:   file,
	})

	if err != nil {
		t.Errorf("failed to upload file to S3: %v", err)
	}
}

func prepareSecret(t *testing.T, resourceProps map[string]interface{}) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		t.Errorf("unable to load SDK config, %v", err)
	}

	secretsManagerClient := secretsmanager.NewFromConfig(cfg)

	secretName := resourceProps["Target"].(string)
	ResourceType := resourceProps["ResourceType"].(string)

	if ResourceType != "SECRET" && ResourceType != "SECRET_BINARY" && ResourceType != "SECRET_RAW" {
		return
	}

	secretValue := "empty"

	_, err = secretsManagerClient.CreateSecret(context.Background(), &secretsmanager.CreateSecretInput{
		Name:         &secretName,
		SecretString: &secretValue,
	})

	if err != nil {
		t.Logf("failed to create secret: %v", err)
	}
}

func cleanupSecret(t *testing.T, resourceProps map[string]interface{}) {

	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		t.Errorf("unable to load SDK config, %v", err)
	}

	secretsManagerClient := secretsmanager.NewFromConfig(cfg)

	secretName := resourceProps["Target"].(string)
	ResourceType := resourceProps["ResourceType"].(string)

	if ResourceType != "SECRET" && ResourceType != "SECRET_BINARY" {
		return
	}

	forceDelete := true

	_, err = secretsManagerClient.DeleteSecret(context.Background(), &secretsmanager.DeleteSecretInput{
		SecretId:                   &secretName,
		ForceDeleteWithoutRecovery: &forceDelete,
	})

	if err != nil {
		t.Logf("failed to delete secret: %v", err)
	}
}

func snapshotSecret(t *testing.T, resourceProps map[string]interface{}) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		t.Errorf("unable to load SDK config, %v", err)
	}

	secretsManagerClient := secretsmanager.NewFromConfig(cfg)

	secretName := resourceProps["Target"].(string)
	ResourceType := resourceProps["ResourceType"].(string)

	if ResourceType != "SECRET" && ResourceType != "SECRET_RAW" && ResourceType != "SECRET_BINARY" {
		return
	}

	resp, err := secretsManagerClient.GetSecretValue(context.Background(), &secretsmanager.GetSecretValueInput{
		SecretId: &secretName,
	})
	snaps.WithConfig(snaps.Filename(t.Name())).MatchSnapshot(t, resp.Name, resp.SecretBinary, resp.SecretString)

	if err != nil {
		t.Logf("failed to snapshot secret: %v", err)
	}
}

func cleanupParameter(t *testing.T, resourceProps map[string]interface{}) {
	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		t.Errorf("unable to load SDK config, %v", err)
	}

	ssmClient := ssm.NewFromConfig(cfg)
	parameterNames := []string{}
	parameterName := resourceProps["Target"].(string)
	ResourceType := resourceProps["ResourceType"].(string)

	if ResourceType == "PARAMETER_MULTI" {
		input := &ssm.DescribeParametersInput{}
		for {
			output, err := ssmClient.DescribeParameters(context.Background(), input)
			if err != nil {
				t.Errorf("failed to describe parameters: %v", err)
				return
			}
			for _, param := range output.Parameters {
				if strings.HasPrefix(*param.Name, parameterName) {
					parameterNames = append(parameterNames, *param.Name)
				}
			}
			if output.NextToken == nil {
				break
			}
			input.NextToken = output.NextToken
		}
	} else if ResourceType == "PARAMETER" {
		parameterNames = append(parameterNames, parameterName)
	} else {
		t.Logf("unknown resource type: %s", ResourceType)
		return
	}

	for _, p := range parameterNames {
		t.Logf("deleting parameter %s", p)
		_, err = ssmClient.DeleteParameter(context.Background(), &ssm.DeleteParameterInput{
			Name: &p,
		})
		if err != nil {
			t.Logf("failed to delete parameter: %v", err)
		}
	}
}

func snapshotParameter(t *testing.T, resourceProps map[string]interface{}) {

	cfg, err := config.LoadDefaultConfig(context.Background())
	if err != nil {
		t.Errorf("unable to load SDK config, %v", err)
	}

	ssmClient := ssm.NewFromConfig(cfg)
	parameterNames := []string{}
	parameterName := resourceProps["Target"].(string)
	ResourceType := resourceProps["ResourceType"].(string)

	if ResourceType == "PARAMETER_MULTI" {
		input := &ssm.DescribeParametersInput{}
		for {
			output, err := ssmClient.DescribeParameters(context.Background(), input)
			if err != nil {
				t.Errorf("failed to describe parameters: %v", err)
				return
			}
			for _, param := range output.Parameters {
				if strings.HasPrefix(*param.Name, parameterName) {
					parameterNames = append(parameterNames, *param.Name)
				}
			}
			if output.NextToken == nil {
				break
			}
			input.NextToken = output.NextToken
		}
	} else if ResourceType == "PARAMETER" {
		parameterNames = append(parameterNames, parameterName)
	} else {
		t.Logf("unknown resource type: %s", ResourceType)
		return
	}

	for _, p := range parameterNames {
		t.Logf("snapshotting parameter %s", p)
		resp, err := ssmClient.GetParameter(context.Background(), &ssm.GetParameterInput{
			Name:           &p,
			WithDecryption: aws.Bool(true),
		})
		snaps.WithConfig(snaps.Filename(t.Name())).MatchSnapshot(t, resp.Parameter.Name, resp.Parameter.Value)
		if err != nil {
			t.Logf("failed to snapshot parameter: %v", err)
		}
	}

}

func TestIntegration_HandleRequestWithClients(t *testing.T) {
	//if os.Getenv("INTEGRATION") == "" {
	//	return
	//}
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

			prepareUploadS3(t, tt.event.ResourceProperties)
			prepareSecret(t, tt.event.ResourceProperties)

			clients := client.CreateAwsClients(context.Background())

			physicalResourceID, data, err := HandleRequestWithClients(clients, tt.event)

			snapshotSecret(t, tt.event.ResourceProperties)
			snapshotParameter(t, tt.event.ResourceProperties)

			assert.NoError(t, err)
			assert.NotEmpty(t, physicalResourceID)
			assert.NotNil(t, data)
		})
	}
}

func TestIntegration_Cleanup(t *testing.T) {
	if os.Getenv("INTEGRATION") == "" {
		return
	}
	t.Logf("Running at: %s", time.Now().String()) // Forces re-run
	os.Setenv("AWS_REGION", "eu-central-1")

	files, err := os.ReadDir("events")
	if err != nil {
		t.Fatalf("failed to read events directory: %v", err)
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

		t.Run(file.Name(), func(t *testing.T) {

			cleanupSecret(t, resourceProperties)
			cleanupParameter(t, resourceProperties)

		})
	}
}

func Test_CleanUpHard(t *testing.T) {
	if os.Getenv("DANGER_HARD_CLEAN") == "" {
		return
	}
	deleteAllSecrets()
	deleteAllParametersExcept([]string{"/cdk-bootstrap/hnb659fds/version"})
}

func deleteAllSecrets() {
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRetryMaxAttempts(5))

	if err != nil {
		panic(err)
	}

	secretsManagerClient := secretsmanager.NewFromConfig(cfg)

	input := &secretsmanager.ListSecretsInput{
		IncludePlannedDeletion: aws.Bool(true),
	}
	for {
		output, err := secretsManagerClient.ListSecrets(context.Background(), input)
		if err != nil {
			panic(err)
		}
		for _, secret := range output.SecretList {
			_, err := secretsManagerClient.DeleteSecret(context.Background(), &secretsmanager.DeleteSecretInput{
				ForceDeleteWithoutRecovery: aws.Bool(true),
				SecretId:                   secret.ARN,
			})
			if err != nil {
				panic(err)
			}
		}
		if output.NextToken == nil {
			break
		}
		input.NextToken = output.NextToken
	}
}

func deleteAllParametersExcept(except []string) {
	cfg, err := config.LoadDefaultConfig(context.Background(), config.WithRetryMaxAttempts(5))

	if err != nil {
		panic(err)
	}

	ssmClient := ssm.NewFromConfig(cfg)

	input := &ssm.DescribeParametersInput{}
	for {
		output, err := ssmClient.DescribeParameters(context.Background(), input)
		if err != nil {
			panic(err)
		}
		for _, param := range output.Parameters {
			if !slices.Contains(except, *param.Name) {
				_, err := ssmClient.DeleteParameter(context.Background(), &ssm.DeleteParameterInput{
					Name: param.Name,
				})
				if err != nil {
					panic(err)
				}
			}
		}
		if output.NextToken == nil {
			break
		}
		input.NextToken = output.NextToken
	}
}
