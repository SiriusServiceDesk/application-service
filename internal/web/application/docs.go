package application

import (
	"github.com/SiriusServiceDesk/application-service/internal/models"
)

type GetApplicationUserResponse struct {
	Id              string        `json:"id" example:"000000001"`
	Title           string        `json:"title" example:"сломался кампутир"`
	Status          models.Status `json:"status" example:"Создана"`
	Performer       string        `json:"performer" example:"Методический отдел"`
	Priority        string        `json:"priority" example:"низкий"`
	Comment         *string       `json:"comment" example:"любой коментарий ваще"`
	ApplicantId     string        `json:"applicant" example:"23ger34-khsdb23G-23afh75-sdHvd"`
	ExecutionPeriod string        `json:"execution_period" example:"7 рабочих дней"`
	CreatedAt       string        `json:"create_date" example:"21.05.2024"`
}

type CreateApplicationRequest struct {
	Title     string  `json:"title" example:"сломался компутир"`
	Comment   *string `json:"comment" example:"любой коментарий ваще"`
	Performer string  `json:"performer" example:"Методический отдел"`
}

type GetApplicationUserResponseDoc struct {
	RawResponse
	Payload struct {
		GetApplicationUserResponse
	} `json:"payload"`
}
