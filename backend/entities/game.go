package entities

import "gorm.io/gorm"

type Game struct {
	gorm.Model
	id          string
	name        string
	category    int
	description string
	plataforms  string
}
