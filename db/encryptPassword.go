package db

import "golang.org/x/crypto/bcrypt"

// EncryptPassword - User for secure reasons
func EncryptPassword(pass string) (string, error) {

	cost := 8

	//only accept bytes slice
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)

	//return the encrypted password
	return string(bytes), err
}
