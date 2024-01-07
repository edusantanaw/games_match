package entities

import "gorm.io/gorm"

type Users struct {
	gorm.Model
	Id       string `gorm:"primaryKey"`
	username string
	email    string `gorm:"unique"`
	password string
	phone    *string
	Roles    []string `gorm:"serializer:json"`
}
