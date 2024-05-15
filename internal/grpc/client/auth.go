package client

import (
	"context"
	"github.com/SiriusServiceDesk/gateway-service/pkg/auth_v1"
	"github.com/urcop/go-fiber-template/internal/config"
	"github.com/urcop/go-fiber-template/pkg/logger"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func createConnectionToAuthService() (auth_v1.AuthV1Client, error) {
	cfg := config.GetConfig()
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, cfg.AuthService.Address, grpc.WithInsecure())
	if err != nil {
		logger.Error("Failed to connect gRPC auth service", zap.Error(err))
		return nil, err
	}

	return auth_v1.NewAuthV1Client(conn), nil
}

func GetUserIdFromToken(token string) (string, error) {
	conn, err := createConnectionToAuthService()
	if err != nil {
		return "", err
	}
	ctx := context.Background()

	response, err := conn.GetUserIdFromToken(ctx, &auth_v1.GetUserIdFromTokenRequest{Token: token})
	if err != nil {
		return "", err
	}

	return response.GetUserId(), nil
}
