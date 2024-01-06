package controllers

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/edusantanaw/games_match.git/db"
	"github.com/edusantanaw/games_match.git/entities"
	"github.com/google/uuid"
)

type CreateGameData struct {
	Name        string   `json:"name"`
	Categories  []string `json:"categories"`
	Description string   `json:"description"`
	Plataforms  []string `json:"plataforms"`
}

func RegisterGame(data CreateGameData) (*entities.Game, error) {
	db := db.GetConnection()
	fmt.Println(data)
	if data.Name == "" {
		return &entities.Game{}, errors.New("O nome é obrigatorio!")
	}
	if len(data.Categories) == 0 {
		return &entities.Game{}, errors.New("Pelo menos uma categoria deve ser selecionada!")
	}
	if data.Description == "" {
		return &entities.Game{}, errors.New("A descrição é obrigatoria!")
	}
	if len(data.Plataforms) == 0 {
		return &entities.Game{}, errors.New("Pelo menos uma plataforma deve ser selecionada!")
	}
	parsedCategories, err := json.Marshal(data.Categories)
	if err != nil {
		return &entities.Game{}, errors.New("Erro ao converter array para json")
	}
	parsedPlataforms, err := json.Marshal(data.Plataforms)
	if err != nil {
		return &entities.Game{}, errors.New("Erro ao converter array para json")
	}
	game := &entities.Game{Name: data.Name, Id: uuid.New().String(), Categories: string(parsedCategories), Plataforms: string(parsedPlataforms), Description: data.Description}
	db.Create(game)
	return game, nil
}
