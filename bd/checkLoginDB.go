package bd

import (
	"github.com/FelixMH/tuitapp/models"
	"golang.org/x/crypto/bcrypt"
)

//CheckLogin : realiza el chequeo de login a la base datos.
func CheckLogin(email string, password string) (models.Users, bool) {
	user, findUs, _ := CheckUsersTaken(email)
	if findUs == false {
		return user, false
	}
	passwordBytes := []byte(password)
	passwordDB := []byte(user.Password)

	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)
	if err != nil {
		return user, false
	}
	return user, true
}