package repository

import (
	"errors"
	"log"

	"github.com/denilbhatt0814/email-scheduler/internal/domain"
	"gorm.io/gorm"
)

type ScheduleRepository interface {
	CreateEmailSchedule(email domain.ScheduledEmail) (domain.ScheduledEmail, error)
	FindScheduledEmails() ([]*domain.ScheduledEmail, error)
	FindScheduledEmailById(id int) (*domain.ScheduledEmail, error)
	DeleteScheduledEmail(id int) error
}

type scheduleRepository struct {
	db *gorm.DB
}

func NewScheduleRepository(db *gorm.DB) ScheduleRepository {
	return &scheduleRepository{
		db: db,
	}
}

func (r scheduleRepository) CreateEmailSchedule(email domain.ScheduledEmail) (domain.ScheduledEmail, error) {
	err := r.db.Create(&email).Error
	if err != nil {
		log.Println("create email schedule error: ", err)
		return domain.ScheduledEmail{}, errors.New("failed to create email schedule")
	}
	return email, err
}

func (r scheduleRepository) FindScheduledEmails() ([]*domain.ScheduledEmail, error) {
	var emails []*domain.ScheduledEmail

	err := r.db.Find(&emails).Error
	if err != nil {
		log.Println("error retrieving emails: ", err)
		return nil, err
	}

	return emails, nil
}

func (r scheduleRepository) FindScheduledEmailById(id int) (*domain.ScheduledEmail, error) {
	var email *domain.ScheduledEmail

	err := r.db.First(&email, id).Error
	if err != nil {
		log.Printf("error retrieving email with id[%v]: %v\n", id, err.Error())
		return nil, err
	}

	return email, nil
}

func (r scheduleRepository) DeleteScheduledEmail(id int) error {
	err := r.db.Delete(&domain.ScheduledEmail{}, id).Error

	if err != nil {
		log.Printf("error deleting email with id[%v]: %v\n", id, err.Error())
		return errors.New("error deleting scheduled email")
	}

	return nil
}
