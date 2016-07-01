package types

// Config is the functions config
type Config struct {
	AwsRegion       string `env:"AWS_REGION" envDefault:"us-west-2"`
	FormTable       string `env:"FORM_TABLE" envDefault:"test-form-table"`
	EmailFromAddres string `env:"EMAIL_FROM_ADDRES" envDefault:"test-in-1@coinsignals.com"`
	SubmissionTable string `env:"SUMBMISSION_TABLE" envDefault:"test-submission-table"`
	BaseURL         string `env:"BASE_URL" envDefault:"https://fqcx5ghvc3.execute-api.us-west-2.amazonaws.com/prod"`
}
