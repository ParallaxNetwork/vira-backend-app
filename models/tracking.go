package models

import "vira-backend-app/enums"

type Tracking struct {
	TrackingId 						string				`json:"tracking_id" bson:"_id"`
	ProjectId 						string				`json:"project_id" bson:"project_id"`
	Batch 								int 					`json:"batch" bson:"batch"`
	SeedPlantingDate 			string				`json:"seed_planting_date" bson:"seed_planting_date"`
	OptimisticHarvestDate string				`json:"optimistic_harvest_date" bson:"optimistic_harvest_date"`
	HarvestDate 					string				`json:"harvest_date" bson:"harvest_date"`
	Status 								enums.Status 	`json:"status" bson:"status"`
}