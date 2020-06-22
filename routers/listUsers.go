package routers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/digitalHanzo/Meganeura-Social-Network/db"
)

//ListUsers - Getting users with options
func ListUsers(w http.ResponseWriter, r *http.Request) {

	if len(r.URL.Query.Get("page")) < 1 {
		http.Error(w, "Debe enviar el parámetro a la pagina", http.StatusBadRequest)
		return
	}

	typeUser := r.URL.Query().Get("type")
	page := r.URL.Query().Get("page")
	search := r.URL.Query().Get("search")

	pagTemp, err := strconv.Atoi(page)
	if err != nil {
		http.Error(w, "Hay un error en el formato de page, debe ser mayor a 0", http.StatusBadRequest)
		return
	}

	pag := int64(pagTemp)

	result, status := db.GetAllUsers(IDUser, pag, search, typeUser)

	if status == false {
		http.Error(w, "Error al leer los usuarios", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(result)

}
