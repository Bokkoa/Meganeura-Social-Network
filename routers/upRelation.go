package routers

import (
	"net/http"

	"github.com/digitalHanzo/Meganeura-Social-Network/db"
	"github.com/digitalHanzo/Meganeura-Social-Network/models"
)

//upRelation - For NXM
func UpRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "El ID es obligatorio", http.StatusBadRequest)
		return
	}

	var m models.Relation
	m.UserID = IDUser
	m.UserRelationID = ID

	status, err := db.InsertRelation(m)

	if err != nil {
		http.Error(w, "Error al insertar relacion", http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se pudo insertar la relacion", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
