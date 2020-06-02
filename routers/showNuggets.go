package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/digitalHanzo/Meganeura-Social-Network/db"
)

//ShowNuggets - Method of get for router
func ShowNuggets(w http.ResponseWriter, r *http.Request) {

	//getting id from URL
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parametro ID", http.StatusBadRequest)
		return
	}

	//check for pagination
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Debe enviar el parametro page", http.StatusBadRequest)
		return
	}

	//casting string to INT
	page, err := strconv.Atoi(r.URL.Query().Get("page"))

	if err != nil {
		http.Error(w, "Fallo al castear o page es menor a 0", http.StatusBadRequest)
		return
	}

	//casting to INT
	pag := int64(page)

	response, correct := db.GetNuggets(ID, pag)

	//bad db read
	if correct == false {
		http.Error(w, "Error al leer los nuggets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

	//sending data
	json.NewEncoder(w).Encode(response)
}
