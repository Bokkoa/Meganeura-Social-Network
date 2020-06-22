package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/digitalHanzo/Meganeura-Social-Network/db"
)

//ShowNuggetsFollowing - The http for get all nuggets of my follows
func ShowNuggetsFollowing(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Debe enviar el parámetro a la pagina", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Hay un error en el formato de page, debe ser mayor a 0", http.StatusBadRequest)
		return
	}

	response, correct := db.GetNuggetsFollowing(IDUser, page)

	if correct == false {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
