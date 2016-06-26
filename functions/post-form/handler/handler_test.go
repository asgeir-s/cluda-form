package handler_test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/cluda/cluda-form/functions/post-form/handler"
)

var config = handler.Config{
	AwsRegion:       "us-west-2",
	FormFreeTable:   "test-form-table",
	EmailFromAddres: "test-in-1@coinsignals.com",
}

var clients = handler.Clients{
	Ses:    ses.New(session.New(), &aws.Config{Region: aws.String(config.AwsRegion)}),
	Dynamo: dynamodb.New(session.New(), &aws.Config{Region: aws.String(config.AwsRegion)}),
}

func TestHandler(t *testing.T) {
	event := handler.Event{
		Receiver: "sogasg@gmail.com",
		Data:     "data=some",
	}

	res, err := handler.Handle(event, config, clients)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println("handler res:", res)
	}
}
