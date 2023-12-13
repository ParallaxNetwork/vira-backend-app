package models

type User struct {
	UserId        string `json:"user_id" bson:"_id"`
	WalletAddress string `json:"wallet_address" bson:"wallet_address"`
	BalanceAmount int64  `json:"balance_amount" bson:"balance_amount"`
	CreatedAt     int64  `json:"created_at" bson:"created_at"`
}
