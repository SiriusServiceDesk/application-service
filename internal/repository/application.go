package repository

import (
	"fmt"
	"github.com/SiriusServiceDesk/application-service/internal/config"
	"github.com/SiriusServiceDesk/application-service/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ApplicationRepository interface {
	Get() ([]*models.Application, error)
	GetByApplicationId(id uint) (*models.Application, error)
	GetByUserId(userId string) ([]*models.Application, error)
	Create(application *models.Application) error
	Update(application *models.Application) error
	GetApplicationsByUser(userId string) ([]*models.Application, error)
	GetProcessedApplications(statuses []models.Status) ([]*models.Application, error)
	GetProcessedApplicationsWithDate(statuses []models.Status, date string) ([]*models.Application, error)
}

func (a ApplicationRepositoryImpl) GetProcessedApplications(statuses []models.Status) ([]*models.Application, error) {
	var applications []*models.Application
	if err := a.db.Where("status in ?", statuses).Find(&applications).Error; err != nil {
		return nil, err
	}
	return applications, nil
}

func (a ApplicationRepositoryImpl) GetProcessedApplicationsWithDate(statuses []models.Status, date string) ([]*models.Application, error) {
	var applications []*models.Application
	if err := a.db.Where("DATE(created_at) = ? and status in ?", date, statuses).Find(&applications).Error; err != nil {
		return nil, err
	}
	return applications, nil
}

func (a ApplicationRepositoryImpl) GetApplicationsByUser(userId string) ([]*models.Application, error) {
	var applications []*models.Application
	if err := a.db.Where(models.Application{ApplicantId: userId}).Find(&applications).Error; err != nil {
		return nil, err
	}

	return applications, nil
}

func (a ApplicationRepositoryImpl) Get() ([]*models.Application, error) {
	var applications []*models.Application
	if err := a.db.Find(&applications).Error; err != nil {
		return nil, err
	}
	return applications, nil
}

func (a ApplicationRepositoryImpl) GetByApplicationId(id uint) (*models.Application, error) {
	var application *models.Application
	if err := a.db.Where(models.Application{Id: id}).First(&application).Error; err != nil {
		return nil, err
	}
	return application, nil
}

func (a ApplicationRepositoryImpl) GetByUserId(userId string) ([]*models.Application, error) {
	var application []*models.Application
	if err := a.db.Where(models.Application{ApplicantId: userId}).Find(&application).Error; err != nil {
		return nil, err
	}
	return application, nil
}

func (a ApplicationRepositoryImpl) Create(application *models.Application) error {
	if err := a.db.Create(&application).Error; err != nil {
		return err
	}
	return nil
}

func (a ApplicationRepositoryImpl) Update(application *models.Application) error {
	if err := a.db.Save(&application).Error; err != nil {
		return err
	}
	return nil
}

type ApplicationRepositoryImpl struct {
	db *gorm.DB
}

func NewApplicationRepository() ApplicationRepository {
	cfg := config.GetConfig()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
		cfg.Db.Host, cfg.Db.User, cfg.Db.Password, cfg.Db.Name, cfg.Db.Port)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	pgSvc := &ApplicationRepositoryImpl{db: db}
	err = db.AutoMigrate(&models.Application{})
	if err != nil {
		panic(err)
	}
	return pgSvc
}
