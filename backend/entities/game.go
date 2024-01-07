package entities

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Game struct {
	gorm.Model
	Id          string `gorm:"primaryKey"`
	Name        string
	Categories  []string `gorm:"serializer:json"`
	Description string
	Plataforms  []string `gorm:"serializer:json"`
}

func CreateValidGame(name string, categories []string, description string, plataforms []string) (*Game, error) {
	game := &Game{}
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
	game.Id = uuid.New().String()
	game.Name = name
	game.Categories = categories
	game.Plataforms = plataforms
	game.Description = description
	return game, nil
}
