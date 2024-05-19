package application

import (
	"github.com/SiriusServiceDesk/application-service/internal/models"
	"time"
)

type GetApplicationUserResponse struct {
	Id              uint          `json:"id"`
	Title           string        `json:"title"`
	Status          models.Status `json:"status"`
	Performer       string        `json:"performer"`
	Priority        string        `json:"priority"`
	Comment         *string       `json:"comment"`
	ApplicantId     string        `json:"applicant"`
	ExecutionPeriod string        `json:"execution_period"`
	CreatedAt       time.Time     `json:"create_date"`
}

type CreateApplicationRequest struct {
	Title   string  `json:"title"`
	Comment *string `json:"comment"`
}
