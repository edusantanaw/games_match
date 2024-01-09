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
	Group       []Group
}

func (game *Game) ValidGame(name string, categories []string, description string, plataforms []string) error {
	if name == "" {
		return errors.New("O nome é obrigatorio!")
	}
	if len(categories) == 0 {
		return errors.New("Pelo menos uma categoria deve ser selecionada!")
	}
	if description == "" {
		return errors.New("A descrição é obrigatoria!")
	}
	if len(plataforms) == 0 {
		return errors.New("Pelo menos uma plataforma deve ser selecionada!")
	}
	game.Id = uuid.New().String()
	game.Name = name
	game.Categories = categories
	game.Plataforms = plataforms
	game.Description = description
	return nil
}
