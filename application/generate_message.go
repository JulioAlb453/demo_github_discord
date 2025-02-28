package application

import "fmt"

func GenerateDiscordMessage(Base, Head, html_url, User, repository_full_name string)string {
	return fmt.Sprintf("Nuevo pull request a la rama %s en el repositorio %s Rama inicial: %s  Usuario: %s Repositorio ", Base, html_url, Head,User, )
}