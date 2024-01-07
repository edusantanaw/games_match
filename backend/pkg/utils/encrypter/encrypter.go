package encrypter

import "golang.org/x/crypto/bcrypt"

var rounds = 10

func Encrypter(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), rounds)
	return string(hash), err
}

func CompareHash(hash string, password string) bool {
	isEqual := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return isEqual != nil
}
