package database

import (
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/cluda/cluda-form/functions/post-form/types"
)

func GetFormByOriginId(dynamo *dynamodb.DynamoDB, tableName, origin, id string) (types.Form, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"origin": {
				S: aws.String(origin),
			},
			"id": {
				S: aws.String(id),
			},
		},
		TableName: aws.String(tableName),
	}

	itemRes, err := dynamo.GetItem(input)
	if err != nil {
		return types.Form{}, err
	}

	form := types.Form{
		ID:          id,                       // sort/range key
		Email:       *itemRes.Item["email"].S, // secoundary sort/range key
		Origin:      origin,                   // primary key
		Verifyed:    *itemRes.Item["verifyed"].BOOL,
		Subscribing: *itemRes.Item["subscribing"].BOOL,
	}
	return form, nil
}

func GetFormByOriginEmail(dynamo *dynamodb.DynamoDB, tableName, origin, email string) (types.Form, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"origin": {
				S: aws.String(origin),
			},
			"email": {
				S: aws.String(email),
			},
		},
		TableName: aws.String(tableName),
	}

	itemRes, err := dynamo.GetItem(input)
	if err != nil {
		return types.Form{}, err
	}

	form := types.Form{
		ID:          *itemRes.Item["id"].S, // sort/range key
		Email:       email,                 // secoundary sort/range key
		Origin:      origin,                // primary key
		Verifyed:    *itemRes.Item["verifyed"].BOOL,
		Subscribing: *itemRes.Item["subscribing"].BOOL,
	}
	return form, nil
}

func AddNewForm(dynamo *dynamodb.DynamoDB, tableName string, form types.Form) error {
	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"origin": {
				S: aws.String(form.Origin),
			},
			"id": {
				S: aws.String(form.ID),
			},
			"email": {
				S: aws.String(form.Email),
			},
			"secret": {
				S: aws.String(form.Secret),
			},
			"verifyed": {
				BOOL: aws.Bool(form.Verifyed),
			},
			"subscription": {
				BOOL: aws.Bool(form.Subscribing),
			},
		},
		TableName:           aws.String(tableName),
		Expected: map[string]*dynamodb.ExpectedAttributeValue{
			"origin": {
				Exists: aws.Bool(false),
			},
			"id": {
				Exists: aws.Bool(false),
			},
		},
	}
	_, err := dynamo.PutItem(input)
	return err
}

func FormSubmissionPut(dynamo *dynamodb.DynamoDB, tableName string, formID string, timestamp int64, data string) error {
	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"formId": {
				S: aws.String(formID),
			},
			"timestamp": {
				N: aws.String(strconv.FormatInt(timestamp, 10)),
			},
			"data": {
				S: aws.String(data),
			},
		},
		TableName:           aws.String(tableName),
		ConditionExpression: aws.String("attribute_not_exists(id)"),
	}
	_, err := dynamo.PutItem(input)
	return err
}
