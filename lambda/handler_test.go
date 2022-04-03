package main

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/aws/aws-lambda-go/cfn"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/gkampitakis/go-snaps/snaps"
)

func Test_FullWorkflow_Create(t *testing.T) {
	mocks := &AWS{
		secretsmanager: &SecretsManagerMockClient{},
		s3downlaoder:   &S3ManagerMockClient{},
	}
	d := time.Now().Add(50 * time.Millisecond)
	ctx, _ := context.WithDeadline(context.Background(), d)
	ctx = lambdacontext.NewContext(ctx, &lambdacontext.LambdaContext{
		AwsRequestID:       "AwsRequestID",
		InvokedFunctionArn: "arn:aws:lambda:us-east-2:123456789012:function:cdk-sops-secrets",
	})
	inputJson := ReadJSONFromFile(t, "event_create.json")
	var event cfn.Event
	err := json.Unmarshal(inputJson, &event)

	phys, data, err := mocks.syncSopsToSecretsmanager(ctx, event)
	check(err)
	snaps.MatchSnapshot(t, phys, data, err)
}

func Test_FullWorkflow_Update(t *testing.T) {
	mocks := &AWS{
		secretsmanager: &SecretsManagerMockClient{},
		s3downlaoder:   &S3ManagerMockClient{},
	}
	d := time.Now().Add(50 * time.Millisecond)
	ctx, _ := context.WithDeadline(context.Background(), d)
	ctx = lambdacontext.NewContext(ctx, &lambdacontext.LambdaContext{
		AwsRequestID:       "AwsRequestID",
		InvokedFunctionArn: "arn:aws:lambda:us-east-2:123456789012:function:blank-go",
	})
	inputJson := ReadJSONFromFile(t, "event_update.json")
	var event cfn.Event
	err := json.Unmarshal(inputJson, &event)

	phys, data, err := mocks.syncSopsToSecretsmanager(ctx, event)
	check(err)
	snaps.MatchSnapshot(t, phys, data, err)
}

func Test_FullWorkflow_Delete(t *testing.T) {
	mocks := &AWS{
		secretsmanager: &SecretsManagerMockClient{},
		s3downlaoder:   &S3ManagerMockClient{},
	}
	d := time.Now().Add(50 * time.Millisecond)
	ctx, _ := context.WithDeadline(context.Background(), d)
	ctx = lambdacontext.NewContext(ctx, &lambdacontext.LambdaContext{
		AwsRequestID:       "AwsRequestID",
		InvokedFunctionArn: "arn:aws:lambda:us-east-2:123456789012:function:blank-go",
	})
	inputJson := ReadJSONFromFile(t, "event_delete.json")
	var event cfn.Event
	err := json.Unmarshal(inputJson, &event)

	phys, data, err := mocks.syncSopsToSecretsmanager(ctx, event)
	check(err)
	snaps.MatchSnapshot(t, phys, data, err)
}
