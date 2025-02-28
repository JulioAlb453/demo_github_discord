package infraestructure

import (
	"github/application"
	value_object "github/domain/value_objects"
	"log"
	"net/http"

	"encoding/json"
	"github.com/gin-gonic/gin"
)

func HandlePullRequestEvent(ctx *gin.Context) {

	eventType := ctx.GetHeader("X-GitHub-Event")
	deliveryD := ctx.GetHeader("X-GitHub-Delivery")

	log.Printf("Nuevo evento: %s con ID: %s", eventType, deliveryD)

	rawData, err := ctx.GetRawData()

	if err != nil {
		log.Printf("Error al leer el cuerpo de la solicitud: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "leer datos"})
	}

	var statusCode int

	switch eventType {
	case "ping":
		ctx.JSON(http.StatusOK, gin.H{"status": "success"})
	case "pull_request":
		statusCode = application.ProcessPullRequestEvent(rawData)
	}

	var payload value_object.PullRequestEvent

	if err := json.Unmarshal(rawData, &payload); err != nil {
        log.Printf("Error al deserializar el payload del pull request: %v", err)
        ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar el payload del pull request"})
    }


	if payload.Action == "closed" {
		log.Printf("Pull request cerrado")
		log.Println("Respositorio", payload.Repository.Name)
		log.Println("Usuario", payload.PullRequest.User.Login)
		log.Println("Desde", payload.PullRequest.Head.Ref)
		log.Println("Hacia", payload.PullRequest.Base.Ref)
	}

	switch statusCode {
	case 200:
		ctx.JSON(http.StatusOK, gin.H{"success": "Pull Request procesado con éxito"})
	case 403:
		log.Printf("Error al deserializar el payload del pull request: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar el payload del pull request"})
	default:
		ctx.JSON(http.StatusOK, gin.H{"success": "Normal"})
	}

}
