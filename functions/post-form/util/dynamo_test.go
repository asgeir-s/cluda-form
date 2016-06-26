package util_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/cluda/cluda-form/functions/post-form/util"
)

const fromTable = "test-form-table"
const formItemID = "sogasg@msn.com"

var dynamo = dynamodb.New(session.New(), &aws.Config{Region: aws.String("us-west-2")})

func TestFormDataPut(t *testing.T) {
	_, err := dynamo.PutItem(util.NewFormDataPut(fromTable, formItemID, util.RandString(10)))

	if err != nil {
		strings.Contains(err.Error(), "ConditionalCheckFailedException")
	} else {
		t.Error("ConditionalCheckFailedException did not work")
	}
}

func TestFormDataRequest(t *testing.T) {
	resp, err := dynamo.GetItem(util.FormDataRequest(fromTable, formItemID))

	if err != nil {
		t.Error(err)
		fmt.Println(err.Error())
	}

	if *resp.Item["id"].S != formItemID {
		t.Error("wring responds")
	}
}
