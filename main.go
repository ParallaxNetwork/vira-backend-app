// main.go
package main

import (
	"vira-backend-app/routers"
)

func main() {
	router := routers.SetupRouter()

	router.Run(":8080")
}
