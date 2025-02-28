package application

import (
	"encoding/json"
	value_object "github/domain/value_objects"

	"log"
)

func ProcessPullRequestEvent(rawData []byte) string {
	var eventPayload value_object.PullRequestEvent

	if err := json.Unmarshal(rawData, &eventPayload); err != nil {
		log.Println("Error al serializar payload")
		return "ERROR"
	}

	log.Printf("Evento pull request recibido con accion de %s", eventPayload.Action)

	base := eventPayload.PullRequest.Base.Ref
	titulo := eventPayload.PullRequest.Title
	repoFullName := eventPayload.Repository.FullName
	user := eventPayload.PullRequest.User.Login
	urlPullRequest := eventPayload.PullRequest.HTMLURL

	return GenerateDiscordMessage(base, titulo, repoFullName, user, urlPullRequest)
}
