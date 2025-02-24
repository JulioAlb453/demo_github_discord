package application

import (
	"encoding/json"
	value_objects "github/domain/value_objects"

	"log"
)

func ProcessPullRequestEvent(rawData []byte) int {
	var eventPayload value_objects.PullRequestEvent

	if err := json.Unmarshal(rawData, &eventPayload); err != nil {
		return 403
	}

	log.Printf("Evento pull request recibido con accion de %s", eventPayload.Action)

	return 200
}