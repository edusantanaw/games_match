package controllers

import (
	"encoding/json"
	"fmt"

	"github.com/edusantanaw/games_match.git/db"
	"github.com/edusantanaw/games_match.git/entities"
	"github.com/edusantanaw/games_match.git/httpResponse"
	"github.com/google/uuid"
)

type CreateGameData struct {
	Name        string   `json:"name"`
	Categories  []string `json:"categories"`
	Description string   `json:"description"`
	Plataforms  []string `json:"plataforms"`
}

func RegisterGame(data *CreateGameData) httpResponse.HttpResponse {
	db := db.GetConnection()
	fmt.Println(data)
	if data.Name == "" {
		return httpResponse.BadRequest("O nome é obrigatorio!")
	}
	if len(data.Categories) == 0 {
		return httpResponse.BadRequest("Pelo menos uma categoria deve ser selecionada!")
	}
	if data.Description == "" {
		return httpResponse.BadRequest("A descrição é obrigatoria!")
	}
	if len(data.Plataforms) == 0 {
		return httpResponse.BadRequest("Pelo menos uma plataforma deve ser selecionada!")
	}
	parsedCategories, err := json.Marshal(data.Categories)
	if err != nil {
		return httpResponse.BadRequest("Erro ao converter array para json")
	}
	parsedPlataforms, err := json.Marshal(data.Plataforms)
	if err != nil {
		return httpResponse.BadRequest("Erro ao converter array para json")
	}
	game := &entities.Game{Name: data.Name, Id: uuid.New().String(), Categories: string(parsedCategories), Plataforms: string(parsedPlataforms), Description: data.Description}
	db.Create(game)
	return httpResponse.Created[entities.Game](*game)
}
