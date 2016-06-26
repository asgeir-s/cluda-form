package util_test

import (
	"fmt"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/ses"
	"github.com/cluda/cluda-form/functions/post-form/util"
)

const secret = "test-secret"

var sesClient = ses.New(session.New(), &aws.Config{Region: aws.String("us-west-2")})

func TestSendEmail(t *testing.T) {
	templateData := util.EmailData{
		Text1:  "",
		Text2:  "To activate your form, please confirm your email address by clicking the link below.",
		Button: "Confirm email address",
		Secret: secret,
	}

	body, err := util.ParseTemplate("email-templates/action.html", templateData)
	if err != nil {
		t.Error(err)

	}

	resp, err := sesClient.SendEmail(util.SendEmialInput("test-in-1@coinsignals.com", "sogasg@gmail.com", "Test 22", body))
	if err != nil {
		t.Error(err)
	}

	fmt.Println("res: ", resp)
}
