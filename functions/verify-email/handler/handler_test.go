package handler_test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/cluda/cluda-form/functions/verify-email/handler"
)

var config = handler.Config{
	AwsRegion: "us-west-2",
	FormTable: "test-form-table",
}

var clients = handler.Clients{
	Dynamo: dynamodb.New(session.New(), &aws.Config{Region: aws.String(config.AwsRegion)}),
}

func TestHandler(t *testing.T) {
	event := handler.Event{
		Receiver: "sogasg@gmail.com",
		Secret:   "DPnssIYBRXYQ",
		Origin:   "example.com",
	}

	res, err := handler.Handle(event, config, clients)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println("handler res:", res)
	}
}
