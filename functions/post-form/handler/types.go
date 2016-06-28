package handler

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/ses"
)

// Event is the event message
type Event struct {
	Receiver string `json:"receiver"` // email or userId
	Data     string `json:"data"`
}

// Config is the functions config
type Config struct {
	AwsRegion       string `env:"AWS_REGION" envDefault:"us-west-2"`
	FormFreeTable   string `env:"FORM_FREE_TABLE" envDefault:"test-form-table"`
	EmailFromAddres string `env:"EMAIL_FROM_ADDRES" envDefault:"test-in-1@coinsignals.com"`
	SubmissionTable string `env:"SUMBMISSION_TABLE" envDefault:"test-submission-table"`
	BaseURL         string `env:"BASE_URL" envDefault:"https://fqcx5ghvc3.execute-api.us-west-2.amazonaws.com/prod"`
}

// Clients is the dunctions exsternal resources
type Clients struct {
	Ses    *ses.SES
	Dynamo *dynamodb.DynamoDB
}
