package models

type FundingTransaction struct {
	Amount    float64 `json:"amount" bson:"amount"`
	AdminFee  float64 `json:"admin_fee" bson:"admin_fee"`
	Total     float64 `json:"total" bson:"total"`
	CreatedAt int64   `json:"created_at" bson:"created_at"`
}

type Funding struct {
	FundingId          string               `json:"funding_id" bson:"_id"`
	ProjectId          string               `json:"project_id" bson:"project_id"`
	TokenId            string               `json:"token_id" bson:"token_id"`
	UserId             string               `json:"user_id" bson:"user_id"`
	TransactionHistory []FundingTransaction `json:"transaction_history" bson:"transaction_history"`
	CreatedAt          int64                `json:"created_at" bson:"created_at"`
}
