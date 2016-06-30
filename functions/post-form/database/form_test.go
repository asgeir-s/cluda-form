package database_test

import (
	"strings"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/cluda/cluda-form/functions/post-form/database"
	"github.com/cluda/cluda-form/functions/post-form/types"
	"github.com/cluda/cluda-form/functions/post-form/util"
)

const formTable = "test-form-table"
const email = "sogasg@msn.com"
const origin = "test.com"
const secret = "test.secret"

var dynamo = dynamodb.New(session.New(), &aws.Config{Region: aws.String("us-west-2")})

func TestAddNewForm(t *testing.T) {
	id := util.RandString(10)

	form := types.Form{
		ID:          id,     // sort/range key
		Email:       email,  // secoundary sort/range key
		Origin:      origin, // primary key
		Secret:      secret, // the user and I kows this secret
		Verifyed:    false,
		Subscribing: false,
	}

	err := database.AddNewForm(dynamo, formTable, form)
	if err != nil {
		t.Fatal("Could not add form to database. Error:", err)
	}

	err = database.AddNewForm(dynamo, formTable, form)
	if err != nil {
		strings.Contains(err.Error(), "ConditionalCheckFailedException")
	} else {
		t.Fatal("should not be possible to add the same form again")
	}

	form.Origin = "other.com"
	err = database.AddNewForm(dynamo, formTable, form)
	if err != nil {
		t.Fatal("Could not add form with other origin but same id to database. Error:", err)
	}

	form.ID = util.RandString(10)
	err = database.AddNewForm(dynamo, formTable, form)
	if err != nil {
		t.Fatal("Could not add form with same origin but other id to database. Error:", err)
	}
}

/*
func TestFormDataRequest(t *testing.T) {
	resp, err := dynamo.GetItem(util.FormDataRequest(fromTable, origin, email))

	if err != nil {
		t.Error(err)
		fmt.Println(err.Error())
	}

	if *resp.Item["id"].S != email {
		t.Error("wring responds")
	}
}
*/
