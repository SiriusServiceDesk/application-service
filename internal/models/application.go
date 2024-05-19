package models

import "time"

type Application struct {
	Id              uint      `json:"id" gorm:"primaryKey"`
	Title           string    `json:"title"`
	Status          Status    `json:"status"`
	Priority        string    `json:"priority"`
	PerformerId     string    `json:"performer"`
	Comment         *string   `json:"comment"`
	ApplicantId     string    `json:"applicant"`
	ExecutionPeriod string    `json:"execution_period"`
	CreatedAt       time.Time `json:"create_date"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type Status string

const (
	InProgress Status = "В работе"
	Pending    Status = "Заявка принята"
	Executed   Status = "Выполнена"
	Canceled   Status = "Отменена"
)
