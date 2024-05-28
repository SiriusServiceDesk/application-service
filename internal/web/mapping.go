package web

import (
	"github.com/SiriusServiceDesk/application-service/internal/grpc/client"
	"github.com/SiriusServiceDesk/application-service/internal/helpers"
	"github.com/SiriusServiceDesk/application-service/internal/models"
	"sort"
)

func MappingApplicationForUser(application *models.Application) GetApplicationResponse {
	return GetApplicationResponse{
		Id:              helpers.FormatIdFromUintToString(application.Id),
		Title:           application.Title,
		Status:          application.Status,
		Comment:         application.Comment,
		ApplicantId:     application.ApplicantId,
		Performer:       cleanPerformer(application.PerformerId),
		ExecutionPeriod: application.ExecutionPeriod,
		CreatedAt:       helpers.FormatDate(application.CreatedAt),
	}
}

func MappingApplicationsForUser(applications []*models.Application) []GetApplicationResponse {
	var result []GetApplicationResponse
	for _, application := range applications {
		result = append(result, GetApplicationResponse{
			Id:              helpers.FormatIdFromUintToString(application.Id),
			Title:           application.Title,
			Status:          application.Status,
			Comment:         application.Comment,
			ApplicantId:     application.ApplicantId,
			Performer:       cleanPerformer(application.PerformerId),
			ExecutionPeriod: application.ExecutionPeriod,
			CreatedAt:       helpers.FormatDate(application.CreatedAt),
		})
	}

	if len(result) == 0 {
		result = []GetApplicationResponse{}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Id > result[j].Id
	})

	return result
}

func MappingApplicationForAdmin(application *models.Application) GetApplicationResponse {
	user, _ := client.GetUserById(application.ApplicantId)
	return GetApplicationResponse{
		Id:              helpers.FormatIdFromUintToString(application.Id),
		Title:           application.Title,
		Status:          application.Status,
		Performer:       cleanPerformer(application.PerformerId),
		Priority:        application.Priority,
		Comment:         application.Comment,
		ApplicantId:     user.GetName() + " " + user.GetSurname(),
		ExecutionPeriod: application.ExecutionPeriod,
		FeedBack:        application.FeedBack,
		CreatedAt:       helpers.FormatDate(application.CreatedAt),
	}
}

func MappingApplicationsForAdmin(applications []*models.Application) []GetApplicationResponse {
	var result []GetApplicationResponse
	for _, application := range applications {
		user, _ := client.GetUserById(application.ApplicantId)
		result = append(result, GetApplicationResponse{
			Id:              helpers.FormatIdFromUintToString(application.Id),
			Title:           application.Title,
			Status:          application.Status,
			Performer:       cleanPerformer(application.PerformerId),
			Priority:        application.Priority,
			Comment:         application.Comment,
			ApplicantId:     user.GetName() + " " + user.GetSurname(),
			ExecutionPeriod: application.ExecutionPeriod,
			FeedBack:        application.FeedBack,
			CreatedAt:       helpers.FormatDate(application.CreatedAt),
		})
	}

	if len(result) == 0 {
		result = []GetApplicationResponse{}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Id > result[j].Id
	})

	return result
}

func cleanPerformer(performer string) string {
	switch performer {
	case "":
		return "Не назначен"
	default:
		return performer
	}
}
