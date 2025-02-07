package main

import (
	"fmt"
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
	var outData *[]byte
	var outDataErr error
	if props.properties.ResourceType == event.SECRET {
		seperator := "."
		if props.properties.FlattenSeparator != nil {
			seperator = *props.properties.FlattenSeparator
		}
		if err := props.secretDecryptedData.Flatten(seperator); err != nil {
			return props.properties.GeneratePhysicalResourceId(), nil, err
		}
		if err := props.secretDecryptedData.StringifyValues(); err != nil {
			return props.properties.GeneratePhysicalResourceId(), nil, err
		}
		outData, outDataErr = props.secretDecryptedData.ToJSON()
	} else {
		outData, outDataErr = props.secretDecryptedData.GetRaw()
	}

	if outDataErr != nil {
		return props.properties.GeneratePhysicalResourceId(), nil, outDataErr
	}

	putSecretValueResp, putSecretValueRespErr := props.clients.SecretsManagerPutSecretValue(*props.secretDecryptedData.Hash, props.properties.Target, outData)
	if putSecretValueRespErr != nil {
		return props.properties.GeneratePhysicalResourceId(), nil, putSecretValueRespErr
	}
	return *putSecretValueResp.ARN, map[string]interface{}{
		"ARN":           *putSecretValueResp.ARN,
		"Name":          *putSecretValueResp.Name,
		"VersionStages": putSecretValueResp.VersionStages,
		"VersionId":     *putSecretValueResp.VersionId,
	}, nil
}

func handleParameterMulti(props BaseProps) (physicalResourceID string, data map[string]interface{}, err error) {
	seperator := "/"
	if props.properties.FlattenSeparator != nil {
		seperator = *props.properties.FlattenSeparator
	}
	if err := props.secretDecryptedData.Flatten(seperator); err != nil {
		return props.properties.GeneratePhysicalResourceId(), nil, err
	}
	outData, outDataErr := props.secretDecryptedData.ToStringMap()
	if outDataErr != nil {
		return props.properties.GeneratePhysicalResourceId(), nil, outDataErr
	}
	if props.properties.EncryptionKey == nil {
		return props.properties.GeneratePhysicalResourceId(), nil, fmt.Errorf("EncryptionKey must be set for PARAMETER_MULTI")
	}
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
	}

	return props.properties.GeneratePhysicalResourceId(), map[string]interface{}{
		"Prefix": props.properties.Target,
		"Count":  len(outData),
	}, nil
}

func handleParameter(props BaseProps) (physicalResourceID string, data map[string]interface{}, err error) {
	outData, outDataErr := props.secretDecryptedData.GetRaw()
	if outDataErr != nil {
		return props.properties.GeneratePhysicalResourceId(), nil, outDataErr
	}
	if props.properties.EncryptionKey == nil {
		return "", nil, fmt.Errorf("EncryptionKey must be set for PARAMETER")
	}
	putParameterResponse, putParameterResponseErr := props.clients.SsmPutParameter(*&props.properties.Target, outData, *props.properties.EncryptionKey)
	if putParameterResponseErr != nil {
		return props.properties.GeneratePhysicalResourceId(), nil, fmt.Errorf("failed to update ssm parameter:\n%v", err)
	}
	return props.properties.GeneratePhysicalResourceId(), map[string]interface{}{
		"ParameterName": props.properties.Target,
		"Version":       putParameterResponse.Version,
		"Tier":          putParameterResponse.Tier,
	}, nil
}
