package repositories

import (
	"fmt"
	"vira-backend-app/utils"

	"firebase.google.com/go/messaging"
)

type NotificationRepository interface {
	NotificationPush() (error)
}

type notificationRepository struct {}

func NewNotificationRepository() NotificationRepository {
	return &notificationRepository{}
}

func (r *notificationRepository) NotificationPush() (error) {
	ctx, client := utils.SetupFirebase()
	topic := "vira-backend-app"

	message := &messaging.Message{
		Notification: &messaging.Notification{
			Title: "Test Notification",
			Body:  "This is a test message",
		},
		Topic: topic,
	}

	response, err := client.Send(ctx, message)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	fmt.Println("Successfully sent message:", response)

	return nil
}