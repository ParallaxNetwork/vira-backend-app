package models

import "vira-backend-app/enums"

type Tracking struct {
	TrackingId 						string				`json:"tracking_id" bson:"_id"`
	ProjectId 						string				`json:"project_id" bson:"project_id"`
	Batch 								int64					`json:"batch" bson:"batch"`
	SeedPlantingDate 			int64					`json:"seed_planting_date" bson:"seed_planting_date"`
	OptimisticHarvestDate int64					`json:"optimistic_harvest_date" bson:"optimistic_harvest_date"`
	HarvestDate 					int64					`json:"harvest_date" bson:"harvest_date"`
	Status 								enums.Status 	`json:"status" bson:"status"`
}