package main

import (
	"encoding/json"
	"log"

	"github.com/apex/go-apex"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/caarlos0/env"

	"github.com/cluda/cluda-form/functions/post-form/handler"
	"github.com/cluda/cluda-form/functions/post-form/types"
)

func main() {

	config := types.Config{}
	env.Parse(&config)

	clients := types.Clients{
		Ses:    ses.New(session.New(), &aws.Config{Region: aws.String(config.AwsRegion)}),
		Dynamo: dynamodb.New(session.New(), &aws.Config{Region: aws.String(config.AwsRegion)}),
	}

	apex.HandleFunc(func(rawEvent json.RawMessage, ctx *apex.Context) (interface{}, error) {
		var event types.Event

		err := json.Unmarshal(rawEvent, &event)
		if err != nil {
			return nil, err
		}

		res, err := handler.Handle(event, config, clients)
		if err != nil {
			log.Println(err.Error())
			return nil, err // TODO: should not return the error.
		}
		return res, nil
	})
}
