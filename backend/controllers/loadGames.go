package controllers

import (
	"github.com/edusantanaw/games_match.git/db"
	"github.com/edusantanaw/games_match.git/entities"
	"github.com/edusantanaw/games_match.git/httpResponse"
	"github.com/edusantanaw/games_match.git/structs"
)

func LoadGames(data structs.IPagination) httpResponse.HttpResponse {
	db := db.GetConnection()
	var games []entities.Game
	var total int64
	db.Model(entities.Game{}).Limit(data.Take).Offset(data.Offset).Find(&games)
	db.Model(entities.Game{}).Count(&total)
	println(len(games))
	response := structs.PaginationResponse[*[]entities.Game]{Data: &games, Total: total}
	return httpResponse.Ok(response)
}
