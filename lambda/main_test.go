package main

import (
	"context"
	"testing"

	"github.com/aws/aws-lambda-go/cfn"
	"github.com/markussiebert/cdk-sops-secrets/internal/client"
	"github.com/stretchr/testify/assert"
)

func TestHandleRequestWithClients_CreateRequest(t *testing.T) {
	mockClients := client.MockClient{}
	mockEvent := cfn.Event{
		RequestType: cfn.RequestCreate,
	}

	physicalResourceID, data, err := HandleRequestWithClients(&mockClients, mockEvent)

	assert.NoError(t, err)
	assert.NotEmpty(t, physicalResourceID)
	assert.NotNil(t, data)
}

func TestHandleRequestWithClients_DeleteRequest(t *testing.T) {
	mockClients := client.MockClient{}
	mockEvent := cfn.Event{
		RequestType: cfn.RequestDelete,
	}

	physicalResourceID, data, err := HandleRequestWithClients(&mockClients, mockEvent)

	assert.NoError(t, err)
	assert.Empty(t, physicalResourceID)
	assert.Nil(t, data)
}

func TestHandleRequestWithClients_UnsupportedRequest(t *testing.T) {
	mockClients := client.MockClient{}
	mockEvent := cfn.Event{
		RequestType: "UnsupportedRequestType",
	}

	physicalResourceID, data, err := HandleRequestWithClients(&mockClients, mockEvent)

	assert.Error(t, err)
	assert.Empty(t, physicalResourceID)
	assert.Nil(t, data)
}

func TestHandleRequest(t *testing.T) {
	mockEvent := cfn.Event{
		RequestType: cfn.RequestCreate,
	}

	physicalResourceID, data, err := HandleRequest(context.Background(), mockEvent)

	assert.NoError(t, err)
	assert.NotEmpty(t, physicalResourceID)
	assert.NotNil(t, data)
}

func TestGeneratePhysicalResourceId(t *testing.T) {
	input := "arn:aws:secretsmanager:region:account-id:secret:example"
	expected := "arn:custom:sopssync:region:account-id:secret:example"

	result := generatePhysicalResourceId(input)

	assert.Equal(t, expected, result)
}
