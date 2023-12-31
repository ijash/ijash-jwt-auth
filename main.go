package main

import (
	"fmt"
	"ijash-jwt-auth/src/configs"
	"ijash-jwt-auth/src/routes"
	"os"

	"github.com/labstack/echo/v4"
)

func main() {
	// LoadEnv()
	configs.InitDatabase()
	e := echo.New()
	routes.InitRoute(e)
	// e.Start(":8000")
	port := fmt.Sprintf(":%s", os.Getenv("PORT"))
	e.Start(port)

}

// func LoadEnv() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	fmt.Printf("Environment: %s", os.Getenv("ENV"))
// }
