package infraestructure

import (
	"log"

	"github.com/gin-gonic/gin"
)

func HandlePullRequestEvent(ctx *gin.Context) {

	eventType := ctx.GetHeader("X-GitHub-Event")
	deliveryD := ctx.GetHeader("X-GitHub-Delivery")
	
	log.Println("Nuevo evento: %s con ID: %s ", eventType, deliveryD)
	

}