package entities

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Id       string
	username string
	email    string
	password string
	phone    string
}
