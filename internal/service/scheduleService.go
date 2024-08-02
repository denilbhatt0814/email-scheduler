package service

import (
	"github.com/denilbhatt0814/email-scheduler/config"
	"github.com/denilbhatt0814/email-scheduler/internal/repository"
)

type ScheduleService struct {
	Repo   repository.ScheduleRepository
	Config config.AppConfig
}
