package util

import (
	"bytes"
	"html/template"
)

type EmailData struct {
	Text1  string
	Text2  string
	Button string
  Secret string
}

func ParseTemplate(templateFileName string, emailData EmailData) (string, error) {
	t, err := template.ParseFiles(templateFileName)
	if err != nil {
		return "", err
	}

	buf := new(bytes.Buffer)
	err = t.Execute(buf, emailData)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
