package handler

import (
	"errors"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/cluda/cluda-form/functions/post-form/database"
)

// Handle will handel a new event
func Handle(e Event, conf Config, cli Clients) (interface{}, error) {

	// get form
	form, err := database.GetForm(cli.Dynamo, conf.FormTable, e.Origin, e.Receiver)
	if err != nil {
		return nil, err
	}
	if form.Verified {
		log.Println("form with origin:", e.Origin, "and receiver:", e.Receiver, "already verified")
		return "already verified", nil
	} else if form.Secret == e.Secret && form.Origin == e.Origin && (e.Receiver == form.Email || e.Receiver == form.ID) {
		// valide verification
		return "receiver verified", verifyForm(cli.Dynamo, conf.FormTable, e.Origin, form.ID)
	}
	// if we get here their is some error with the verifcation (or hacking :P)
	return "", errors.New("error or wrong secret or origin for " + e.Receiver + ". Correct secret: " + form.Secret +
		", this secret: " + e.Secret + ". Correct origin: " + form.Origin + ", this origin: " + e.Origin)
}

func verifyForm(dynamo *dynamodb.DynamoDB, table string, origin, id string) error {
	param := &dynamodb.UpdateItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"origin": {
				S: aws.String(origin),
			},
			"id": {
				S: aws.String(id),
			},
		},
		TableName: aws.String(table),
		AttributeUpdates: map[string]*dynamodb.AttributeValueUpdate{
			"verified": { // Required
				Action: aws.String("PUT"),
				Value: &dynamodb.AttributeValue{
					BOOL: aws.Bool(true),
				},
			},
		},
	}
	_, err := dynamo.UpdateItem(param)
	return err
}
