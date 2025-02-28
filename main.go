package main

import (
	"github/infraestructure"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	
	godotenv.Load()

	port := os.Getenv("PORT")

	if port == "" {
		port = "8080"
	}

	engine := gin.Default()

	infraestructure.Routes(engine)

	engine.Run(":" + port)
}