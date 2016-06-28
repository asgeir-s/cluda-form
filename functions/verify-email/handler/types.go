package handler

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// Event is the event message
type Event struct {
	Receiver string `json:"receiver"` // email or userId
	Secret   string `json:"secret"`
}

// Config is the functions config
type Config struct {
	AwsRegion       string `env:"AWS_REGION" envDefault:"us-west-2"`
	FormFreeTable   string `env:"FORM_FREE_TABLE" envDefault:"test-form-table"`
}

// Clients is the dunctions exsternal resources
type Clients struct {
	Dynamo *dynamodb.DynamoDB
}
