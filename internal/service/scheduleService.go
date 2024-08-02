package service

import (
	"errors"
	"log"

	"github.com/denilbhatt0814/email-scheduler/config"
	"github.com/denilbhatt0814/email-scheduler/internal/domain"
	"github.com/denilbhatt0814/email-scheduler/internal/dto"
	"github.com/denilbhatt0814/email-scheduler/internal/repository"
)

type ScheduleService struct {
	Repo   repository.ScheduleRepository
	Config config.AppConfig
}

func (s *ScheduleService) ScheduleEmail(input *dto.CreateScheduledEmail) error {

	sEmail := domain.ScheduledEmail{
		Recipient: input.Recipient,
		Subject:   input.Subject,
		Body:      input.Body,
		Schedule:  input.Schedule,
	}

	scheduledEmail, err := s.Repo.CreateEmailSchedule(sEmail)
	if err != nil {
		return err
	}
	log.Println("email scheduled successfully: ", scheduledEmail)

	return nil
}

func (s *ScheduleService) GetScheduledEmails() ([]*domain.ScheduledEmail, error) {
	emails, err := s.Repo.FindScheduledEmails()
	if err != nil {
		return nil, err
	}

	return emails, nil
}

func (s *ScheduleService) GetScheduledEmail(id int) (*domain.ScheduledEmail, error) {
	email, err := s.Repo.FindScheduledEmailById(id)
	if err != nil {
		return nil, errors.New("scheduled email does not exist")
	}

	return email, nil
}

func (s *ScheduleService) DeleteScheduledEmail(id int) error {
	_, err := s.Repo.FindScheduledEmailById(id)
	if err != nil {
		return errors.New("schduled email does not exist")
	}

	err = s.Repo.DeleteScheduledEmail(id)
	if err != nil {
		return err
	}

	return nil
}
