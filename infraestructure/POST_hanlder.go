package infraestructure

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github/application"
	value_object "github/domain/value_objects"
)

func HandlePullRequestEvent(ctx *gin.Context) {

	eventType := ctx.GetHeader("X-GitHub-Event")
	deliveryD := ctx.GetHeader("X-GitHub-Delivery")

	log.Printf("Nuevo evento: %s con ID: %s", eventType, deliveryD)

	rawData, err := ctx.GetRawData()
	if err != nil {
		log.Printf("Error al leer el cuerpo de la solicitud: %v", err)
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error al leer los datos"})
		return
	}

	var statusCode int

	switch eventType {
	case "ping":
		ctx.JSON(http.StatusOK, gin.H{"status": "success"})
		return

	case "pull_request":
		message := application.ProcessPullRequestEvent(rawData)
		log.Printf("Mensaje de procesamiento: %s", message)

		if message == "ERROR" {
			statusCode = http.StatusInternalServerError
		} else {
			statusCode = post_discord(message)
		}
	}

	var payload value_object.PullRequestEvent
	if err := json.Unmarshal(rawData, &payload); err != nil {
		log.Printf("Error al deserializar el payload del pull request: %v", err)
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Error al procesar el payload del pull request"})
		return
	}

	if payload.Action == "closed" {
		log.Printf("Pull request cerrado")
		log.Println("Repositorio:", payload.Repository.FullName)
		log.Println("Usuario:", payload.PullRequest.User.Login)
		log.Println("Desde:", payload.PullRequest.Head.Ref)
		log.Println("Hacia:", payload.PullRequest.Base.Ref)
	}

	switch statusCode {
	case http.StatusOK:
		ctx.JSON(http.StatusOK, gin.H{"success": "Pull Request procesado con éxito"})
	case http.StatusForbidden:
		ctx.JSON(http.StatusForbidden, gin.H{"error": "No autorizado para procesar el pull request"})
	case http.StatusInternalServerError:
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Error interno al procesar el PR"})
	default:
		ctx.JSON(http.StatusOK, gin.H{"success": "Normal"})
	}
}
