package authentication

import (
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

var secret = os.Getenv("JWT_SECRET")

func CreateJwtToken(userId string, role []string) (string, error) {
	atClaims := jwt.MapClaims{}
	atClaims["authorized"] = true
	atClaims["user_id"] = userId
	atClaims["roles"] = role
	at := jwt.NewWithClaims(jwt.SigningMethodHS256, atClaims)
	token, err := at.SignedString([]byte(secret))
	return token, err
}

func ExtractToken(r *gin.Context) (string, error) {
	bearToken := r.Request.Header.Get("Authorization")
	splitedToken := strings.Split(bearToken, " ")
	if splitedToken[0] != "Bearer" || len(splitedToken) != 2 {
		return "", errors.New("Token invalido")
	}
	return splitedToken[1], nil
}

func TokenValid(r *gin.Context) error {
	token, err := VerifyToken(r)
	if err != nil {
		return err
	}
	if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
		return err
	}
	return nil
}

func ExtractTokenMetadata(r *gin.Context) (string, error) {
	token, err := VerifyToken(r)
	if err != nil {
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userId := fmt.Sprint(claims["user_id"])
		return userId, nil
	}
	return "", err
}

func VerifyToken(r *gin.Context) (*jwt.Token, error) {
	tokenString, err := ExtractToken(r)
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("ACCESS_SECRET")), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}
