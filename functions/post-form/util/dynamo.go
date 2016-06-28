package util

import (
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// FormDataRequest returns the "get item"-input for dynamo
func FormDataRequest(tableName string, userID string) *dynamodb.GetItemInput {
	return &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(userID),
			},
		},
		TableName: aws.String(tableName),
	}
}

func NewFormDataPut(tableName string, key string, randomSecret string) *dynamodb.PutItemInput {
	return &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(key),
			},
			"secret": {
				S: aws.String(randomSecret),
			},
			"verifyed": {
				BOOL: aws.Bool(false),
			},
		},
		TableName:           aws.String(tableName),
		ConditionExpression: aws.String("attribute_not_exists(id)"),
	}
}

func FormSubmissionPut(tableName string, userID string, timestamp int64, data string) *dynamodb.PutItemInput {

	return &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"userId": {
				S: aws.String(userID),
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
}
