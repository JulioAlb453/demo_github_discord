package main

import (
	"github/infraestructure"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	
	infraestructure.Routes(engine)

	engine.Run(":4000") 
}