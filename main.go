// main.go
package main

import (
	"log"
	"os"
	"vira-backend-app/routers"

	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}

	port := os.Getenv("PORT")
	router := routers.SetupRouter()

	router.Run("0.0.0.0:" + port)
}
