package routers

import (
	"encoding/json"
	"net/http"

	"github.com/digitalHanzo/Meganeura-Social-Network/db"
	"github.com/digitalHanzo/Meganeura-Social-Network/models"
)

//Register - Create users
func Register(w http.ResponseWriter, r *http.Request) {

	var u models.User

	//data streaming from body request
	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		//response for error
		http.Error(w, "Error en los datos recibidos"+err.Error(), 400)
		return
	}

	if len(u.Email) == 0 {
		//response for error
		http.Error(w, "Email requerido", 400)
		return
	}

	if len(u.Password) < 6 {
		//response for error
		http.Error(w, "La password necesita minimo 6 caracteres", 400)
		return
	}

	//Check user existence
	_, found, _ := db.CheckUserExist(u.Email)

	if found {
		http.Error(w, "El correo ya ha sido registrado", 400)
		return
	}

	//Inserting user on DB
	_, status, err := db.InsertUser(u)
	if err != nil {
		http.Error(w, "Hubo un error al insertar el usuario", 400)
		return
	}

	//Check if user has not been inserted
	if status == false {
		http.Error(w, "No se registró el usuario", 400)
		return
	}

}
