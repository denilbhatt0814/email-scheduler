package service

import (
	"log"

	"github.com/robfig/cron/v3"
)

type CronService interface {
	ScheduleJob(spec string, job cron.Job) (int, error)
	Parse(spec string) error
}

type cronService struct {
	cron   *cron.Cron
	parser cron.Parser
}

func NewCronService() CronService {
	c := cron.New()
	c.Start()

	return &cronService{
		cron:   c,
		parser: cron.NewParser(cron.Minute | cron.Hour | cron.Dom | cron.Month | cron.Dow),
	}
}

func (s cronService) Parse(spec string) error {
	_, err := s.parser.Parse(spec)
	return err
}

func (s cronService) ScheduleJob(spec string, job cron.Job) (int, error) {

	schedule, err := s.parser.Parse("* * * * *")
	if err != nil {
		return -1, err
	}
	id := s.cron.Schedule(schedule, job)

	log.Println("Schedule id:", id)

	return int(id), nil
}
