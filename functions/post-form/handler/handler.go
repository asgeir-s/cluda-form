package handler

import (
	"errors"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/cluda/cluda-form/functions/post-form/database"
	"github.com/cluda/cluda-form/functions/post-form/types"
	"github.com/cluda/cluda-form/functions/post-form/util"
)

// Handle will handel a new event
func Handle(e types.Event, conf types.Config, cli types.Clients) (interface{}, error) {
	form, err := database.GetForm(cli.Dynamo, conf.FormTable, e.Origin, e.Receiver)
	if err != nil {
		if strings.Contains(err.Error(), "[does not exist]") {
			// new free form
			if !strings.Contains(e.Receiver, "@") {
				return nil, errors.New("encountered a new form that was not using an email as receiver. " +
					"Forms with ID as receiver should already have be added to the forms table.")
			}
			form := types.Form{
				ID:          util.RandString(12), // range key
				Email:       e.Receiver,          // secoundary sort/range key
				Origin:      e.Origin,            // primary key
				Secret:      util.RandString(12), // the user and I kows this secret
				Verified:    false,
				Subscribing: false,
			}

			// add form
			err := database.AddNewFreeForm(cli.Dynamo, conf.FormTable, form)
			if err != nil {
				return nil, err
			}

			// send verification email
			templateData := util.EmailData{
				Text1:  "New form for " + form.Origin + ".",
				Text2:  "To activate your form, please confirm your email address by clicking the link below.",
				Button: "Confirm email address",
				Url:    conf.BaseURL + "/verify?receiver=" + form.Email + "&secret=" + form.Secret + "&origin=" + form.Origin,
			}

			body, err := util.ParseTemplate("email-templates/action.html", templateData)
			if err != nil {
				log.Println(err.Error())
				return nil, err
			}

			_, err = cli.Ses.SendEmail(util.SendEmialInput(conf.EmailFromAddres, e.Receiver, "Activate your new form at "+form.Origin, body))
			if err != nil {
				log.Println(err.Error())
				return nil, err
			}
			return "verification email sent", nil

		}
		return nil, err
	}

	if !form.Verified {
		return nil, errors.New("form is not verified")
	}

	// add to submission table
	sumbmission := types.Submission{
		FormID:     form.ID,               // primary key (part 2)
		FormOrigin: form.Origin,           // primary key (part 1)
		Timestamp:  time.Now().UnixNano(), // range key
		Data:       e.Data,
	}

	err = database.AddFormSubmission(cli.Dynamo, conf.SubmissionTable, sumbmission)
	if err != nil {
		return nil, err
	}

	// send submission to associated email
	data, err := url.ParseQuery(e.Data)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	body := util.CreateEmailBody(data)

	_, err = cli.Ses.SendEmail(util.SendEmialInput(conf.EmailFromAddres, "sogasg@gmail.com", "New form submitted", body))
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return "submission handled", nil
}
