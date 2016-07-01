package handler_test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/cluda/cluda-form/functions/post-form/handler"
	"github.com/cluda/cluda-form/functions/post-form/types"
)

var config = types.Config{
	AwsRegion:       "us-west-2",
	FormTable:       "test-form-table",
	EmailFromAddres: "test-in-1@coinsignals.com",
	SubmissionTable: "test-submission-table",
	BaseURL:         "https://fqcx5ghvc3.execute-api.us-west-2.amazonaws.com/prod",
}

var clients = types.Clients{
	Ses:    ses.New(session.New(), &aws.Config{Region: aws.String(config.AwsRegion)}),
	Dynamo: dynamodb.New(session.New(), &aws.Config{Region: aws.String(config.AwsRegion)}),
}

func TestHandler(t *testing.T) {
	event := types.Event{
		Receiver: "sogasg@gmail.com",
		Data:     "data1=some&more=someOtherdata",
		Origin:   "example.com",
	}

	res, err := handler.Handle(event, config, clients)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println("handler res:", res)
	}
}
