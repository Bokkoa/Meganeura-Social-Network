package routers

import (
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/digitalHanzo/Meganeura-Social-Network/db"
	"github.com/digitalHanzo/Meganeura-Social-Network/models"
)

//UploadAvatar - Uploading a file for avatar
func UploadAvatar(w http.ResponseWriter, r *http.Request) {

	//getting from form
	file, handler, err := r.FormFile("avatar")

	//getting the .png/.jpg
	var extension = strings.Split(handler.Filename, ".")[1]

	//setting de path file
	var fileName string = "uploads/avatars/" + IDUser + "." + extension

	//open file writeOnly and CREATE, all permissions (0666)
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, 0666)

	if err != nil {
		http.Error(w, "Error al subir la imagen ! "+err.Error(), http.StatusBadRequest)
		return
	}

	//make a copy of open file (f) and setting on path (file)
	_, err = io.Copy(f, file)

	if err != nil {
		http.Error(w, "Error al copiar la imagen ! "+err.Error(), http.StatusBadRequest)
		return
	}

	var user models.User
	var status bool

	//setting the file name on model
	user.Avatar = IDUser + "." + extension

	//saving on DB
	status, err = db.EditUser(user, IDUser)

	if err != nil || status == false {
		http.Error(w, "Error al guardar la imagen en DB ! "+err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
