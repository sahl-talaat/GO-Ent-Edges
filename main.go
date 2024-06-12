package main

import (
	"log"
	"os"
	"test/config"
	"test/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error loading .env file")
	}

	config.InitClient()

	app := gin.Default()

	routes.RegisterRoutes(app)

	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "9999"
	}

	err = app.Run(":" + port)
	if err != nil {
		log.Fatalf("error starting server: %v", err)
	}

}
