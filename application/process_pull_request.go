package application

import (
	"encoding/json"
	value_objects "github/domain/value_objects"
	// "html"
	"log"
)

func ProcessPullRequestEvent(rawData []byte) int {
	var eventPayload value_objects.PullRequestEvent

	if err := json.Unmarshal(rawData, &eventPayload); err != nil {
		return 403
	}

	log.Printf("Evento pull request recibido con accion de %s", eventPayload.Action)

	// base := eventPayload.PullRequest.Base.Ref
	// head := eventPayload.PullRequest.
	// html_url
	// user
	// reposotiry_full_name
	return 200
}