package middleware

import (
	"github.com/SiriusServiceDesk/application-service/internal/grpc/client"
	"github.com/SiriusServiceDesk/application-service/pkg/response"
	"github.com/gofiber/fiber/v2"
)

func NewAdminMiddleware() fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authHeaders := ctx.GetReqHeaders()[fiber.HeaderAuthorization]

		userId, err := client.GetUserIdFromToken(authHeaders)
		if err != nil {
			return response.Response().WithDetails(err).ServerInternalError(ctx, "can't get user id")
		}

		user, err := client.GetUserById(userId)
		if err != nil {
			return response.Response().WithDetails(err).ServerInternalError(ctx, "can't get user")
		}

		if user.GetRole() != "Админ" {
			return response.Response().BadRequest(ctx, "permissions denied")
		}

		ctx.Locals("user", user)
		return ctx.Next()
	}
}
