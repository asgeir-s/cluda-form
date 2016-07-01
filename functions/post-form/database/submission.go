package database

import (
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/cluda/cluda-form/functions/post-form/types"
)

// AddFormSubmission adds a form submission to the table with hash: originId (formOrigin + formId), range: timestamp (time.Now().UnixNano())
func AddFormSubmission(dynamo *dynamodb.DynamoDB, tableName string, submission types.Submission) error {
	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"origin_and_id": {
				S: aws.String(submission.FormOrigin + "-" + submission.FormID),
			},
			"timestamp": {
				N: aws.String(strconv.FormatInt(submission.Timestamp, 10)),
			},
			"data": {
				S: aws.String(submission.Data),
			},
		},
		TableName:           aws.String(tableName),
		ConditionExpression: aws.String("attribute_not_exists(origin_and_id)"),
	}
	_, err := dynamo.PutItem(input)
	return err
}
