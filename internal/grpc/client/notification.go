package client

import (
	"context"
	"github.com/SiriusServiceDesk/application-service/internal/config"
	"github.com/SiriusServiceDesk/gateway-service/pkg/notification_v1"
	"google.golang.org/grpc"
)

type Message struct {
	Subject      string   `json:"subject"`
	To           []string `json:"to"`
	Data         string   `json:"data"`
	TemplateName string   `json:"template_name"`
	Type         string   `json:"type"`
}

type UpdateApplicationData struct {
	Id      string `json:"AppId"`
	Status  string `json:"Status"`
	Comment string `json:"Comment"`
}

func SendMessage(message Message) error {
	cfg := config.GetConfig()
	ctx := context.Background()

	conn, err := grpc.DialContext(ctx, cfg.NotificationService.Address, grpc.WithInsecure())
	if err != nil {
		return err
	}

	client := notification_v1.NewNotificationV1Client(conn)
	if _, err = client.CreateMessage(ctx, &notification_v1.CreateMessageRequest{
		Subject:      message.Subject,
		To:           message.To,
		Data:         message.Data,
		TemplateName: message.TemplateName,
		Type:         message.Type,
	}); err != nil {
		return err
	}

	return nil
}
