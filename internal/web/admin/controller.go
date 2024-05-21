package admin

import (
	"github.com/SiriusServiceDesk/application-service/internal/middleware"
	"github.com/SiriusServiceDesk/application-service/internal/services"
	"github.com/SiriusServiceDesk/application-service/internal/web"
	"github.com/SiriusServiceDesk/application-service/pkg/response"
	"github.com/gofiber/fiber/v2"
)

type Controller struct {
	applicationService services.ApplicationService
}

func NewController(applicationService services.ApplicationService) *Controller {
	return &Controller{
		applicationService: applicationService,
	}
}

func (ctrl *Controller) DefineRouter(app *fiber.App) {
	adminGroup := app.Group("/v1/admin/applications")

	adminGroup.Use(middleware.SetupCORS())
	adminGroup.Use(middleware.NewAdminMiddleware())
	adminGroup.Use(middleware.BenchmarkMiddleware())

	adminGroup.Get("/", ctrl.getApplications)

}

// @Summary Get Applications
// @Description Get all applications (Admin only)
// @Tags admin
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Success 200 {object} web.GetApplicationsResponseDoc
// @Failure 400 {object} response.RawResponse
// @Failure 500 {object} response.RawResponse
// @Router /v1/admin/applications/ [get]
func (ctrl *Controller) getApplications(ctx *fiber.Ctx) error {
	applications, err := ctrl.applicationService.GetAllApplications()
	if err != nil {
		return response.Response().WithDetails(err).ServerInternalError(ctx, "failed to get applications")
	}
	return response.Response().StatusOK(ctx, web.MappingApplicationsForAdmin(applications))
}
