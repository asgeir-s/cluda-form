package util

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// FormDataRequest returns the "get item"-input for dynamo 
func FormDataRequest(tableName string, key string) *dynamodb.GetItemInput {
	params := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(key),
			},
		},
		TableName: aws.String(tableName),
	}
	return params
}
