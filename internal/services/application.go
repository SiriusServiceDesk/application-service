package services

import (
	"github.com/SiriusServiceDesk/application-service/internal/models"
	"github.com/SiriusServiceDesk/application-service/internal/repository"
	"time"
)

type ApplicationService interface {
	GetAllApplications() ([]*models.Application, error)
	GetApplicationByUserId(userId string) ([]*models.Application, error)
	GetApplicationById(id uint) (*models.Application, error)
	GetApplicationsByUser(userId string) ([]*models.Application, error)
	CreateApplication(application *models.Application) error
	UpdateApplication(application *models.Application, id uint) error
}

func (a ApplicationServiceImpl) GetApplicationsByUser(userId string) ([]*models.Application, error) {
	return a.repos.GetApplicationsByUser(userId)
}

func (a ApplicationServiceImpl) GetAllApplications() ([]*models.Application, error) {
	return a.repos.Get()
}

func (a ApplicationServiceImpl) GetApplicationByUserId(userId string) ([]*models.Application, error) {
	return a.repos.GetByUserId(userId)
}

func (a ApplicationServiceImpl) GetApplicationById(id uint) (*models.Application, error) {
	return a.repos.GetByApplicationId(id)
}

func (a ApplicationServiceImpl) CreateApplication(application *models.Application) error {
	return a.repos.Create(application)
}

func (a ApplicationServiceImpl) UpdateApplication(application *models.Application, id uint) error {
	existing, err := a.repos.GetByApplicationId(id)
	if err != nil {
		return err
	}

	existing = &models.Application{
		Id:              existing.Id,
		Title:           existing.Title,
		Status:          application.Status,
		Priority:        application.Priority,
		PerformerId:     existing.PerformerId,
		Comment:         existing.Comment,
		ApplicantId:     existing.ApplicantId,
		ExecutionPeriod: application.ExecutionPeriod,
		FeedBack:        application.FeedBack,
		CreatedAt:       existing.CreatedAt,
		UpdatedAt:       time.Now(),
	}

	if err := a.repos.Update(existing); err != nil {
		return err
	}

	return nil
}

type ApplicationServiceImpl struct {
	repos repository.ApplicationRepository
}

func NewApplicationService(repository repository.ApplicationRepository) ApplicationService {
	return ApplicationServiceImpl{repos: repository}
}
