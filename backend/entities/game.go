package entities

import (
	"encoding/json"
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	Id          string
	Name        string
	Categories  string
	Description string
	Plataforms  string
}

func CreateValidGame(name string, categories []string, description string, plataforms []string) (Game, error) {
	game := Game{}
	if name == "" {
		return game, errors.New("O nome é obrigatorio!")
	}
	if len(categories) == 0 {
		return game, errors.New("Pelo menos uma categoria deve ser selecionada!")
	}
	if description == "" {
		return game, errors.New("A descrição é obrigatoria!")
	}
	if len(plataforms) == 0 {
		return game, errors.New("Pelo menos uma plataforma deve ser selecionada!")
	}
	parsedCategories, err := json.Marshal(categories)
	if err != nil {
		return game, errors.New("Erro ao converter array para json")
	}
	parsedPlataforms, err := json.Marshal(plataforms)
	if err != nil {
		return game, errors.New("Erro ao converter array para json")
	}
	game.Id = uuid.New().String()
	game.Name = name
	game.Categories = string(parsedCategories)
	game.Plataforms = string(parsedPlataforms)
	game.Description = description
	return game, nil
}
