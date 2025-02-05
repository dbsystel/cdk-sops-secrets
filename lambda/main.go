package main

import (
	"context"
	"fmt"
	"regexp"

	"github.com/aws/aws-lambda-go/cfn"
	runtime "github.com/aws/aws-lambda-go/lambda"
	"github.com/markussiebert/cdk-sops-secrets/internal/client"
	"github.com/markussiebert/cdk-sops-secrets/internal/event"
)

func HandleRequestWithClients(clients client.AwsClient, e cfn.Event) (physicalResourceID string, data map[string]interface{}, err error) {
	// If it's a delete request, we don't have to do anything
	if e.RequestType == cfn.RequestDelete {
		return "", nil, nil
	}
	// We have to run this code only, if it is a CloudFormation Create or Update request
	if e.RequestType != cfn.RequestCreate && e.RequestType != cfn.RequestUpdate {
		return "", nil, fmt.Errorf("requestType '%s' not supported", e.RequestType)
	}

	props, err := event.FromCfnEvent(e)
	if err != nil {
		return "", nil, err
	}

	tempArn := props.GeneratePhysicalResourceId()

	secretEncrypted, secretEncryptedErr := props.GetEncryptedSopsSecret(clients)

	if secretEncryptedErr != nil {
		return "", nil, secretEncryptedErr
	}

	secretDecrypted, secretDecryptedErr := secretEncrypted.Decrypt()

	if secretDecryptedErr != nil {
		return "", nil, secretDecryptedErr
	}

	secretDecryptedData, secretDecryptedDataErr := secretDecrypted.ToData()

	if secretDecryptedDataErr != nil {
		return "", nil, secretDecryptedDataErr
	}

	if props.Flatten {
		err := secretDecryptedData.Flatten(props.FlattenSeparator)
		if err != nil {
			return "", nil, err
		}
	}

	if props.StringifyValues {
		err := secretDecryptedData.StringifyValues()
		if err != nil {
			return "", nil, err
		}
	}
	baseProps := BaseProps{
		clients:             clients,
		tempArn:             tempArn,
		secretDecryptedData: secretDecryptedData,
	}

	switch props.ResourceType {
	case "SECRET":
		return handleSecret(HandleSecretProps{
			BaseProps:    baseProps,
			hash:         secretEncrypted.Hash,
			secretArn:    props.SecretARN,
			outputFormat: props.OutputFormat,
		})
	case "PARAMETER_MULTI":
		return handleParameterMulti(HandleParameterMultiProps{
			BaseProps:          baseProps,
			parameterKeyPrefix: props.ParameterKeyPrefix,
			encryptionKey:      props.EncryptionKey,
		})
	case "PARAMETER":
		return handleParameter(HandleParameterProps{
			BaseProps:     baseProps,
			parameterName: props.ParameterName,
			encryptionKey: props.EncryptionKey,
			outputFormat:  props.OutputFormat,
		})
	default:
		return "", nil, fmt.Errorf("unsupported resource type %s", props.ResourceType)
	}
}

// Just a Wrapper function to allow injecting clients for testing
func HandleRequest(ctx context.Context, event cfn.Event) (physicalResourceID string, data map[string]interface{}, err error) {

	clients := client.CreateAwsClients(ctx)

	return HandleRequestWithClients(clients, event)
}

func main() {
	runtime.Start(cfn.LambdaWrap(HandleRequest))
}

func generatePhysicalResourceId(input string) string {
	re := regexp.MustCompile(`(^arn:.*:secretsmanager:|ssm:)(.*)`)
	return re.ReplaceAllString(input, `arn:custom:sopssync:$2`)
}
