package util

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// FormDataRequest returns the "get item"-input for dynamo
func FormDataRequest(tableName string, key string) *dynamodb.GetItemInput {
	return &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(key),
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
