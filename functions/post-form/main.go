package main

import (
	"encoding/json"

	"github.com/apex/go-apex"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/cluda/cluda-form/functions/post-form/handler"

	"github.com/caarlos0/env"
)

func main() {

	config := handler.Config{}
	env.Parse(&config)

	clients := handler.Clients{
		Ses:    ses.New(session.New(), &aws.Config{Region: aws.String(config.AwsRegion)}),
		Dynamo: dynamodb.New(session.New(), &aws.Config{Region: aws.String(config.AwsRegion)}),
	}

	apex.HandleFunc(func(rawEvent json.RawMessage, ctx *apex.Context) (interface{}, error) {
		var event handler.Event

		err := json.Unmarshal(rawEvent, &event)
		if err != nil {
			return nil, err
		}

		res, err := handler.Handle(event, config, clients)
		if err != nil {
			return nil, err
		}
		return res, nil
	})
}
