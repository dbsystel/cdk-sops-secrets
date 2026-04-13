package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/cfn"
	runtime "github.com/aws/aws-lambda-go/lambda"
	"github.com/markussiebert/cdk-sops-secrets/internal/client"
	"github.com/markussiebert/cdk-sops-secrets/internal/event"
)

// staticAgeKey captures the value of SOPS_AGE_KEY at Lambda cold-start so that
// loadAgeKeysFromSSM can always rebuild the combined key list correctly across
// warm invocations without duplicating the static portion.
var staticAgeKey = os.Getenv("SOPS_AGE_KEY")

// loadAgeKeysFromSSM reads parameter names from SOPS_AGE_KEY_PARAMS (newline-
// separated), fetches each SecureString from SSM Parameter Store with decryption,
// and sets SOPS_AGE_KEY to the combined value of any static key already present
// at cold-start plus all SSM-fetched keys.  The function is a no-op when
// SOPS_AGE_KEY_PARAMS is unset or empty.
func loadAgeKeysFromSSM(clients client.AwsClient) error {
	logger := slog.With("Package", "main", "Function", "loadAgeKeysFromSSM")

	paramList := os.Getenv("SOPS_AGE_KEY_PARAMS")
	if paramList == "" {
		return nil
	}

	var ageKeys []string
	if staticAgeKey != "" {
		ageKeys = append(ageKeys, staticAgeKey)
	}

	for _, paramName := range strings.Split(paramList, "\n") {
		paramName = strings.TrimSpace(paramName)
		if paramName == "" {
			continue
		}
		logger.Info("Fetching age key from SSM Parameter Store", "Parameter", paramName)
		value, err := clients.SsmGetParameter(paramName)
		if err != nil {
			return fmt.Errorf("failed to fetch age key from SSM parameter %s: %v", paramName, err)
		}
		if value == nil {
			return fmt.Errorf("received nil value for SSM parameter %s", paramName)
		}
		ageKeys = append(ageKeys, *value)
	}

	return os.Setenv("SOPS_AGE_KEY", strings.Join(ageKeys, "\n"))
}

func HandleRequestWithClients(clients client.AwsClient, e cfn.Event) (physicalResourceID string, data map[string]interface{}, err error) {
	logger := slog.With("Package", "main", "Function", "HandleRequestWithClients")
	logger.Debug("Incoming Event", "Event", e)

	// If it's a delete request, we don't have to do anything
	if e.RequestType == cfn.RequestDelete {
		return event.GenerateTempPhysicalResourceId(), nil, nil
	}
	// We have to run this code only, if it is a CloudFormation Create or Update request
	if e.RequestType != cfn.RequestCreate && e.RequestType != cfn.RequestUpdate {
		return event.GenerateTempPhysicalResourceId(), nil, fmt.Errorf("requestType '%s' not supported", e.RequestType)
	}

	// Get the event input from the cloudformation event
	props, err := event.FromCfnEvent(e)
	if err != nil {
		return "", nil, err
	}

	// Get the encrypted secret input provided by the user
	secretEncrypted, secretEncryptedErr := props.GetEncryptedSopsSecret(clients)
	if secretEncryptedErr != nil {
		return props.GeneratePhysicalResourceId(), nil, secretEncryptedErr
	}

	// Fetch SSM age keys only when the SOPS file actually uses age encryption.
	if secretEncrypted.UsesAgeEncryption() {
		if err := loadAgeKeysFromSSM(clients); err != nil {
			return props.GeneratePhysicalResourceId(), nil, err
		}
	}

	// Decrypt the secret input with sops
	secretDecrypted, secretDecryptedErr := secretEncrypted.Decrypt()
	if secretDecryptedErr != nil {
		return props.GeneratePhysicalResourceId(), nil, secretDecryptedErr
	}

	// Generate a data object by parsing the decrypted secret depending on the data input type
	secretDecryptedData, secretDecryptedDataErr := secretDecrypted.ToData()
	if secretDecryptedDataErr != nil {
		return props.GeneratePhysicalResourceId(), nil, secretDecryptedDataErr
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
		return props.GeneratePhysicalResourceId(), nil, fmt.Errorf("unsupported resource type %s", props.ResourceType)
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
