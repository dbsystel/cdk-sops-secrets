package main

import (
	"fmt"
	"log/slog"
	"strings"

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
	if props.properties.ResourceType == event.SECRET {
		seperator := "."
		if props.properties.FlattenSeparator != nil {
			seperator = *props.properties.FlattenSeparator
		}
		logger.Info("Flattening secret data", "seperator", seperator)
		if err := props.secretDecryptedData.Flatten(seperator); err != nil {
			return props.properties.GeneratePhysicalResourceId(), nil, err
		}
		logger.Info("Stringifying secret data")
		if err := props.secretDecryptedData.StringifyValues(); err != nil {
			return props.properties.GeneratePhysicalResourceId(), nil, err
		}
		logger.Info("Converting secret data to JSON")
		outData, outDataErr = props.secretDecryptedData.ToJSON()
	} else {
		logger.Info("Getting raw secret data")
		outData, outDataErr = props.secretDecryptedData.GetRaw()
	}

	if outDataErr != nil {
		return props.properties.GeneratePhysicalResourceId(), nil, outDataErr
	}

	putSecretValueResp, putSecretValueRespErr := props.clients.SecretsManagerPutSecretValue(*props.secretDecryptedData.Hash, props.properties.Target, outData)
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
		// Ass we flatten array to [number] path notations, we have to fix this for parameter store
		fixedKey := strings.ReplaceAll(key, "[", seperator)
		fixedKey = strings.ReplaceAll(fixedKey, "]", seperator)
		fixedKey = strings.TrimSuffix(fixedKey, seperator)
		fixedKey = strings.ReplaceAll(fixedKey, seperator+seperator, seperator)
		// Whitespaces are also not allowed, maybe more characters
		fixedKey = strings.ReplaceAll(fixedKey, " ", "_")
		// Prefix with Target
		fixedKey = props.properties.Target + fixedKey
		_, err := props.clients.SsmPutParameter(fixedKey, &value, *props.properties.EncryptionKey)
		if err != nil {
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
		return props.properties.GeneratePhysicalResourceId(), nil, fmt.Errorf("failed to update ssm parameter:\n%v", err)
	}
	logger.Info("Parameter written", "Parameter", props.properties.Target)

	logger.Info("Secret data updated", "ARN", props.properties.GeneratePhysicalResourceId())
	return props.properties.GeneratePhysicalResourceId(), map[string]interface{}{
		"ParameterName": props.properties.Target,
		"Version":       putParameterResponse.Version,
		"Tier":          putParameterResponse.Tier,
	}, nil
}
