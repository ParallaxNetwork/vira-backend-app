package models

type User struct {
	UserId         string `json:"user_id" bson:"_id"`
	WalletAddress  string `json:"wallet_address" bson:"wallet_address"`
	BalanceAmount  int    `json:"balance_amount" bson:"balance_amount"`
	CreatedAt      string `json:"created_at" bson:"created_at"`
}
