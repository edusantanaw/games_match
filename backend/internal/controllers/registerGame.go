package controllers

import (
	"github.com/edusantanaw/games_match.git/internal/db"
	"github.com/edusantanaw/games_match.git/internal/entities"
	utils "github.com/edusantanaw/games_match.git/pkg/utils/httpResponse"
)

type CreateGameData struct {
	Name        string   `json:"name"`
	Categories  []string `json:"categories"`
	Description string   `json:"description"`
	Plataforms  []string `json:"plataforms"`
}

func RegisterGame(data *CreateGameData) utils.HttpResponse {
	db := db.GetConnection()
	game := &entities.Game{}
	err := game.ValidGame(data.Name, data.Categories, data.Description, data.Plataforms)
	if err != nil {
		return utils.BadRequest[string](err.Error())
	}
	db.Create(&game)
	return utils.Created[*entities.Game](game)
}
