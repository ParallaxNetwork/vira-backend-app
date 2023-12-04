package repositories

import (
	"fmt"

	"vira-backend-app/enums"
	"vira-backend-app/models"
	"vira-backend-app/utils"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProjectRepository interface {
	FindById(projectId string) (*models.Project, error)
	FindAll() ([]models.Project, error)
	InsertOne(project models.Project) (*models.Project, error)
	DeleteById(projectId string) (error)
	UpdateById(projectId string, updatedProject *models.Project) (*models.Project, error)
}

type projectRepository struct {}

func NewProjectRepository() ProjectRepository {
	return &projectRepository{}
}

func (r *projectRepository) FindById(projectId string) (*models.Project, error) {
	db, ctx, err := utils.MongodbConnect()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	var project models.Project

	filter := bson.M{"_id": projectId}

	err = db.Collection("projects").FindOne(ctx, filter).Decode(&project)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, fmt.Errorf("No Project Found: %v", err)
		}

		return nil, fmt.Errorf("Database Error: %v", err)
	}

	return &project, nil
}

func (r *projectRepository) FindAll() ([]models.Project, error) {
	db, ctx, err := utils.MongodbConnect()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	var projects []models.Project

	cursor, err := db.Collection("projects").Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	err = cursor.All(ctx, &projects)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	for i := range projects {
    cursor, err := db.Collection("fundings").Find(ctx, bson.M{"project_id": projects[i].ProjectId})
    if err != nil {
        return nil, fmt.Errorf(err.Error())
    }

    var fundings []models.Funding
    var fundingCountTotal float64 = 0
    err = cursor.All(ctx, &fundings)
    if err != nil {
        return nil, fmt.Errorf(err.Error())
    }

    for _, funding := range fundings {
        fundingCountTotal += funding.Amount
    }

    projects[i].CollectedValue = fundingCountTotal
	}

	return projects, nil
}

func (r *projectRepository) InsertOne(project models.Project) (*models.Project, error) {
	db, ctx, err := utils.MongodbConnect()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	project.ProjectId = utils.GenerateRandomID()
	project.CreatedAt = utils.GetCurrentTime()
	project.Status 		= enums.Crowdfunding

	_, err = db.Collection("projects").InsertOne(ctx, project)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return &project, nil
}

func (r *projectRepository) DeleteById(projectId string) (error) {
	db, ctx, err := utils.MongodbConnect()
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	filter := bson.M{"_id": projectId}
	_, err = db.Collection("projects").DeleteOne(ctx, filter)
	if err != nil {
		return fmt.Errorf(err.Error())
	}

	return nil
}

func (r *projectRepository) UpdateById(projectId string, updatedProject *models.Project) (*models.Project, error) {
	db, ctx, err := utils.MongodbConnect()
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	filter := bson.M{"_id": projectId}
	update := bson.M{"$set": updatedProject}

	_, err = db.Collection("projects").UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return updatedProject, nil
}