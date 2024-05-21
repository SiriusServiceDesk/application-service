package web

import (
	"github.com/SiriusServiceDesk/application-service/internal/models"
	"github.com/SiriusServiceDesk/application-service/pkg/response"
)

type GetApplicationResponse struct {
	Id              string          `json:"id" example:"000000001"`
	Title           string          `json:"title" example:"сломался кампутир"`
	Status          models.Status   `json:"status" example:"Создана"`
	Performer       string          `json:"performer" example:"Методический отдел"`
	Priority        models.Priority `json:"priority" example:"низкий"`
	Comment         *string         `json:"comment" example:"любой коментарий ваще"`
	ApplicantId     string          `json:"applicant" example:"23ger34-khsdb23G-23afh75-sdHvd"`
	ExecutionPeriod string          `json:"execution_period" example:"7 рабочих дней"`
	FeedBack        string          `json:"feedback" example:"тут инфа видно только админу"`
	CreatedAt       string          `json:"create_date" example:"21.05.2024"`
}

type GetApplicationResponseDoc struct {
	response.RawResponse
	Payload struct {
		GetApplicationResponse
	} `json:"payload"`
}

type ArrayOfApplications []GetApplicationResponse
type GetApplicationsResponseDoc struct {
	response.RawResponse
	Payload struct {
		ArrayOfApplications `json:"applications"`
	} `json:"payload"`
}

type CreateApplicationRequest struct {
	Title     string  `json:"title" example:"Любой заголовок до 20 символов"`
	Comment   *string `json:"comment" example:"любой коммент или пустота"`
	Performer string  `json:"performer" example:"Методический отдел"`
}

type UpdateApplicationRequest struct {
	Status          models.Status   `json:"status" example:"В работе"`
	Priority        models.Priority `json:"priority" example:"Низкий"`
	ExecutionPeriod string          `json:"execution_period" example:"7 рабочих дней"`
	FeedBack        string          `json:"feedback" example:"причина отказа или комментарий админа"`
}

type AnalyticResponse struct {
	NewApplicationsToday     int `json:"pending"`
	ProcessedToday           int `json:"processed_today"`
	InProgress               int `json:"in_progress"`
	AllProcessedApplications int `json:"processed"`
}

type AnalyticResponseDoc struct {
	response.RawResponse
	Payload struct {
		AnalyticResponse
	} `json:"payload"`
}
