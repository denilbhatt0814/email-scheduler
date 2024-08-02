package service

import (
	"fmt"
	"log"

	"github.com/denilbhatt0814/email-scheduler/config"
	"github.com/denilbhatt0814/email-scheduler/internal/dto"
	"github.com/resend/resend-go/v2"
)

type MailService interface {
	SendEmail(input dto.Email) error
}

type mailService struct {
	client   *resend.Client
	fromMail string
}

func NewMailService(config config.AppConfig) MailService {
	client := resend.NewClient(config.ResendApiKey)
	return &mailService{
		client:   client,
		fromMail: config.FromMail,
	}
}

func (s mailService) SendEmail(input dto.Email) error {

	params := &resend.SendEmailRequest{
		From:    s.fromMail,
		To:      []string{input.Recipient},
		Html:    fmt.Sprintf("<p>%s</p>", input.Body),
		Subject: input.Subject,
	}

	sent, err := s.client.Emails.Send(params)
	if err != nil {
		return err
	}
	log.Println(sent.Id)
	return nil
}
