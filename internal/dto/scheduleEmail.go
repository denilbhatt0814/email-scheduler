package dto

type CreateScheduledEmail struct {
	Recipient string `json:"recipient"`
	Subject   string `json:"subject"`
	Body      string `json:"body"`
	Schedule  string `json:"schedule_expression"`
}
