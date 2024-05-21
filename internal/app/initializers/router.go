package initializers

import (
	"github.com/SiriusServiceDesk/application-service/internal/app/dependencies"
	"github.com/SiriusServiceDesk/application-service/internal/web"
	"github.com/SiriusServiceDesk/application-service/internal/web/admin"
	"github.com/SiriusServiceDesk/application-service/internal/web/application"
	"github.com/SiriusServiceDesk/application-service/internal/web/status"
	"github.com/SiriusServiceDesk/application-service/internal/web/swagger"
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, container *dependencies.Container) {
	ctrl := buildRouters(container)

	for i := range ctrl {
		ctrl[i].DefineRouter(app)
	}
}

func buildRouters(container *dependencies.Container) []web.Controller {
	return []web.Controller{
		status.NewStatusController(),
		swagger.NewSwaggerController(),
		application.NewApplicationsController(container.ApplicationService),
		admin.NewController(container.ApplicationService),
	}
}
