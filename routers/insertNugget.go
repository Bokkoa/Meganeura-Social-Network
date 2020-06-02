package routers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/digitalHanzo/Meganeura-Social-Network/db"
	"github.com/digitalHanzo/Meganeura-Social-Network/models"
)

//InsertNugget - Insertion process for http route
func InsertNugget(w http.ResponseWriter, r *http.Request) {
	var message models.Nugget

	//getting data from body
	err := json.NewDecoder(r.Body).Decode(&message)

	//creating object on recordnugget
	//IDUser from process token
	record := models.RecordNugget{
		UserID:  IDUser,
		Message: message.Message,
		Date:    time.Now(),
	}

	//inserting on  mongo
	_, status, err := db.InsertNugget(record)

	if err != nil {
		http.Error(w, "Sucedió un error al almacenar "+err.Error(), 400)
		return
	}

	if status == false {
		http.Error(w, "No se logró insertar el nugget", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
