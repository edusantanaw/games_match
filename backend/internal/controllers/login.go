package controllers

import (
	"github.com/edusantanaw/games_match.git/internal/db"
	"github.com/edusantanaw/games_match.git/internal/entities"
	authentication "github.com/edusantanaw/games_match.git/pkg/authentication"
	"github.com/edusantanaw/games_match.git/pkg/utils/encrypter"
	utils "github.com/edusantanaw/games_match.git/pkg/utils/httpResponse"
)

type LoginData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func Login(data *LoginData) utils.HttpResponse {
	db := db.GetConnection()
	var user []entities.Users
	if data.Email == "" {
		return utils.BadRequest[string]("O email é invalido!")
	}
	if data.Password == "" {
		return utils.BadRequest[string]("A senha é obrigatorio!")
	}
	db.Model(entities.Users{}).Where("email = ?", data.Email).Find(&user)
	if len(user) == 0 {
		return utils.BadRequest[string]("Usuario não encontrado!")
	}
	isPassValid := encrypter.CompareHash(user[0].Password, data.Password)
	println(isPassValid)
	if !isPassValid {
		return utils.BadRequest[string]("Email/Password invalidos!")
	}
	token, _ := authentication.CreateJwtToken(user[0].Id, user[0].Roles)
	return utils.Ok[*DataResponse](&DataResponse{User: &user[0], Token: token})

}
