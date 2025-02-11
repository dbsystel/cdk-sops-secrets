package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/aws/aws-lambda-go/cfn"
	runtime "github.com/aws/aws-lambda-go/lambda"
	"github.com/markussiebert/cdk-sops-secrets/internal/client"
	"github.com/markussiebert/cdk-sops-secrets/internal/event"
)

func HandleRequestWithClients(clients client.AwsClient, e cfn.Event) (physicalResourceID string, data map[string]interface{}, err error) {
	logger := slog.With("Package", "main", "Function", "HandleRequestWithClients")
	logger.Debug("Incoming Event", "Event", e)

	// If it's a delete request, we don't have to do anything
	if e.RequestType == cfn.RequestDelete {
		return "", nil, nil
	}
	// We have to run this code only, if it is a CloudFormation Create or Update request
	if e.RequestType != cfn.RequestCreate && e.RequestType != cfn.RequestUpdate {
		return "", nil, fmt.Errorf("requestType '%s' not supported", e.RequestType)
	}

	// Get the event input from the cloudformation event
	props, err := event.FromCfnEvent(e)
	if err != nil {
		return "", nil, err
	}

	// Get the encrypted secret input provided by the user
	secretEncrypted, secretEncryptedErr := props.GetEncryptedSopsSecret(clients)
	if secretEncryptedErr != nil {
		return "", nil, secretEncryptedErr
	}

	// Decrypt the secret input with sops
	secretDecrypted, secretDecryptedErr := secretEncrypted.Decrypt()
	if secretDecryptedErr != nil {
		return "", nil, secretDecryptedErr
	}

	// Generate a data object by parsing the decrypted secret depending on the data input type
	secretDecryptedData, secretDecryptedDataErr := secretDecrypted.ToData()
	if secretDecryptedDataErr != nil {
		return "", nil, secretDecryptedDataErr
	}

	baseProps := BaseProps{
		properties:          props,
		clients:             clients,
		secretDecryptedData: secretDecryptedData,
	}

	// Fill the secret values in the ressource depending on the ressource type
	switch props.ResourceType {
	case event.SECRET, event.SECRET_RAW, event.SECRET_BINARY:
		return handleSecret(baseProps)
	case event.PARAMETER_MULTI:
		return handleParameterMulti(baseProps)
	case event.PARAMETER:
		return handleParameter(baseProps)
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
