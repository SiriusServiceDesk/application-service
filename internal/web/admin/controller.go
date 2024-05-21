package admin

import (
	"github.com/SiriusServiceDesk/application-service/internal/middleware"
	"github.com/SiriusServiceDesk/application-service/internal/models"
	"github.com/SiriusServiceDesk/application-service/internal/services"
	"github.com/SiriusServiceDesk/application-service/internal/web"
	"github.com/SiriusServiceDesk/application-service/pkg/response"
	"github.com/gofiber/fiber/v2"
	"time"
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
	adminGroup.Get("/analytic/", ctrl.analytic)

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

// Analytic provides statistics on applications
// @Summary Get application analytics
// @Description Retrieve statistics on applications, including the number of new applications today, all processed applications, applications processed today, and those in progress.
// @Tags analytics
// @Accept json
// @Produce json
// @Param Authorization header string true "Bearer <token>"
// @Success 200 {object} web.AnalyticResponse
// @Failure 500 {object} response.RawResponse "Failed to get applications"
// @Router /v1/admin/applications/analytic/ [get]
func (ctrl *Controller) analytic(ctx *fiber.Ctx) error {
	processedStatuses := []models.Status{models.Canceled, models.Executed}
	allProcessed, err := ctrl.applicationService.GetProcessedApplications(processedStatuses)
	if err != nil {
		return response.Response().WithDetails(err).ServerInternalError(ctx, "failed to get applications")
	}

	inProgressStatus := []models.Status{models.InProgress}
	inProgress, err := ctrl.applicationService.GetProcessedApplications(inProgressStatus)
	if err != nil {
		return response.Response().WithDetails(err).ServerInternalError(ctx, "failed to get applications")
	}

	today := time.Now().Format("2006-01-02")
	todayStatues := []models.Status{models.Canceled, models.InProgress, models.Pending, models.Executed}

	newApplicationsToday, err := ctrl.applicationService.GetProcessedApplicationsWithDate(todayStatues, today)
	if err != nil {
		return response.Response().WithDetails(err).ServerInternalError(ctx, "failed to get applications")
	}

	todayProcessed, err := ctrl.applicationService.GetProcessedApplicationsWithDate(processedStatuses, today)
	if err != nil {
		return response.Response().WithDetails(err).ServerInternalError(ctx, "failed to get applications")
	}

	analyticResponse := web.AnalyticResponse{
		NewApplicationsToday:     len(newApplicationsToday),
		AllProcessedApplications: len(allProcessed),
		ProcessedToday:           len(todayProcessed),
		InProgress:               len(inProgress),
	}

	return response.Response().StatusOK(ctx, analyticResponse)
}
