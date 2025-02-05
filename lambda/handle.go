package main

import (
	"fmt"
	"log"

	"github.com/markussiebert/cdk-sops-secrets/internal/client"
	"github.com/markussiebert/cdk-sops-secrets/internal/data"
	"github.com/markussiebert/cdk-sops-secrets/internal/sops"
)

type BaseProps struct {
	clients             client.AwsClient
	tempArn             string
	secretDecryptedData *data.Data
}

type HandleSecretProps struct {
	BaseProps
	outputFormat sops.Format
	hash         string
	secretArn    *string
}

func handleSecret(props HandleSecretProps) (physicalResourceID string, data map[string]interface{}, err error) {
	outData, outDataErr := produceOutput(props.outputFormat, props.secretDecryptedData)
	if outDataErr != nil {
		return props.tempArn, nil, outDataErr
	}
	if props.secretArn == nil {
		return props.tempArn, nil, fmt.Errorf("SecretARN must be set for SECRET")
	}

	putSecretValueResp, putSecretValueRespErr := props.clients.SecretsManagerPutSecretValue(props.hash, *props.secretArn, outData)
	if putSecretValueRespErr != nil {
		return props.tempArn, nil, putSecretValueRespErr
	}
	return *putSecretValueResp.ARN, map[string]interface{}{
		"ARN":           *putSecretValueResp.ARN,
		"Name":          *putSecretValueResp.Name,
		"VersionStages": putSecretValueResp.VersionStages,
		"VersionId":     *putSecretValueResp.VersionId,
	}, nil
}

type HandleParameterMultiProps struct {
	BaseProps
	parameterKeyPrefix *string
	encryptionKey      *string
}

func handleParameterMulti(props HandleParameterMultiProps) (physicalResourceID string, data map[string]interface{}, err error) {
	if props.parameterKeyPrefix == nil || props.encryptionKey == nil {
		return "", nil, fmt.Errorf("ParameterKeyPrefix and EncryptionKey must be set for PARAMETER_MULTI")
	}
	log.Printf("Patching multiple string parameters")
	v, e := props.secretDecryptedData.ToStringMap()
	if e != nil {
		return props.tempArn, nil, e
	}
	for key, value := range v {
		_, err := props.clients.SsmPutParameter(fmt.Sprintf("%s%s", *props.parameterKeyPrefix, key), &value, *props.encryptionKey)
		if err != nil {
			return props.tempArn, nil, fmt.Errorf("failed to update ssm parameter:\n%v", err)
		}
	}

	return props.tempArn, map[string]interface{}{
		"Prefix": props.parameterKeyPrefix,
		"Count":  len(v),
	}, nil
}

type HandleParameterProps struct {
	BaseProps
	outputFormat  sops.Format
	parameterName *string
	encryptionKey *string
}

func handleParameter(props HandleParameterProps) (physicalResourceID string, data map[string]interface{}, err error) {
	outData, outDataErr := produceOutput(props.outputFormat, props.secretDecryptedData)
	if outDataErr != nil {
		return props.tempArn, nil, outDataErr
	}
	if props.parameterName == nil || props.encryptionKey == nil {
		return "", nil, fmt.Errorf("ParameterName and EncryptionKey must be set for PARAMETER")
	}
	putParameterResponse, putParameterResponseErr := props.clients.SsmPutParameter(*props.parameterName, outData, *props.encryptionKey)
	if putParameterResponseErr != nil {
		return props.tempArn, nil, fmt.Errorf("failed to update ssm parameter:\n%v", err)
	}
	return props.tempArn, map[string]interface{}{
		"ParameterName": props.parameterName,
		"Version":       putParameterResponse.Version,
		"Tier":          putParameterResponse.Tier,
	}, nil
}

func produceOutput(outputFormat sops.Format, secretDecryptedData *data.Data) (*[]byte, error) {
	switch outputFormat {
	case sops.JSON:
		return secretDecryptedData.ToJSON()
	case sops.YAML:
		return secretDecryptedData.ToYAML()
	case sops.DOTENV:
		return secretDecryptedData.ToDotEnv()
	case sops.BINARY:
		return secretDecryptedData.GetRaw()
	default:
		return nil, fmt.Errorf("unsupported output format %s", outputFormat)
	}
}
