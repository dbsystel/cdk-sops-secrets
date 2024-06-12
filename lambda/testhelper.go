package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"io/ioutil"
	"os"
	"testing"
	"time"

	"github.com/aws/aws-lambda-go/cfn"
	"github.com/aws/aws-lambda-go/lambdacontext"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func UnmarshalAny(input []byte) map[string]interface{} {
	var returnValue map[string]interface{}
	err := json.Unmarshal(input, &returnValue)
	check(err)
	return returnValue
}

func ReadFile(path string) []byte {
	ret, err := os.ReadFile(path)
	check(err)
	return ret
}

func ReadJSONFromFile(t *testing.T, inputFile string) []byte {
	inputJSON, err := ioutil.ReadFile(inputFile)
	if err != nil {
		t.Errorf("could not open test file. details: %v", err)
	}

	return inputJSON
}

func base64FromFile(file string) string {
	inputData, err := ioutil.ReadFile(file)
	check(err)
	return base64.StdEncoding.EncodeToString(inputData)
}

func fileToInline(event cfn.Event) cfn.Event {
	sopsProps, err := toSopsSyncResourcePropertys(&event.ResourceProperties)
	check(err)

	inlineProps := SopsInline{
		Content: base64FromFile(sopsProps.SopsS3File.Key),
		Hash:    "MyHash",
	}
	event.ResourceProperties["SopsInline"] = inlineProps
	delete(event.ResourceProperties, "SopsS3File")
	return event
}
func getMocks(t *testing.T) *AWS {
	return &AWS{
		secretsmanager: &SecretsManagerMockClient{
			t: t,
		},
		s3downlaoder: &S3ManagerMockClient{
			t: t,
		},
		ssm: &MockSSMClient{
			t: t,
		},
	}
}
func prepareHandler(t *testing.T, eventFile string) (*AWS, context.Context, cfn.Event) {
	os.Setenv("SOPS_AGE_KEY", "AGE-SECRET-KEY-1EFUWJ0G2XJTJFWTAM2DGMA4VCK3R05W58FSMHZP3MZQ0ZTAQEAFQC6T7T3")
	mocks := getMocks(t)
	d := time.Now().Add(50 * time.Millisecond)
	ctx, _ := context.WithDeadline(context.Background(), d)
	ctx = lambdacontext.NewContext(ctx, &lambdacontext.LambdaContext{
		AwsRequestID:       "AwsRequestID",
		InvokedFunctionArn: "arn:aws:lambda:us-east-2:123456789012:function:cdk-sops-secrets",
	})
	inputJson := ReadJSONFromFile(t, eventFile)
	var event cfn.Event
	err := json.Unmarshal(inputJson, &event)
	check(err)
	return mocks, ctx, event
}
