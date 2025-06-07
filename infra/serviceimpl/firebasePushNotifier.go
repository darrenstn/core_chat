package serviceimpl

import (
	"context"
	"errors"
	"fmt"

	chatservice "core_chat/application/chat/service"
	postservice "core_chat/application/post/service"
	pushnotificationdto "core_chat/application/pushnotification/dto"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/messaging"
	"google.golang.org/api/option"
)

var (
	_ chatservice.PushNotifier = (*FirebasePushNotifier)(nil)
	_ postservice.PushNotifier = (*FirebasePushNotifier)(nil)
)

type FirebasePushNotifier struct {
	client *messaging.Client
}

func NewFirebasePushNotifier(firebaseCredPath string) (*FirebasePushNotifier, error) {
	ctx := context.Background()

	app, err := firebase.NewApp(ctx, nil, option.WithCredentialsFile(firebaseCredPath))
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Firebase app: %w", err)
	}

	client, err := app.Messaging(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize Firebase messaging: %w", err)
	}

	return &FirebasePushNotifier{client: client}, nil
}

func (f *FirebasePushNotifier) Send(input pushnotificationdto.SendNotificationInput) error {
	if input.FirebaseToken == "" {
		return errors.New("no Firebase token provided")
	}

	msg := &messaging.Message{
		Token: input.FirebaseToken,
		Notification: &messaging.Notification{
			Title: input.Title,
			Body:  input.Body,
		},
		Data: map[string]string{
			"type":    input.Type,
			"sender":  input.Sender,
			"payload": input.Payload,
		},
	}

	_, err := f.client.Send(context.Background(), msg)
	return err
}
