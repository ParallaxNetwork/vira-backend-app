package models

type Content struct {
	ContentId   string `json:"content_id" bson:"_id"`
	ImageUrl    string `json:"image_url" bson:"image_url"`
	Title       string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	CreatedAt   string `json:"created_at" bson:"created_at"`
}