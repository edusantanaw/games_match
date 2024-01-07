package controllers

import (
	"github.com/edusantanaw/games_match.git/internal/db"
	"github.com/edusantanaw/games_match.git/internal/entities"
	authentication "github.com/edusantanaw/games_match.git/pkg/authentication"
	"github.com/edusantanaw/games_match.git/pkg/utils/encrypter"
	utils "github.com/edusantanaw/games_match.git/pkg/utils/httpResponse"
)

type RegisterUserData struct {
	username string
	email    string
	password string
	phone    string
	Roles    []string
}

type DataResponse struct {
	token string
	user  entities.Users
}

func ResgisterUser(data *RegisterUserData) utils.HttpResponse {
	db := db.GetConnection()
	var emailAlreadyUsed *entities.Users
	db.Model(entities.Users{}).First(&emailAlreadyUsed).Where("email = ?", data.email)
	if emailAlreadyUsed != nil {
		return utils.BadRequest[string]("O email já está sendo usado")
	}
	user := &entities.Users{}
	hashedPassword, err := encrypter.Encrypter(data.password)
	if err != nil {
		return utils.BadRequest[string](err.Error())
	}
	user.ValidUser(data.username, data.email, hashedPassword, data.phone, data.Roles)
	db.Model(entities.Users{}).Create(&user)
	token, err := authentication.CreateJwtToken(user.Id, user.Roles)
	if err != nil {
		return utils.BadRequest[string](err.Error())
	}
	return utils.Created[*DataResponse](&DataResponse{user: *user, token: token})
}
