package application

import (
	"github.com/SiriusServiceDesk/application-service/internal/grpc/client"
	"github.com/SiriusServiceDesk/application-service/internal/middleware"
	"github.com/SiriusServiceDesk/application-service/internal/models"
	"github.com/SiriusServiceDesk/application-service/internal/services"
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type Controller struct {
	applicationService services.ApplicationService
}

func NewApplicationsController(applicationService services.ApplicationService) *Controller {
	return &Controller{
		applicationService: applicationService,
	}
}

func (ctrl *Controller) DefineRouter(app *fiber.App) {
	applicationGroup := app.Group("/v1/applications")
	applicationGroup.Use(middleware.BenchmarkMiddleware())
	applicationGroup.Use(middleware.SetupCORS())

	applicationGroup.Get("/", ctrl.getApplications)
	applicationGroup.Get("/:id", ctrl.getApplication)
	applicationGroup.Post("/", ctrl.createApplication)
	applicationGroup.Put("/:id", ctrl.updateApplication)
}

func (ctrl *Controller) getApplications(ctx *fiber.Ctx) error {
	panic("not implemented")
}

func (ctrl *Controller) getApplication(ctx *fiber.Ctx) error {
	applicationId := ctx.Params("id")

	applicationIdInt, err := strconv.Atoi(applicationId)
	if err != nil {
		return Response().WithDetails(err).ServerInternalError(ctx, "failed to convert from parameter")
	}

	application, err := ctrl.applicationService.GetApplicationById(uint(applicationIdInt))
	if err != nil {
		return Response().WithDetails(err).ServerInternalError(ctx, "failed to get application from database")
	}

	return Response().StatusOK(ctx, mappingApplicationForUser(application))
}

func (ctrl *Controller) createApplication(ctx *fiber.Ctx) error {
	var request CreateApplicationRequest
	authHeader := ctx.GetReqHeaders()[fiber.HeaderAuthorization]

	if err := ctx.BodyParser(&request); err != nil {
		return Response().WithDetails(err).ServerInternalError(ctx, "failed to parse data")
	}

	userId, err := client.GetUserIdFromToken(authHeader)
	if err != nil {
		return Response().WithDetails(err).ServerInternalError(ctx, "failed to get userId")
	}

	application := &models.Application{
		Title:       request.Title,
		Status:      models.Pending,
		PerformerId: userId,
		Comment:     request.Comment,
	}

	if err := ctrl.applicationService.CreateApplication(application); err != nil {
		return Response().WithDetails(err).ServerInternalError(ctx, "failed to create application")
	}

	return Response().StatusOK(ctx, "application created successfully")
}

func (ctrl *Controller) updateApplication(ctx *fiber.Ctx) error {
	panic("not implemented")
}
