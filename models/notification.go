package models

import "time"

type Notification struct {
	Id              string    `bson:"_id,omitempty" json:"id"`
	Title           string    `bson:"title" json:"title"`
	Message         string    `bson:"message" json:"message"`
	UserId          string    `bson:"user_id" json:"user_id"`
	IsRead          bool      `bson:"is_read" json:"is_read"`
	NotificationType string   `bson:"notification_type" json:"notification_type"`
	CreatedAt       time.Time `bson:"created_at" json:"created_at"`
}
