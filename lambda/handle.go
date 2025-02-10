package main

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/aws/smithy-go/ptr"
	"github.com/markussiebert/cdk-sops-secrets/internal/client"
	"github.com/markussiebert/cdk-sops-secrets/internal/data"
	"github.com/markussiebert/cdk-sops-secrets/internal/event"
)

type BaseProps struct {
	properties          *event.SopsSyncResourcePropertys
	clients             client.AwsClient
	secretDecryptedData *data.Data
}

func handleSecret(props BaseProps) (physicalResourceID string, data map[string]interface{}, err error) {
	logger := slog.With("Package", "main", "Function", "handleSecret")
	var outData *[]byte
	var outDataErr error
	logger.Info("Handling secret data", "ResourceType", props.properties.ResourceType)

	var binary *bool

	// Prepare the data depending on the resource type
	switch props.properties.ResourceType {
	case event.SECRET:
		// The secretsmanager does not support nested json objects, so we have to flatten the data
		seperator := "."
		if props.properties.FlattenSeparator != nil {
			seperator = *props.properties.FlattenSeparator
		}
		logger.Info("Flattening secret data", "seperator", seperator, "ResourceType", props.properties.ResourceType)
		if err := props.secretDecryptedData.Flatten(seperator); err != nil {
			return props.properties.GeneratePhysicalResourceId(), nil, err
		}

		// Make sure the values are strings to avoid issues while storing them in the secrets manager
		logger.Info("Stringifying secret data", "ResourceType", props.properties.ResourceType)
		if err := props.secretDecryptedData.StringifyValues(); err != nil {
			return props.properties.GeneratePhysicalResourceId(), nil, err
		}

		// Secretsmanager values have to be stored as JSON
		logger.Info("Converting secret data to JSON", "ResourceType", props.properties.ResourceType)
		outData, outDataErr = props.secretDecryptedData.ToJSON()
		binary = ptr.Bool(false)
	case event.SECRET_RAW:
		logger.Info("Getting raw secret data", "ResourceType", props.properties.ResourceType)
		outData, outDataErr = props.secretDecryptedData.GetRaw()
		binary = ptr.Bool(false)
	case event.SECRET_BINARY:
		logger.Info("Getting raw secret data", "ResourceType", props.properties.ResourceType)
		outData, outDataErr = props.secretDecryptedData.GetRaw()
		binary = ptr.Bool(true)
	default:
		return props.properties.GeneratePhysicalResourceId(), nil, fmt.Errorf("unknown ResourceType: %s", props.properties.ResourceType)
	}

	if outDataErr != nil {
		return props.properties.GeneratePhysicalResourceId(), nil, outDataErr
	}

	putSecretValueResp, putSecretValueRespErr := props.clients.SecretsManagerPutSecretValue(*props.secretDecryptedData.Hash, props.properties.Target, outData, binary)
	if putSecretValueRespErr != nil {
		return props.properties.GeneratePhysicalResourceId(), nil, putSecretValueRespErr
	}

	logger.Info("Secret data updated", "ARN", *putSecretValueResp.ARN)

	return *putSecretValueResp.ARN, map[string]interface{}{
		"ARN":           *putSecretValueResp.ARN,
		"Name":          *putSecretValueResp.Name,
		"VersionStages": putSecretValueResp.VersionStages,
		"VersionId":     *putSecretValueResp.VersionId,
	}, nil
}

func handleParameterMulti(props BaseProps) (physicalResourceID string, data map[string]interface{}, err error) {
	logger := slog.With("Package", "main", "Function", "handleParameterMulti")
	seperator := "/"
	if props.properties.FlattenSeparator != nil {
		seperator = *props.properties.FlattenSeparator
	}
	logger.Info("Flattening secret data", "seperator", seperator)
	if err := props.secretDecryptedData.Flatten(seperator); err != nil {
		return props.properties.GeneratePhysicalResourceId(), nil, err
	}
	logger.Info("Stringify values in data")
	if err := props.secretDecryptedData.StringifyValues(); err != nil {
		return props.properties.GeneratePhysicalResourceId(), nil, err
	}
	logger.Info("Converting secret data to map")
	outData, outDataErr := props.secretDecryptedData.ToStringMap()
	if outDataErr != nil {
		return props.properties.GeneratePhysicalResourceId(), nil, outDataErr
	}
	if props.properties.EncryptionKey == nil {
		return props.properties.GeneratePhysicalResourceId(), nil, fmt.Errorf("EncryptionKey must be set for PARAMETER_MULTI")
	}
	logger.Info("Writing secret data to parameter store")
	for key, value := range outData {
		// As we flatten array to [number] path notations, we have to fix this for parameter store
		fixedKey := strings.ReplaceAll(key, "[", seperator)
		fixedKey = strings.ReplaceAll(fixedKey, "]", seperator)
		fixedKey = strings.TrimSuffix(fixedKey, seperator)
		fixedKey = strings.ReplaceAll(fixedKey, seperator+seperator, seperator)
		// The secret name can contain ASCII letters, numbers, and the following characters: /_+=.@-
		allowedChars := "/_+=.@-"
		for i, char := range fixedKey {
			if !(char >= 'a' && char <= 'z') && !(char >= 'A' && char <= 'Z') && !(char >= '0' && char <= '9') && !strings.ContainsRune(allowedChars, char) {
				fixedKey = fixedKey[:i] + "_" + fixedKey[i+1:]
			}
		}
		// Prefix with Target
		fixedKey = props.properties.Target + fixedKey
		// FixValue (can't be empty)
		if len(value) == 0 {
			value = []byte("<empty>")
		}
		_, err := props.clients.SsmPutParameter(fixedKey, &value, *props.properties.EncryptionKey)
		if err != nil {
			logger.Error("Failed to write parameter", "Parameter", fixedKey, "Error", err)
			return props.properties.GeneratePhysicalResourceId(), nil, fmt.Errorf("failed to update ssm parameter:\n%v", err)
		}
		logger.Info("Parameter written", "Parameter", fixedKey)
	}

	logger.Info("Secret data updated", "ARN", props.properties.GeneratePhysicalResourceId())

	return props.properties.GeneratePhysicalResourceId(), map[string]interface{}{
		"Prefix": props.properties.Target,
		"Count":  len(outData),
	}, nil
}

func handleParameter(props BaseProps) (physicalResourceID string, data map[string]interface{}, err error) {
	logger := slog.With("Package", "main", "Function", "handleParameter")
	logger.Info("Getting raw secret data")
	outData, outDataErr := props.secretDecryptedData.GetRaw()
	if outDataErr != nil {
		return props.properties.GeneratePhysicalResourceId(), nil, outDataErr
	}
	if props.properties.EncryptionKey == nil {
		return "", nil, fmt.Errorf("EncryptionKey must be set for PARAMETER")
	}
	putParameterResponse, putParameterResponseErr := props.clients.SsmPutParameter(props.properties.Target, outData, *props.properties.EncryptionKey)
	if putParameterResponseErr != nil {
		logger.Error("Failed to write parameter", "Parameter", props.properties.Target, "Error", putParameterResponseErr)
		return props.properties.GeneratePhysicalResourceId(), nil, fmt.Errorf("failed to update ssm parameter:\n%v", putParameterResponseErr)
	}
	logger.Info("Parameter written", "Parameter", props.properties.Target)

	logger.Info("Secret data updated", "ARN", props.properties.GeneratePhysicalResourceId())
	return props.properties.GeneratePhysicalResourceId(), map[string]interface{}{
		"ParameterName": props.properties.Target,
		"Version":       putParameterResponse.Version,
		"Tier":          putParameterResponse.Tier,
	}, nil
}
