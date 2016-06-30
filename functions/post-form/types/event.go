package types


// Event is the event message
type Event struct {
	Receiver string `json:"receiver"` // email or formId
	Data     string `json:"data"`
	Origin   string `json:"origin"`
}