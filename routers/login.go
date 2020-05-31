package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/digitalHanzo/Meganeura-Social-Network/db"
	"github.com/digitalHanzo/Meganeura-Social-Network/jwt"
	"github.com/digitalHanzo/Meganeura-Social-Network/models"
)

// Login - Need to explain?
func Login(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("content-type", "application/json")

	var user models.User

	err := json.NewDecoder(r.Body).Decode(&user)

	if err != nil {
		http.Error(w, "Usuario / Contraseña inválidos"+err.Error(), 400)
		return
	}

	if len(user.Email) == 0 {
		http.Error(w, "El email del usuario es requerido", 400)
		return
	}

	document, exists := db.TryLogin(user.Email, user.Password)

	if exists == false {
		http.Error(w, "Credenciales inválidas", 400)
		return
	}

	//create the jwt
	jwtKey, err := jwt.GenerateJWT(document)

	if err != nil {
		http.Error(w, "Error de token "+err.Error(), 400)
		return
	}

	//storing JWT on model
	resp := models.LoginResponse{
		Token: jwtKey,
	}

	//set jwt to headers
	w.Header().Set("Content-Type", "application/json")
	//200 OR 201
	w.WriteHeader(http.StatusCreated)
	//encoding response
	// this always return on login (the JWT )
	json.NewEncoder(w).Encode(resp)

	//make a Cookie for JWT
	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})

}
