package application

import (
	"github.com/SiriusServiceDesk/application-service/internal/models"
)

func mappingApplicationForUser(application *models.Application) GetApplicationUserResponse {
	return GetApplicationUserResponse{
		Id:              application.Id,
		Title:           application.Title,
		Status:          application.Status,
		Priority:        application.Priority,
		Comment:         application.Comment,
		ApplicantId:     application.ApplicantId,
		ExecutionPeriod: application.ExecutionPeriod,
		CreatedAt:       application.CreatedAt,
	}
}
