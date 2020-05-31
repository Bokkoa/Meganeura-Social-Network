package db

import (
	"github.com/digitalHanzo/Meganeura-Social-Network/models"
	"golang.org/x/crypto/bcrypt"
)

// TryLogin - Check if its a successful login
func TryLogin(email string, password string) (models.User, bool) {

	usr, found, _ := CheckUserExist(email)

	if found == false {
		return usr, false
	}

	//storing passwords for match
	passwordBytes := []byte(password)
	passwordDB := []byte(usr.Password)

	//comparing
	err := bcrypt.CompareHashAndPassword(passwordDB, passwordBytes)

	if err != nil {
		return usr, false
	}

	//if all good
	return usr, true
}
