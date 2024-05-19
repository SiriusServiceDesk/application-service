package application

import (
	"github.com/SiriusServiceDesk/application-service/internal/models"
)

func mappingApplicationForUser(application *models.Application) GetApplicationUserResponse {
	return GetApplicationUserResponse{
		Id:              application.Id,
		Title:           application.Title,
		Status:          application.Status,
		Comment:         application.Comment,
		ApplicantId:     application.ApplicantId,
		Performer:       application.PerformerId,
		ExecutionPeriod: application.ExecutionPeriod,
		CreatedAt:       application.CreatedAt,
	}
}

func mappingApplicationsForUser(applications []*models.Application) []GetApplicationUserResponse {
	var result []GetApplicationUserResponse
	for _, application := range applications {
		result = append(result, GetApplicationUserResponse{
			Id:              application.Id,
			Title:           application.Title,
			Status:          application.Status,
			Comment:         application.Comment,
			ApplicantId:     application.ApplicantId,
			Performer:       application.PerformerId,
			ExecutionPeriod: application.ExecutionPeriod,
			CreatedAt:       application.CreatedAt,
		})
	}
	return result
}
