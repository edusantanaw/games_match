package controllers

import (
	"github.com/edusantanaw/games_match.git/db"
	"github.com/edusantanaw/games_match.git/entities"
	"github.com/edusantanaw/games_match.git/httpResponse"
)

type CreateGameData struct {
	Name        string   `json:"name"`
	Categories  []string `json:"categories"`
	Description string   `json:"description"`
	Plataforms  []string `json:"plataforms"`
}

func RegisterGame(data *CreateGameData) httpResponse.HttpResponse {
	db := db.GetConnection()
	game, err := entities.CreateValidGame(data.Name, data.Categories, data.Description, data.Plataforms)
	if err != nil {
		return httpResponse.BadRequest[string](err.Error())
	}
	db.Create(&game)
	return httpResponse.Created[entities.Game](game)
}
