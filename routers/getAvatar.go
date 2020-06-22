package routers

import (
	"io"
	"net/http"
	"os"

	"github.com/digitalHanzo/Meganeura-Social-Network/db"
)

//GetAvatar - Dont neet description, is for everybody
func GetAvatar(w http.ResponseWriter, r *http.Request) {

	//getting ID from URL
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Se debe enviar el ID", http.StatusBadRequest)
		return
	}

	//getting user
	profile, err := db.SearchUser(ID)

	if err != nil {
		http.Error(w, "Usuario no encontrado", http.StatusBadRequest)
		return
	}

	//getting the user's avatar
	OpenFile, err := os.Open("uploads/avatars/" + profile.Avatar)
	if err != nil {
		http.Error(w, "Imagen no encontrada", http.StatusBadRequest)
		return
	}

	//sending image to responsewriter
	_, err = io.Copy(w, OpenFile)

	if err != nil {
		http.Error(w, "Error al copiar la imagen", http.StatusBadRequest)
	}

}
