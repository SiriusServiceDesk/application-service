package application

import (
	"github.com/SiriusServiceDesk/application-service/internal/grpc/client"
	"github.com/SiriusServiceDesk/application-service/internal/helpers"
	"github.com/SiriusServiceDesk/application-service/internal/middleware"
	"github.com/SiriusServiceDesk/application-service/internal/models"
	"github.com/SiriusServiceDesk/application-service/internal/services"
	"github.com/SiriusServiceDesk/application-service/internal/web"
	"github.com/SiriusServiceDesk/application-service/pkg/logger"
	"github.com/SiriusServiceDesk/application-service/pkg/response"
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

// getApplications получает список заявок
// @Summary Получить заявки
// @Description Получает список заявок для администратора или пользователя
// @Param Authorization header string true "Bearer <token>"
// @Tags applications
// @Produce json
// @Success 200 {array} web.GetApplicationsResponseDoc
// @Failure 500 {object} response.RawResponse
// @Router /v1/applications [get]
// @Security ApiKeyAuth
func (ctrl *Controller) getApplications(ctx *fiber.Ctx) error {
	var applications []*models.Application
	authHeaders := ctx.GetReqHeaders()[fiber.HeaderAuthorization]

	userId, err := client.GetUserIdFromToken(authHeaders)
	if err != nil {
		return response.Response().WithDetails(err).ServerInternalError(ctx, "cant get user id")
	}

	applications, err = ctrl.applicationService.GetApplicationByUserId(userId)

	if err != nil {
		return response.Response().WithDetails(err).ServerInternalError(ctx, "failed to get applications from database")
	}

	return response.Response().StatusOK(ctx, web.MappingApplicationsForUser(applications))
}

// getApplication получает заявку по ID
// @Summary Получить заявку
// @Description Получает заявку по ID для администратора или пользователя
// @Tags applications
// @Param Authorization header string true "Bearer <token>"
// @Produce json
// @Param id path int true "ID заявки"
// @Success 200 {object} web.GetApplicationResponseDoc
// @Failure 400 {object} response.RawResponse
// @Failure 500 {object} response.RawResponse
// @Router /v1/applications/{id} [get]
func (ctrl *Controller) getApplication(ctx *fiber.Ctx) error {
	applicationId := ctx.Params("id")
	authHeaders := ctx.GetReqHeaders()[fiber.HeaderAuthorization]

	applicationIdInt, err := helpers.FormatIdFromStringToUint(applicationId)
	if err != nil {
		return response.Response().WithDetails(err).ServerInternalError(ctx, "failed to convert from parameter")
	}

	application, err := ctrl.applicationService.GetApplicationById(applicationIdInt)
	if err != nil {
		return response.Response().WithDetails(err).ServerInternalError(ctx, "failed to get application from database")
	}

	userId, err := client.GetUserIdFromToken(authHeaders)
	if err != nil {
		logger.Debug("cant get user id from auth-service", zap.Error(err))
		return response.Response().WithDetails(err).ServerInternalError(ctx, "failed to get user from database")
	}

	user, err := client.GetUserById(userId)
	if err != nil {
		logger.Debug("cant get user from auth-service", zap.Error(err))
		return response.Response().WithDetails(err).ServerInternalError(ctx, "failed to get user")
	}

	if userId != application.ApplicantId && user.GetRole() != "Админ" && user.GetRole() != application.PerformerId {
		return response.Response().BadRequest(ctx, "user dont have permissions")
	}

	return response.Response().StatusOK(ctx, web.MappingApplicationForUser(application))
}

// createApplication создает новую заявку
// @Summary Создать заявку
// @Description Создает новую заявку
// @Tags applications
// @Param Authorization header string true "Bearer <token>"
// @Accept json
// @Produce json
// @Param application body web.CreateApplicationRequest true "Создание заявки"
// @Success 200 {object} response.RawResponse
// @Failure 400 {object} response.RawResponse
// @Failure 500 {object} response.RawResponse
// @Router /v1/applications [post]
// @Security ApiKeyAuth
func (ctrl *Controller) createApplication(ctx *fiber.Ctx) error {
	var request web.CreateApplicationRequest
	authHeader := ctx.GetReqHeaders()[fiber.HeaderAuthorization]

	if err := ctx.BodyParser(&request); err != nil {
		return response.Response().WithDetails(err).ServerInternalError(ctx, "failed to parse data")
	}

	userId, err := client.GetUserIdFromToken(authHeader)
	if err != nil {
		return response.Response().WithDetails(err).ServerInternalError(ctx, "failed to get userId")
	}

	application := &models.Application{
		Title:       request.Title,
		Status:      models.Pending,
		ApplicantId: userId,
		Comment:     request.Comment,
		PerformerId: request.Performer,
		Priority:    models.NotSet,
	}

	if err := ctrl.applicationService.CreateApplication(application); err != nil {
		return response.Response().WithDetails(err).ServerInternalError(ctx, "failed to create application")
	}

	return response.Response().StatusOK(ctx, "application created successfully")
}

// updateApplication updates an application
// @Summary Update Application
// @Description Update an existing application by ID
// @Tags applications
// @Accept json
// @Produce json
// @Param id path string true "Application ID"
// @Param Authorization header string true "Authorization token"
// @Param request body web.UpdateApplicationRequest true "Update Application Request"
// @Success 200 {object} response.RawResponse "application updated"
// @Failure 400 {object} response.RawResponse "Bad Request"
// @Failure 500 {object} response.RawResponse "Internal Server Error"
// @Router /v1/admin/applications/{id} [put]
func (ctrl *Controller) updateApplication(ctx *fiber.Ctx) error {
	authHeaders := ctx.GetReqHeaders()[fiber.HeaderAuthorization]
	id := ctx.Params("id")
	var request web.UpdateApplicationRequest

	if err := ctx.BodyParser(&request); err != nil {
		return response.Response().WithDetails(err).ServerInternalError(ctx, "failed to parse data")
	}

	userId, err := client.GetUserIdFromToken(authHeaders)
	if err != nil {
		return response.Response().WithDetails(err).ServerInternalError(ctx, "failed to get user id")
	}

	user, err := client.GetUserById(userId)
	if err != nil {
		return response.Response().WithDetails(err).ServerInternalError(ctx, "failed to get user")
	}

	if user.GetRole() == "Пользователь" {
		return response.Response().BadRequest(ctx, "permission denied")
	}
	uintId, err := helpers.FormatIdFromStringToUint(id)
	if err != nil {
		return response.Response().WithDetails(err).ServerInternalError(ctx, "failed to format id")
	}

	newApplication := &models.Application{
		Status:          request.Status,
		Priority:        request.Priority,
		ExecutionPeriod: request.ExecutionPeriod,
		FeedBack:        request.FeedBack,
	}

	if err := ctrl.applicationService.UpdateApplication(newApplication, uintId); err != nil {
		return response.Response().WithDetails(err).ServerInternalError(ctx, "failed to update application")
	}

	return response.Response().StatusOK(ctx, "application updated")
}
