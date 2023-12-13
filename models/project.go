package models

import "vira-backend-app/enums"

type Project struct {
	ProjectId     			string 				`json:"project_id" bson:"_id"`
	ProjectName   			string				`json:"project_name" bson:"project_name"`
	ProjectDesc   			string				`json:"project_desc" bson:"project_desc"`
	ImageUrl      			string				`json:"image_url" bson:"image_url"`
	Location      			string				`json:"location" bson:"location"`
	InitialValue  			float64 			`json:"initial_value" bson:"initial_value"`
	CollectedValue			float64 			`json:"collected_value" bson:"collected_value"` // Funding collected value
	FullReturnDuration	int64					`json:"full_return_duration" bson:"full_return_duration"`
	ContractDuration 		int64					`json:"contract_duration" bson:"contract_duration"`
	ReturnPerYear  			float64 			`json:"return_per_year" bson:"return_per_year"`
	Risk           			string 				`json:"risk" bson:"risk"`
	Status         			enums.Status 	`json:"status" bson:"status"`
	CreatedAt     			int64 				`json:"created_at" bson:"created_at"`
}