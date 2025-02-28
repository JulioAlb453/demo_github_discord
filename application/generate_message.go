package application

import "fmt"

func generateDiscordMessage(base, head, html_url, user, repository_full_name string)string {
	return fmt.Sprintln("Nuevo pull reques a la rama %s en el repositorio %s Rama inicial: %s  Usuario: %s  ", base, head, html_url)
}