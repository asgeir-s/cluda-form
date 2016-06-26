package util

import (
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// FormDataRequest returns the "get item"-input for dynamo
func FormDataRequest(tableName string, formID string) *dynamodb.GetItemInput {
	return &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				S: aws.String(formID),
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

func FormSubmissionPut(tableName string, formId string, timestamp int64, data map[string][]string) *dynamodb.PutItemInput {

	dataMap := make(map[string]*dynamodb.AttributeValue)
	for key, value := range data {
		println("key:", key, "value:", value[0])
		dataMap[key] = &dynamodb.AttributeValue{
			SS: aws.StringSlice(value),
		}
		//dataMap[key].SS = aws.StringSlice(value)
	}

	return &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"formId": {
				S: aws.String(formId),
			},
			"timestamp": {
				N: aws.String(strconv.FormatInt(timestamp, 10)),
			},
			"data": {
				M: dataMap,
			},
		},
		TableName:           aws.String(tableName),
		ConditionExpression: aws.String("attribute_not_exists(id)"),
	}
}
