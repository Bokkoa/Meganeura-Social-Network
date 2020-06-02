package routers

import (
	"encoding/json"
	"net/http"

	"github.com/digitalHanzo/Meganeura-Social-Network/db"
	"github.com/digitalHanzo/Meganeura-Social-Network/models"
)

//EditProfile - Action of router for profile updating
func EditProfile(w http.ResponseWriter, r *http.Request) {

	var u models.User

	err := json.NewDecoder(r.Body).Decode(&u)

	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), 400)
		return
	}

	var status bool
	//IDUser is global from ProcessToken route
	status, err = db.EditUser(u, IDUser)

	if err != nil {
		http.Error(w, "Sucedio un error al modificar el usuario "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se ha modificado el registro ", 400)
		return
	}

	//the return status
	w.WriteHeader(http.StatusCreated)
}
