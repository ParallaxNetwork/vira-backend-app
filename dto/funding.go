package dto

type FundingInsertOneDTO struct {
	ProjectId string `json:"project_id"`
	UserId    string `json:"user_id"`
	Amount    int    `json:"amount"`
	AdminFee  int    `json:"admin_fee"`
	Total     int    `json:"total"`
}
