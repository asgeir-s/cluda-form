package util

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ses"
)

func SendEmialInput(sender, resepian, subject, body string) *ses.SendEmailInput {
	return &ses.SendEmailInput{
		Destination: &ses.Destination{
			ToAddresses: []*string{
				aws.String(resepian),
			},
		},
		Message: &ses.Message{
			Body: &ses.Body{
				Html: &ses.Content{
					Data:    aws.String(body),
					Charset: aws.String("UTF-8"),
				},
			},
			Subject: &ses.Content{
				Data:    aws.String(subject),
				Charset: aws.String("UTF-8"),
			},
		},
		Source: aws.String(sender),
	}
}
