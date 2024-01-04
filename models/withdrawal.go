package models

type Withdrawal struct {
	WithdrawalId string  `json:"withdrawal_id" bson:"_id"`
	Address      string  `json:"address" bson:"address"`
	UserId       string  `json:"user_id" bson:"user_id"`
	Amount       float64 `json:"amount" bson:"amount"`
	Date         int64   `json:"date" bson:"date"`
}