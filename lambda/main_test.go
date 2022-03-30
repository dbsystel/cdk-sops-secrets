package main

import (
	"context"
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"testing"
	"time"

	"github.com/aws/aws-lambda-go/cfn"
	"github.com/aws/aws-lambda-go/lambdacontext"
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
	return &secretsmanager.PutSecretValueOutput{
		ARN:       input.SecretId,
		VersionId: &versionId,
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

func Test_GetS3FileContent(t *testing.T) {
	mocks := &AWS{
		secretsmanager: &SecretsManagerMockClient{},
		s3downlaoder:   &S3ManagerMockClient{},
	}
	data, err := mocks.getS3FileContent(SOPSS3File{
		Bucket: "..",
		Key:    "../test-secrets/sopsfile.enc-age.json",
		Format: "..",
	})
	check(err)
	expected, err := os.ReadFile("../test-secrets/sopsfile.enc-age.json")
	if reflect.DeepEqual(data, expected) {
		t.Log("getS3FileContent produces expected result")
	} else {
		t.Errorf("getS3FileContent produces unexcaptable match:\n%s\nexpceted:\n%s", string(data), string(expected))
	}
}

func Test_UpdateSecret(t *testing.T) {
	mocks := &AWS{
		secretsmanager: &SecretsManagerMockClient{},
		s3downlaoder:   &S3ManagerMockClient{},
	}
	arn := "arn:${Partition}:secretsmanager:${Region}:${Account}:secret:${SecretId}"
	secretValue := []byte("some-secret-data")

	response, _ := mocks.updateSecret(arn, secretValue)

	versionId := fmt.Sprintf("%x", sha256.Sum256(secretValue))
	expected := &secretsmanager.PutSecretValueOutput{
		ARN:       &arn,
		VersionId: &versionId,
	}

	if reflect.DeepEqual(response, expected) {
		t.Log("updateSecret produces expected result")
	} else {
		t.Errorf("updateSecret produces unexcaptable match:\n%s\nexpceted:\n%s", response, expected)
	}
}
func Test_DecryptSopsFileContent(t *testing.T) {

	os.Setenv("SOPS_AGE_KEY", "AGE-SECRET-KEY-1EFUWJ0G2XJTJFWTAM2DGMA4VCK3R05W58FSMHZP3MZQ0ZTAQEAFQC6T7T3")

	sopsEncrypted, err := os.ReadFile("../test-secrets/sopsfile.enc-age.json")
	check(err)
	sopsDecrypted, err := decryptSopsFileContent(sopsEncrypted, SOPSS3File{Bucket: "", Key: "", Format: "json"})
	check(err)
	sopsExpected, err := os.ReadFile("../test-secrets/sopsfile.json")
	check(err)

	sopsDecryptedJ := marshallAny(sopsDecrypted)
	sopsExpectedJ := marshallAny(sopsExpected)

	if reflect.DeepEqual(sopsDecryptedJ, sopsExpectedJ) {
		t.Log("decryptSopsFileContent produces expected result")
	} else {
		t.Errorf("decryptSopsFileContent produces unexcaptable match:\n%s\nexpceted:\n%s", sopsDecryptedJ, sopsExpectedJ)
	}
}

func Test_FullWorkflow_Create(t *testing.T) {
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
	inputJson := ReadJSONFromFile(t, "event_create.json")
	var event cfn.Event
	err := json.Unmarshal(inputJson, &event)

	phys, data, err := mocks.syncSopsToSecretsmanager(ctx, event)
	check(err)
	log.Println(phys)
	log.Println(data)
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
	log.Println(phys)
	log.Println(data)
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
	log.Println(phys)
	log.Println(data)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func marshallAny(input []byte) map[string]interface{} {
	var returnValue map[string]interface{}
	err := json.Unmarshal(input, &returnValue)
	check(err)
	return returnValue
}

//func TestMain(t *testing.T) {
//	d := time.Now().Add(50 * time.Millisecond)
//	os.Setenv("AWS_LAMBDA_FUNCTION_NAME", "blank-go")
//	ctx, _ := context.WithDeadline(context.Background(), d)
//	ctx = lambdacontext.NewContext(ctx, &lambdacontext.LambdaContext{
//		AwsRequestID:       "495b12a8-xmpl-4eca-8168-160484189f99",
//		InvokedFunctionArn: "arn:aws:lambda:us-east-2:123456789012:function:blank-go",
//	})

//	if err != nil {
//		t.Errorf("could not unmarshal event. details: %v", err)
//	}
//	//var inputEvent SOPSEvent
//	handleRequest(ctx, event)
//	if err != nil {
//		t.Log(err)
//	}
//}

func ReadJSONFromFile(t *testing.T, inputFile string) []byte {
	inputJSON, err := ioutil.ReadFile(inputFile)
	if err != nil {
		t.Errorf("could not open test file. details: %v", err)
	}

	return inputJSON
}
