package database

import (
	"errors"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"

	"github.com/cluda/cluda-form/functions/post-form/types"
)

// GetForm returns the form corresponding to the origin and key. The key can be formID or email
func GetForm(dynamo *dynamodb.DynamoDB, tableName, origin, key string) (types.Form, error) {
	if strings.Contains(key, "@") {
		return GetFormByOriginEmail(dynamo, tableName, origin, key)
	}
	return GetFormByOriginID(dynamo, tableName, origin, key)
}

// GetFormByOriginID returns a form using the origin/id combination (primary keys) if it exists
func GetFormByOriginID(dynamo *dynamodb.DynamoDB, tableName, origin, id string) (types.Form, error) {
	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"origin": {
				S: aws.String(origin),
			},
			"id": {
				S: aws.String(id),
			},
		},
		TableName: aws.String(tableName),
	}

	itemRes, err := dynamo.GetItem(input)
	if err != nil {
		return types.Form{}, err
	}

	if len(itemRes.Item) == 0 {
		return types.Form{}, errors.New("[does not exist] no form with this origin/id combination")
	}

	form := types.Form{
		ID:          id,                       // sort/range key
		Email:       *itemRes.Item["email"].S, // secoundary sort/range key
		Origin:      origin,                   // primary key
		Verified:    *itemRes.Item["verified"].BOOL,
		Subscribing: *itemRes.Item["subscription"].BOOL,
		Secret:      *itemRes.Item["secret"].S,
	}
	return form, nil
}

// GetFormByOriginEmail returns a form using the origin/email combination (local secondary index) if it exists
func GetFormByOriginEmail(dynamo *dynamodb.DynamoDB, tableName, origin, email string) (types.Form, error) {
	params := &dynamodb.QueryInput{
		TableName:              aws.String(tableName),
		IndexName:              aws.String("origin-email-index"),
		KeyConditionExpression: aws.String("origin = :originval AND email = :emailval"),
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":originval": {S: &origin},
			":emailval":  {S: &email},
		},
		Limit: aws.Int64(1),
	}

	itemRes, err := dynamo.Query(params)
	if err != nil {
		return types.Form{}, err
	}
	if len(itemRes.Items) == 0 {
		return types.Form{}, errors.New("[does not exist] no form with this origin/email combination")
	}

	form := types.Form{
		ID:          *itemRes.Items[0]["id"].S, // sort/range key
		Email:       email,                     // secoundary sort/range key
		Origin:      origin,                    // primary key
		Verified:    *itemRes.Items[0]["verified"].BOOL,
		Subscribing: *itemRes.Items[0]["subscription"].BOOL,
		Secret:      *itemRes.Items[0]["secret"].S,
	}
	return form, nil
}

// AddNewPayedForm adds a new form if no form with this origin/id (primary key) exists
func AddNewPayedForm(dynamo *dynamodb.DynamoDB, tableName string, form types.Form) error {
	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"origin": {
				S: aws.String(form.Origin),
			},
			"id": {
				S: aws.String(form.ID),
			},
			"email": {
				S: aws.String(form.Email),
			},
			"secret": {
				S: aws.String(form.Secret),
			},
			"verified": {
				BOOL: aws.Bool(form.Verified),
			},
			"subscription": {
				BOOL: aws.Bool(form.Subscribing),
			},
		},
		TableName:           aws.String(tableName),
		ConditionExpression: aws.String("attribute_not_exists(origin)"),
	}
	_, err := dynamo.PutItem(input)
	return err
}

// AddNewFreeForm adds a new form if no form with this origin/email (secondary index) exists
func AddNewFreeForm(dynamo *dynamodb.DynamoDB, tableName string, form types.Form) error {
	_, err := GetFormByOriginEmail(dynamo, tableName, form.Origin, form.Email)
	if err == nil {
		return errors.New("A form with this origin/email combination already exists.")
	}

	input := &dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue{
			"origin": {
				S: aws.String(form.Origin),
			},
			"id": {
				S: aws.String(form.ID),
			},
			"email": {
				S: aws.String(form.Email),
			},
			"secret": {
				S: aws.String(form.Secret),
			},
			"verified": {
				BOOL: aws.Bool(form.Verified),
			},
			"subscription": {
				BOOL: aws.Bool(form.Subscribing),
			},
		},
		TableName:           aws.String(tableName),
		ConditionExpression: aws.String("attribute_not_exists(origin)"), // form with this hash and range does not exist AND combination origin/email does not exist
	}
	_, err = dynamo.PutItem(input)
	return err
}
