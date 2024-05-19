package application

import (
	"github.com/SiriusServiceDesk/application-service/internal/grpc/client"
	"github.com/SiriusServiceDesk/application-service/internal/helpers"
	"github.com/SiriusServiceDesk/application-service/internal/middleware"
	"github.com/SiriusServiceDesk/application-service/internal/models"
	"github.com/SiriusServiceDesk/application-service/internal/services"
	"github.com/SiriusServiceDesk/application-service/pkg/logger"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
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
	var applications []*models.Application
	authHeaders := ctx.GetReqHeaders()[fiber.HeaderAuthorization]

	userId, err := client.GetUserIdFromToken(authHeaders)
	if err != nil {
		return Response().WithDetails(err).ServerInternalError(ctx, "cant get user id")
	}

	user, err := client.GetUserById(userId)
	if err != nil {
		return Response().WithDetails(err).ServerInternalError(ctx, "cant get user")
	}

	if user.GetRole() == "Админ" {
		applications, err = ctrl.applicationService.GetAllApplications()
	} else {
		applications, err = ctrl.applicationService.GetApplicationByUserId(userId)
	}

	if err != nil {
		return Response().WithDetails(err).ServerInternalError(ctx, "failed to get applications from database")
	}

	return Response().StatusOK(ctx, mappingApplicationsForUser(applications))
}

func (ctrl *Controller) getApplication(ctx *fiber.Ctx) error {
	applicationId := ctx.Params("id")
	authHeaders := ctx.GetReqHeaders()[fiber.HeaderAuthorization]

	applicationIdInt, err := helpers.FormatIdFromStringToUint(applicationId)
	if err != nil {
		return Response().WithDetails(err).ServerInternalError(ctx, "failed to convert from parameter")
	}

	application, err := ctrl.applicationService.GetApplicationById(applicationIdInt)
	if err != nil {
		return Response().WithDetails(err).ServerInternalError(ctx, "failed to get application from database")
	}

	userId, err := client.GetUserIdFromToken(authHeaders)
	if err != nil {
		logger.Debug("cant get user id from auth-service", zap.Error(err))
		return Response().WithDetails(err).ServerInternalError(ctx, "failed to get user from database")
	}

	user, err := client.GetUserById(userId)
	if err != nil {
		logger.Debug("cant get user from auth-service", zap.Error(err))
		return Response().WithDetails(err).ServerInternalError(ctx, "failed to get user")
	}

	if userId != application.ApplicantId && user.GetRole() != "Админ" && user.GetRole() != application.PerformerId {
		return Response().BadRequest(ctx, "user dont have permissions")
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
		ApplicantId: userId,
		Comment:     request.Comment,
		PerformerId: request.Performer,
	}

	if err := ctrl.applicationService.CreateApplication(application); err != nil {
		return Response().WithDetails(err).ServerInternalError(ctx, "failed to create application")
	}

	return Response().StatusOK(ctx, "application created successfully")
}

func (ctrl *Controller) updateApplication(ctx *fiber.Ctx) error {
	panic("not implemented")
}
