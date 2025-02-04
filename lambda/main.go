package main

import (
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/markussiebert/cdk-sops-secrets/internal/aws"
)

type MyEvent struct {
	Name string `json:"name"`
}

func HandleRequestWithClients(clients *aws.Client, event MyEvent) error {

}

func HandleRequest(event MyEvent) error {

	clients := aws.CreateAwsClients()

	return HandleRequestWithClients(clients, event)
}

func main() {
	lambda.Start(HandleRequest)
}
