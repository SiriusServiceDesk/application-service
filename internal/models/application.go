package models

type Application struct {
	Id              uint    `json:"id" gorm:"primaryKey"`
	Title           string  `json:"title"`
	Status          string  `json:"status"`
	Priority        string  `json:"priority"`
	PerformerId     string  `json:"performer"`
	Comment         *string `json:"comment"`
	ApplicantId     string  `json:"applicant"`
	ExecutionPeriod string  `json:"execution_period"`
	CreatedAt       string  `json:"create_date"`
	UpdatedAt       string  `json:"updated_at"`
}
