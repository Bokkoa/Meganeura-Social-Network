package routers

import (
	"net/http"

	"github.com/digitalHanzo/Meganeura-Social-Network/db"
	"github.com/digitalHanzo/Meganeura-Social-Network/models"
)

//DownRelation - Erase relation
func DownRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	var t models.Relation

	t.UserID = IDUser
	t.UserRelationID = ID
	status, err := db.DeleteRelation(t)

	if err != nil {
		http.Error(w, "Error al eliminar relacion", http.StatusBadRequest)
		return
	}

	if status == false {
		http.Error(w, "No se pudo eliminar la relacion", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
