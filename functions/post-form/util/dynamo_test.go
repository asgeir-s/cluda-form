package util_test

import (
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/cluda/cluda-form/functions/post-form/util"

	"fmt"
	"testing"
)

const fromTable = "test-form-table"
const formItemID = "sogasg@msn.com"

func TestFormDataRequest(t *testing.T) {

	dynamo := dynamodb.New(session.New())

	resp, err := dynamo.GetItem(util.FormDataRequest(fromTable, formItemID))

	if err != nil {
		t.Error(err)
		fmt.Println(err.Error())
	}

	fmt.Println(resp)
}
