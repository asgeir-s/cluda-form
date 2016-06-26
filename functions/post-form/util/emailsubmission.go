package util

import "bytes"

func CreateEmailBody(data map[string][]string) string {
	var bodyBuff bytes.Buffer

	for key, values := range data {
		bodyBuff.WriteString("<b>" + key + ":</b>")
		for _, value := range values {
			bodyBuff.WriteString("<p>" + value + "</p>")
		}
	}
	return bodyBuff.String()
}
