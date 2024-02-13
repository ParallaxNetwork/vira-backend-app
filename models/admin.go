package models

type Admin struct {
	AdminId       string `json:"admin_id" bson:"_id"`
	WalletAddress string `json:"wallet_address" bson:"wallet_address"`
	UserName      string `json:"username" bson:"username"`
	Email         string `json:"email" bson:"email"`
	Password      string `json:"password" bson:"password"`
	CreatedAt     int64  `json:"created_at" bson:"created_at"`
}
