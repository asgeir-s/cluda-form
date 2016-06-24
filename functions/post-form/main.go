package main

import (
	"encoding/json"
	"fmt"

	"github.com/apex/go-apex"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/cluda/cluda-form/functions/post-form/util"

	"github.com/caarlos0/env"
)

type message struct {
	Receiver string `json:"receiver"` // email or formId
	Data     string `json:"data"`
}

type config struct {
	AwsRegion     string `env:"AWS_REGION" envDefault:"us-west-2"`
	FormFreeTable string `env:"FORM_FREE_TABLE" envDefault:"cluda-form-free-table"`
}

func main() {

	fmt.Println("in main")

	config := config{}
	env.Parse(&config)

	sesClient := ses.New(session.New(), &aws.Config{Region: aws.String(config.AwsRegion)})
	dynamo := dynamodb.New(session.New(), &aws.Config{Region: aws.String(config.AwsRegion)})

	apex.HandleFunc(func(event json.RawMessage, ctx *apex.Context) (interface{}, error) {
		var m message

		// one form object for free(email) and one for payable (id)
		// get form object from database. Primary key: formID, secoundaryIndex: email
		resp, err := dynamo.GetItem(util.FormDataRequest(config.FormFreeTable, m.Receiver))

		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}

		fmt.Println(resp)

		// add to submission table

		// send submission to asosiated email

		fmt.Println(sesClient.APIVersion)

		if err := json.Unmarshal(event, &m); err != nil {
			return nil, err
		}

		return m.Receiver, nil
	})
}
