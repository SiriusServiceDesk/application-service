package application

import (
	"github.com/SiriusServiceDesk/application-service/internal/models"
)

type GetApplicationUserResponse struct {
	Id              string        `json:"id"`
	Title           string        `json:"title"`
	Status          models.Status `json:"status"`
	Performer       string        `json:"performer"`
	Priority        string        `json:"priority"`
	Comment         *string       `json:"comment"`
	ApplicantId     string        `json:"applicant"`
	ExecutionPeriod string        `json:"execution_period"`
	CreatedAt       string        `json:"create_date"`
}

type CreateApplicationRequest struct {
	Title     string  `json:"title"`
	Comment   *string `json:"comment"`
	Performer string  `json:"performer"`
}
