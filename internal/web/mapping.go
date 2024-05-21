package web

import (
	"github.com/SiriusServiceDesk/application-service/internal/helpers"
	"github.com/SiriusServiceDesk/application-service/internal/models"
	"sort"
)

func MappingApplicationForUser(application *models.Application) GetApplicationUserResponse {
	return GetApplicationUserResponse{
		Id:              helpers.FormatIdFromUintToString(application.Id),
		Title:           application.Title,
		Status:          application.Status,
		Comment:         application.Comment,
		ApplicantId:     application.ApplicantId,
		Performer:       application.PerformerId,
		ExecutionPeriod: application.ExecutionPeriod,
		CreatedAt:       helpers.FormatDate(application.CreatedAt),
	}
}

func MappingApplicationsForUser(applications []*models.Application) []GetApplicationUserResponse {
	var result []GetApplicationUserResponse
	for _, application := range applications {
		result = append(result, GetApplicationUserResponse{
			Id:              helpers.FormatIdFromUintToString(application.Id),
			Title:           application.Title,
			Status:          application.Status,
			Comment:         application.Comment,
			ApplicantId:     application.ApplicantId,
			Performer:       application.PerformerId,
			ExecutionPeriod: application.ExecutionPeriod,
			CreatedAt:       helpers.FormatDate(application.CreatedAt),
		})
	}

	if len(result) == 0 {
		result = []GetApplicationUserResponse{}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Id > result[j].Id
	})

	return result
}
