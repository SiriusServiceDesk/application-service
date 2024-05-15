package app

import (
	"github.com/SiriusServiceDesk/application-service/internal/app/dependencies"
	"github.com/SiriusServiceDesk/application-service/internal/app/initializers"
	"github.com/SiriusServiceDesk/application-service/internal/repository"
	"github.com/SiriusServiceDesk/application-service/internal/services"
	"github.com/gofiber/fiber/v2"
)

type Application struct{}

func InitApplication(app *fiber.App) {
	applicationRepos := repository.NewApplicationRepository()
	applicationService := services.NewApplicationService(applicationRepos)

	container := &dependencies.Container{
		ApplicationService: applicationService,
	}

	initializers.SetupRoutes(app, container)
}
