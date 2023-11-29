package models

type Funding struct {
	FundingId string `json:"funding_id" bson:"_id"`
	ProjectId string `json:"project_id" bson:"project_id"`
	UserId    string `json:"user_id" bson:"user_id"`
	Amount		float64 `json:"amount" bson:"amount"`
	AdminFee	float64 `json:"admin_fee" bson:"admin_fee"`
	Total			float64 `json:"total" bson:"total"`
	CreatedAt string `json:"created_at" bson:"created_at"`
}