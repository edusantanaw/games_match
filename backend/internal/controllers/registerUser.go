package controllers

import (
	"github.com/edusantanaw/games_match.git/internal/db"
	"github.com/edusantanaw/games_match.git/internal/entities"
	authentication "github.com/edusantanaw/games_match.git/pkg/authentication"
	"github.com/edusantanaw/games_match.git/pkg/utils/encrypter"
	utils "github.com/edusantanaw/games_match.git/pkg/utils/httpResponse"
)

type RegisterUserData struct {
	Username string   `json:"username"`
	Email    string   `json:"email"`
	Password string   `json:"password"`
	Phone    string   `json:"phone"`
	Roles    []string `json:"roles"`
}

type DataResponse struct {
	Token string
	User  *entities.Users
}

func ResgisterUser(data *RegisterUserData) utils.HttpResponse {
	db := db.GetConnection()
	var emailAlreadyUsed []*entities.Users
	db.Model(entities.Users{}).Where("email = ?", data.Email).Find(&emailAlreadyUsed)
	if len(emailAlreadyUsed) > 0 {
		return utils.BadRequest[string]("O email já está sendo usado")
	}
	user := &entities.Users{}
	err := user.ValidUser(data.Username, data.Email, data.Password, data.Phone, data.Roles)
	if err != nil {
		return utils.BadRequest[string](err.Error())
	}
	hashedPassword, err := encrypter.Encrypter(data.Password)
	user.SetPassword(hashedPassword)
	if err != nil {
		return utils.BadRequest[string](err.Error())
	}
	db.Model(entities.Users{}).Create(&user)
	token, err := authentication.CreateJwtToken(user.Id, user.Roles)
	if err != nil {
		return utils.BadRequest[string](err.Error())
	}
	response := DataResponse{User: user, Token: token}
	return utils.Created[*DataResponse](&response)
}
