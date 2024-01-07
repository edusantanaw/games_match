package entities

import (
	"errors"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Id       string `gorm:"primaryKey"`
	Username string
	Email    string `gorm:"unique"`
	password string
	Phone    *string
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
	u.Id = uuid.New().String()
	u.Username = username
	u.Email = email
	u.password = password
	u.Phone = &phone
	u.Roles = roles
	return nil
}

func (u *Users) SetPassword(pass string) {
	u.password = pass
}
