package serviceimpl

import (
	"context"
	appdto "core_chat/application/notification/dto"
	"core_chat/application/notification/service"
	"core_chat/infra/serviceimpl/mapper"
	"encoding/json"
	"errors"
	"fmt"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

type NotifierServiceImpl struct {
	WebSocket service.WebSocketManager
	Firebase  *messaging.Client
}

func NewNotifierServiceImpl(ws service.WebSocketManager, firebaseCredPath string) (service.NotifierService, error) {
	ctx := context.Background()

	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile(firebaseCredPath))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Firebase app: %w", err)
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Firebase messaging: %w", err)
	}

	return &NotifierServiceImpl{
		WebSocket: ws,
		Firebase:  client,
	}, nil
}

func (n *NotifierServiceImpl) SetWebSocketManager(ws service.WebSocketManager) {
	n.WebSocket = ws
}

func (n *NotifierServiceImpl) IsOnline(identifier string) bool {
	if n.WebSocket == nil {
		return false
	}
	return n.WebSocket.IsOnline(identifier)
}

func (n *NotifierServiceImpl) Send(input appdto.SendNotificationInput) error {
	// Convert application DTO to infra DTO (with json tags)
	infraInput := mapper.ToInfraSendNotificationInput(input)

	data, err := json.Marshal(infraInput)
	if err != nil {
		return fmt.Errorf("failed to marshal notification data: %w", err)
	}

	// WebSocket: if user is online
	if n.WebSocket != nil && n.WebSocket.IsOnline(infraInput.Receiver) {
		return n.WebSocket.Send(infraInput.Receiver, data)
	}

	// Firebase: fallback
	if infraInput.FirebaseToken == "" {
		return errors.New("receiver is offline and no Firebase token provided")
	}

	msg := &messaging.Message{
		Token: infraInput.FirebaseToken,
		Notification: &messaging.Notification{
			Title: infraInput.Title,
			Body:  infraInput.Body,
		},
		Data: map[string]string{
			"type":    infraInput.Type,
			"sender":  infraInput.Sender,
			"payload": infraInput.Payload,
		},
	}

	_, err = n.Firebase.Send(context.Background(), msg)
	return err
}
