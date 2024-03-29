package main

import (
	"encoding/json"
	"log"

	"github.com/apex/go-apex"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/cluda/cluda-form/functions/verify-email/handler"

	"github.com/caarlos0/env"
)

func main() {

	config := handler.Config{}
	env.Parse(&config)

	clients := handler.Clients{
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
			log.Println(err.Error())
			return nil, err // TODO: should not return the error.
		}
		return res, nil
	})
}
