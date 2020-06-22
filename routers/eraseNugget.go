package routers

import (
	"net/http"

	"github.com/digitalHanzo/Meganeura-Social-Network/db"
)

//EraseNugget - Router eraser
func EraseNugget(w http.ResponseWriter, r *http.Request) {

	//getting id from url
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviarse el parametro id", http.StatusBadRequest)
		return
	}

	//erase with IDUser from token
	err := db.DeleteNugget(ID, IDUser)

	if err != nil {
		http.Error(w, "Ocurrio un error al eliminar"+err.Error(), http.StatusBadRequest)
		return

	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
