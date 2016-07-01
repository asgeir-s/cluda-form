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
const secret = "test-secret"

var firstID = util.RandString(10)

var dynamo = dynamodb.New(session.New(), &aws.Config{Region: aws.String("us-west-2")})

func TestAddNewPayedForm(t *testing.T) {

	form := types.Form{
		ID:          firstID, // range key
		Email:       email,   // secoundary sort/range key
		Origin:      origin,  // primary key
		Secret:      secret,  // the user and I kows this secret
		Verified:    false,
		Subscribing: false,
	}

	err := database.AddNewPayedForm(dynamo, formTable, form)
	if err != nil {
		t.Fatal("Could not add form to database. Error:", err)
	}

	err = database.AddNewPayedForm(dynamo, formTable, form)
	if err != nil {
		if !strings.Contains(err.Error(), "ConditionalCheckFailedException") {
			t.Fatal("ConditionalCheckFailedException was not returned")
		}
	} else {
		t.Fatal("should not be possible to add the same form again")
	}

	form.Origin = "other.com"
	err = database.AddNewPayedForm(dynamo, formTable, form)
	if err != nil {
		t.Fatal("Could not add form with other origin but same id to database. Error:", err)
	}

	form.ID = util.RandString(10)
	err = database.AddNewPayedForm(dynamo, formTable, form)
	if err != nil {
		t.Fatal("Could not add form with same origin but other id to database. Error:", err)
	}
}

func TestAddNewFreeForm(t *testing.T) {

	form := types.Form{
		ID:          firstID, // sort/range key
		Email:       email,   // secoundary sort/range key
		Origin:      origin,  // primary key
		Secret:      secret,  // the user and I kows this secret
		Verified:    false,
		Subscribing: false,
	}

	form.ID = util.RandString(10)

	err := database.AddNewFreeForm(dynamo, formTable, form)
	if err != nil {
		if !strings.Contains(err.Error(), "already exists") {
			t.Fatal("'already exists' was not returned")
		}
	} else {
		t.Fatal("Should not be possible to add a free form with an already used origin/email combination")
	}

	form.ID = util.RandString(10)
	form.Origin = util.RandString(20) + ".com"
	err = database.AddNewFreeForm(dynamo, formTable, form)
	if err != nil {
		t.Fatal("Could not add form with new origin/email combination. Error:", err)
	}
}

func TestGetFormByOriginId(t *testing.T) {
	resp, err := database.GetFormByOriginID(dynamo, formTable, origin, firstID)
	if err != nil {
		t.Fatal("Could not get the newly added form. Error:", err)
	}

	if resp.Email != email || resp.ID != firstID || resp.Origin != origin || resp.Secret != secret {
		t.Error("The returned form is not equal to the added form.")
	}
}

func TestGetFormByOriginEmail(t *testing.T) {
	resp, err := database.GetFormByOriginEmail(dynamo, formTable, origin, email)
	if err != nil {
		t.Fatal("Could not get the newly added form. Error:", err)
	}

	if resp.Email != email || resp.Origin != origin {
		t.Error("The returned form is not equal to the requested form.")
	}
}
