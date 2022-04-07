package main

import (
	"context"
	"encoding/json"
	"os"
	"testing"
	"time"

	"github.com/aws/aws-lambda-go/cfn"
	"github.com/aws/aws-lambda-go/lambdacontext"
	"github.com/gkampitakis/go-snaps/snaps"
)

func Test_FullWorkflow_Create_S3_JSON_Simple(t *testing.T) {
	os.Setenv("SOPS_AGE_KEY", "AGE-SECRET-KEY-1EFUWJ0G2XJTJFWTAM2DGMA4VCK3R05W58FSMHZP3MZQ0ZTAQEAFQC6T7T3")
	mocks := &AWS{
		secretsmanager: &SecretsManagerMockClient{
			t: t,
		},
		s3downlaoder: &S3ManagerMockClient{
			t: t,
		},
	}
	d := time.Now().Add(50 * time.Millisecond)
	ctx, _ := context.WithDeadline(context.Background(), d)
	ctx = lambdacontext.NewContext(ctx, &lambdacontext.LambdaContext{
		AwsRequestID:       "AwsRequestID",
		InvokedFunctionArn: "arn:aws:lambda:us-east-2:123456789012:function:cdk-sops-secrets",
	})
	inputJson := ReadJSONFromFile(t, "events/event_create_s3_json_simple.json")
	var event cfn.Event
	err := json.Unmarshal(inputJson, &event)

	phys, data, err := mocks.syncSopsToSecretsmanager(ctx, event)
	check(err)
	snaps.MatchSnapshot(t, ">>>syncSopsToSecretsmanager", phys, data, err)
}

func Test_FullWorkflow_Create_S3_JSON_Complex(t *testing.T) {
	os.Setenv("SOPS_AGE_KEY", "AGE-SECRET-KEY-1EFUWJ0G2XJTJFWTAM2DGMA4VCK3R05W58FSMHZP3MZQ0ZTAQEAFQC6T7T3")
	mocks := &AWS{
		secretsmanager: &SecretsManagerMockClient{
			t: t,
		},
		s3downlaoder: &S3ManagerMockClient{
			t: t,
		},
	}
	d := time.Now().Add(50 * time.Millisecond)
	ctx, _ := context.WithDeadline(context.Background(), d)
	ctx = lambdacontext.NewContext(ctx, &lambdacontext.LambdaContext{
		AwsRequestID:       "AwsRequestID",
		InvokedFunctionArn: "arn:aws:lambda:us-east-2:123456789012:function:cdk-sops-secrets",
	})
	inputJson := ReadJSONFromFile(t, "events/event_create_s3_json_complex.json")
	var event cfn.Event
	err := json.Unmarshal(inputJson, &event)

	phys, data, err := mocks.syncSopsToSecretsmanager(ctx, event)
	check(err)
	snaps.MatchSnapshot(t, ">>>syncSopsToSecretsmanager", phys, data, err)
}

func Test_FullWorkflow_Create_S3_JSON_Complex_StringifyValues(t *testing.T) {
	os.Setenv("SOPS_AGE_KEY", "AGE-SECRET-KEY-1EFUWJ0G2XJTJFWTAM2DGMA4VCK3R05W58FSMHZP3MZQ0ZTAQEAFQC6T7T3")
	mocks := &AWS{
		secretsmanager: &SecretsManagerMockClient{
			t: t,
		},
		s3downlaoder: &S3ManagerMockClient{
			t: t,
		},
	}
	d := time.Now().Add(50 * time.Millisecond)
	ctx, _ := context.WithDeadline(context.Background(), d)
	ctx = lambdacontext.NewContext(ctx, &lambdacontext.LambdaContext{
		AwsRequestID:       "AwsRequestID",
		InvokedFunctionArn: "arn:aws:lambda:us-east-2:123456789012:function:cdk-sops-secrets",
	})
	inputJson := ReadJSONFromFile(t, "events/event_create_s3_json_complex_stringify.json")
	var event cfn.Event
	err := json.Unmarshal(inputJson, &event)

	phys, data, err := mocks.syncSopsToSecretsmanager(ctx, event)
	check(err)
	snaps.MatchSnapshot(t, ">>>syncSopsToSecretsmanager", phys, data, err)
}
func Test_FullWorkflow_Create_S3_JSON_Complex_Flat(t *testing.T) {
	os.Setenv("SOPS_AGE_KEY", "AGE-SECRET-KEY-1EFUWJ0G2XJTJFWTAM2DGMA4VCK3R05W58FSMHZP3MZQ0ZTAQEAFQC6T7T3")
	mocks := &AWS{
		secretsmanager: &SecretsManagerMockClient{
			t: t,
		},
		s3downlaoder: &S3ManagerMockClient{
			t: t,
		},
	}
	d := time.Now().Add(50 * time.Millisecond)
	ctx, _ := context.WithDeadline(context.Background(), d)
	ctx = lambdacontext.NewContext(ctx, &lambdacontext.LambdaContext{
		AwsRequestID:       "AwsRequestID",
		InvokedFunctionArn: "arn:aws:lambda:us-east-2:123456789012:function:cdk-sops-secrets",
	})
	inputJson := ReadJSONFromFile(t, "events/event_create_s3_json_complex_flat.json")
	var event cfn.Event
	err := json.Unmarshal(inputJson, &event)

	phys, data, err := mocks.syncSopsToSecretsmanager(ctx, event)
	check(err)
	snaps.MatchSnapshot(t, ">>>syncSopsToSecretsmanager", phys, data, err)
}
