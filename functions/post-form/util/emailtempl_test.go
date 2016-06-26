package util_test

import (
	"fmt"
	"testing"

	"github.com/cluda/cluda-form/functions/post-form/util"
)

func TestEmailTemplate(t *testing.T) {

	// send confirm email
	templateData := util.EmailData{
		Text1:  "YoYo",
		Text2:  "jhdksa",
		Button: "Trykke?",
		Secret: "not-super-secret",
	}

	body, err := util.ParseTemplate("email-templates/action.html", templateData)

	if err != nil {
		t.Error(err)
	} else {
		fmt.Println("body:", body)
	}

	println(body)
}
