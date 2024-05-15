package models

import "time"

type Application struct {
	Id              uint      `json:"id" gorm:"primaryKey"`
	Title           string    `json:"title"`
	Status          status    `json:"status"`
	Priority        string    `json:"priority"`
	PerformerId     string    `json:"performer"`
	Comment         *string   `json:"comment"`
	ApplicantId     string    `json:"applicant"`
	ExecutionPeriod string    `json:"execution_period"`
	CreatedAt       time.Time `json:"create_date"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type status string

const (
	InProgress status = "in_progress"
	Pending    status = "pending"
	Executed   status = "executed"
)
