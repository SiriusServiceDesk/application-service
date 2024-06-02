package client

import (
	"context"
	"github.com/SiriusServiceDesk/application-service/internal/config"
	"github.com/SiriusServiceDesk/application-service/pkg/logger"
	"github.com/SiriusServiceDesk/gateway-service/pkg/auth_v1"
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

func GetUserIdFromToken(header []string) (string, error) {
	conn, err := createConnectionToAuthService()
	if err != nil {
		logger.Info("userId", zap.Error(err))
		return "", err
	}
	ctx := context.Background()

	response, err := conn.GetUserIdFromToken(ctx, &auth_v1.GetUserIdFromTokenRequest{Header: header})
	if err != nil {
		return "", err
	}

	logger.Info("userId", zap.String("userId", response.GetUserId()))

	return response.GetUserId(), nil
}

func GetUserById(userId string) (*auth_v1.GetUserByIdResponse, error) {
	conn, err := createConnectionToAuthService()
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	response, err := conn.GetUserById(ctx, &auth_v1.GetUserByIdRequest{UserId: userId})
	if err != nil {
		return nil, err
	}

	return response, nil
}
