package infraestructure

import (
	"github/application"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandlePullRequestEvent(ctx *gin.Context) {
	eventType := ctx.GetHeader("X-GitHub-Event")
	deliveryD := ctx.GetHeader("X-GitHub-Delivery")

	log.Println("Nuevo evento: %s con ID: %s ", eventType, deliveryD)

	rawData, err := ctx.GetRawData()

	if err != nil {
		log.Println("Error al obtener los datos del cuerpo del evento: ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "error"})
		return
	}

	var statusCode int

	switch eventType {
	case "ping":
		ctx.JSON(http.StatusOK, gin.H{"status": "success"})

	case "pull-requesst":
		statusCode = application.ProcessPullRequestEvent(rawData)
	}

	switch statusCode {
	case 200:
		ctx.JSON(http.StatusOK, gin.H{"status": "Pull request procesdo con exito"})

	case 403:
		log.Println("Error al deserializar el payload del pull request: %v ", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"status": "Error aal procesar el payload del pull request"})

	default:
		ctx.JSON(http.StatusOK, gin.H{"status": "Normal"})
	}

}
