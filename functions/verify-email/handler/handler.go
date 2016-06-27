package handler

import (
	"errors"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/cluda/cluda-form/functions/post-form/util"
)

// Handle will handel a new event
func Handle(e Event, conf Config, cli Clients) (interface{}, error) {

	resp, err := cli.Dynamo.GetItem(util.FormDataRequest(conf.FormFreeTable, e.Receiver))

	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	if len(resp.Item) == 0 {
		log.Println("no form found with key:", e.Receiver)
		return "", errors.New("no form found")
	} else if *resp.Item["verifyed"].BOOL {
		log.Println("form with key:", e.Receiver, "already verified")
		return "already verified", nil
	} else if *resp.Item["secret"].S == e.Secret {
		cli.Dynamo.UpdateItem(verifyFormDynamoIn(conf.FormFreeTable, e.Receiver))
		return "receiver verifyed", nil
	}
	log.Println("error or wrong secret for ", e.Receiver ,". Correct secret:", *resp.Item["secret"].S,  ", this secret:", e.Secret)
	return "", errors.New("unknown error")
}

func verifyFormDynamoIn(table string, receiver string) *dynamodb.UpdateItemInput {
	return &dynamodb.UpdateItemInput{
		Key: map[string]*dynamodb.AttributeValue{ // Required
			"id": { // Required
				S: aws.String(receiver),
			},
		},
		TableName: aws.String(table), // Required
		AttributeUpdates: map[string]*dynamodb.AttributeValueUpdate{
			"verifyed": { // Required
				Action: aws.String("PUT"),
				Value: &dynamodb.AttributeValue{
					BOOL: aws.Bool(true),
				},
			},
		},
	}
}
