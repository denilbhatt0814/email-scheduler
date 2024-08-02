package domain

import "time"

type ScheduledEmail struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Recipient string    `json:"recipient"`
	Subject   string    `json:"subject"`
	Body      string    `json:"body"`
	JobID     int       `json:"job_id"`
	Schedule  string    `json:"schedule_expression"`
	CreatedAt time.Time `json:"created_at" gorm:"default:current_timestamp"`
	UpdatedAt time.Time `json:"updated_at" gorm:"default:current_timestamp"`
}
