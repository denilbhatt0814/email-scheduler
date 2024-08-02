package service

import (
	"errors"
	"fmt"
	"log"

	"github.com/denilbhatt0814/email-scheduler/config"
	"github.com/denilbhatt0814/email-scheduler/internal/domain"
	"github.com/denilbhatt0814/email-scheduler/internal/dto"
	"github.com/denilbhatt0814/email-scheduler/internal/repository"
)

type EmailSchedulerService struct {
	Repo   repository.ScheduleRepository
	Cron   CronService
	Config config.AppConfig
}

func (s *EmailSchedulerService) ScheduleEmail(input *dto.CreateScheduledEmail) error {

	sEmail := domain.ScheduledEmail{
		Recipient: input.Recipient,
		Subject:   input.Subject,
		Body:      input.Body,
		Schedule:  input.Schedule,
	}

	job := emailJob{
		sEmail,
	}
	jobID, err := s.Cron.ScheduleJob(sEmail.Schedule, job)
	if err != nil {
		return errors.New("error scheduling schedule email")
	}

	sEmail.JobID = jobID

	scheduledEmail, err := s.Repo.CreateEmailSchedule(sEmail)
	if err != nil {
		return err
	}
	log.Println("email scheduled successfully: ", scheduledEmail)

	return nil
}

func (s *EmailSchedulerService) GetScheduledEmails() ([]*domain.ScheduledEmail, error) {
	emails, err := s.Repo.FindScheduledEmails()
	if err != nil {
		return nil, err
	}

	return emails, nil
}

func (s *EmailSchedulerService) GetScheduledEmail(id int) (*domain.ScheduledEmail, error) {
	email, err := s.Repo.FindScheduledEmailById(id)
	if err != nil {
		return nil, errors.New("scheduled email does not exist")
	}

	return email, nil
}

func (s *EmailSchedulerService) DeleteScheduledEmail(id int) error {
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

type emailJob struct {
	domain.ScheduledEmail
}

func (j emailJob) Run() {
	fmt.Println(j)
}
