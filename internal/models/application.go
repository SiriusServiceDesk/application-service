package models

import "time"

type Application struct {
	Id              uint      `json:"id" gorm:"primaryKey"`
	Title           string    `json:"title"`
	Status          Status    `json:"status"`
	Priority        Priority  `json:"priority"`
	PerformerId     string    `json:"performer"`
	Comment         *string   `json:"comment"`
	ApplicantId     string    `json:"applicant"`
	ExecutionPeriod string    `json:"execution_period"`
	FeedBack        string    `json:"feedback"`
	CreatedAt       time.Time `json:"create_date"`
	UpdatedAt       time.Time `json:"updated_at"`
}

type Status string

const (
	InProgress Status = "В работе"
	Pending    Status = "Создана"
	Executed   Status = "Выполнена"
	Canceled   Status = "Отклонена"
)

type Priority string

const (
	Low    Priority = "Низкий"
	Medium Priority = "Средний"
	High   Priority = "Высокий"
	NotSet Priority = "Не назначен"
)
