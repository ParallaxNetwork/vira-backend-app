package dto

type FundingInsertOneDTO struct {
	ProjectId string  `json:"project_id"`
	UserId    string  `json:"user_id"`
	Amount    float64 `json:"amount"`
	AdminFee  float64 `json:"admin_fee"`
	Total     float64 `json:"total"`
}
