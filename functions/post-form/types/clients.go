package types

import (
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/ses"
)

// Clients is the dunctions exsternal resources
type Clients struct {
	Ses    *ses.SES
	Dynamo *dynamodb.DynamoDB
}