package routers

import (
	"encoding/json"
	"net/http"

	"github.com/digitalHanzo/Meganeura-Social-Network/db"
	"github.com/digitalHanzo/Meganeura-Social-Network/models"
)

//ShowRelation - The http response for relations between users
func ShowRelation(w http.ResponseWriter, r *http.Request) {

	ID := r.URL.Query().Get("id")

	var m models.Relation

	m.UserID = IDUser
	m.UserRelationID = ID

	//the return of status relation
	var resp models.ResponseRelation

	status, err := db.GetRelation(m)

	if err != nil || status == false {
		resp.Status = false
	} else {
		resp.Status = true
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
