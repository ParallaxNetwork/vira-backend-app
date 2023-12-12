package repositories

import (
	"fmt"
	"vira-backend-app/models"
	"vira-backend-app/utils"

	"go.mongodb.org/mongo-driver/bson"
)

type ContentRepository interface {
	FindAll() ([]models.Content, error)
	FindById(id string) (*models.Content, error)
	InsertOne(content *models.Content) (*models.Content, error)
	DeleteById(id string) error
	UpdateById(id string, updatedContent *models.Content) (*models.Content, error)
}

type contentRepository struct{}

func NewContentRepository() ContentRepository {
	return &contentRepository{}
}

func (r *contentRepository) FindAll() ([]models.Content, error) {
	db, ctx, err := utils.MongodbConnect()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	var contents []models.Content
	cursor, err := db.Collection("contents").Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	if err = cursor.All(ctx, &contents); err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return contents, nil
}

func (r *contentRepository) FindById(id string) (*models.Content, error) {
	db, ctx, err := utils.MongodbConnect()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	var content models.Content
	filter := bson.M{"_id": id}
	err = db.Collection("contents").FindOne(ctx, filter).Decode(&content)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return &content, nil
}

func (r *contentRepository) InsertOne(content *models.Content) (*models.Content, error) {
    db, ctx, err := utils.MongodbConnect()
    if err != nil {
        return nil, fmt.Errorf(err.Error())
    }

		content.ContentId = utils.GenerateRandomID()
		content.CreatedAt = utils.GetCurrentTime()
		fmt.Println(content)

    _, err = db.Collection("contents").InsertOne(ctx, content)
    if err != nil {
        return nil, fmt.Errorf(err.Error())
    }

    return content, nil
}

func (r *contentRepository) DeleteById(id string) error {
    db, ctx, err := utils.MongodbConnect()
    if err != nil {
        return fmt.Errorf(err.Error())
    }

    filter := bson.M{"_id": id}
    _, err = db.Collection("contents").DeleteOne(ctx, filter)
    if err != nil {
        return fmt.Errorf(err.Error())
    }

    return nil
}

func (r *contentRepository) UpdateById(id string, updatedContent *models.Content) (*models.Content, error) {
    db, ctx, err := utils.MongodbConnect()
    if err != nil {
        return nil, fmt.Errorf(err.Error())
    }

    filter := bson.M{"_id": id}
    update := bson.M{"$set": updatedContent}

    _, err = db.Collection("contents").UpdateOne(ctx, filter, update)
    if err != nil {
        return nil, fmt.Errorf(err.Error())
    }

    return updatedContent, nil
}
