package entities

import "gorm.io/gorm"

type Game struct {
	gorm.Model
	Id          string
	Name        string
	Categories  string
	Description string
	Plataforms  string
}
