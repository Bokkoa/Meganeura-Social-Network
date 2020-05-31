package routers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/digitalHanzo/Meganeura-Social-Network/db"
	"github.com/digitalHanzo/Meganeura-Social-Network/models"
)

//Email - Can be exported for other processes and not repeat itself
var Email string

//IDUser - Used for search and validate user jwt for other processes
var IDUser string

//ProcessToken - Validates the token and insert it on Model
func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	myKey := []byte("TrueSelfIsWithoutForm_Zenyatta")

	//not instance, just a pointer,
	//thats how JWT works
	claims := &models.Claim{}

	// Bearer is for standard
	splitToken := strings.Split(tk, "Bearer")

	//if not exists Bearer then
	//is not a valid token
	if len(splitToken) != 2 {
		//errors.New doesnt accept not alphanum chars
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	//delete spaces
	tk = strings.TrimSpace(splitToken[1])

	//receive the token and where gonna be stored,
	//the last one is a anon func that get a jwt token
	//and brings an interface & error

	//this is the validation
	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})

	//valid token
	if err == nil {
		//claims stored the JWT and the payload can be used
		_, found, ID := db.CheckUserExist(claims.Email)

		//user exist
		if found == true {

			//for not repeat
			Email = claims.Email
			IDUser = claims.ID.Hex()

		}

		return claims, found, IDUser, nil

	}

	if !tkn.Valid {
		return claims, false, string(""), errors.New("token Inválido")
	}

	return claims, false, string(""), err

}
