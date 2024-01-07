package entities

import (
	"errors"

	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Id       string `gorm:"primaryKey"`
	username string
	email    string `gorm:"unique"`
	password string
	phone    *string
	Roles    []string `gorm:"serializer:json"`
}

func (u *Users) ValidUser(username string, email string, password string, phone string, roles []string) error {
	if username == "" {
		return errors.New("O username é obrigatorio!")
	}
	if email == "" {
		return errors.New("O email é obrigatorio!")
	}
	if password == "" {
		return errors.New("O senha é obrigatorio!")
	}
	if len(roles) == 0 {
		return errors.New("Os tipos de acesso são obrigatorios!")
	}
	u.username = username
	u.email = email
	u.password = password
	u.phone = &phone
	u.Roles = roles
	return nil
}
