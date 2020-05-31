package jwt

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/digitalHanzo/Meganeura-Social-Network/models"
)

//GenerateJWT - returns an string with user as payload
func GenerateJWT(u models.User) (string, error) {

	myKey := []byte("TrueSelfIsWithoutForm_Zenyatta")

	//make the payload, first the privileges
	payload := jwt.MapClaims{
		"email":     u.Email,
		"name":      u.Name,
		"lastName":  u.LastName,
		"birthDate": u.BirthDate,
		"bio":       u.Bio,
		"location":  u.Location,
		"web":       u.Web,
		"_id":       u.ID.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}
	//expirates on 24 hrs as Unix format

	//creating token with sign mode & payload
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	//signing string
	tokenStr, err := token.SignedString(myKey)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
