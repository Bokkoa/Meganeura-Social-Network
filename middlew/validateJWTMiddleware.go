package middlew

import (
	"net/http"
	"github.com/digitalHanzo/Meganeura-Social-Network/routers"
)

//ValidateJWT - Middleware for true validation of JWT
func ValidateJWT( next http.HandlerFunc ) http.HandlerFunc {
	return func( w http.ResponseWriter, r *http.Request )

	//Only interested for error
	_, _, _, err := routes.ProcessToken( r.Header.Get("Authorization") )

	if err != nil{
		http.Error( w, "Error al procesar el Token ! "+err.Error(), http.StatusBadRequest)
		return
	}

	//pass middleware
	next.ServerHTTP(w, r)
}