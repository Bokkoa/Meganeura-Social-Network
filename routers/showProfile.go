package routers

import (
	"encoding/json"
	"net/http"

	"github.com/digitalHanzo/Meganeura-Social-Network/db"
)

//ShowProfile - Access to user data
func ShowProfile(w http.ResponseWriter, r *http.Request) {
	//get id by url
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "El parametro ID está vacío", http.StatusBadRequest)
		return
	}

	profile, err := db.SearchProfile(ID)

	if err != nil {
		http.Error(w, "Ocurrio un error buscando el perfil"+err.Error(), 400)
		return
	}

	w.Header().Set("context-type", "application/json")
	//good status 201
	w.WriteHeader(http.StatusCreated)

	//this returns, the profile encoded
	json.NewEncoder(w).Encode(profile)
}
