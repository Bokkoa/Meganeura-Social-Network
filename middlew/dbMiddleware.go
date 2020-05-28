package middlew

import (
	"net/http"

	"github.com/digitalHanzo/Meganeura-Social-Network/db"
)

//CheckDB - middleware
func CheckDB(next http.HandlerFunc) http.HandlerFunc {

	//action
	return func(w http.ResponseWriter, r *http.Request) {

		if db.CheckConnection() == 0 {
			http.Error(w, "Conexión con DB perdida", 500)

			return
		}

		// if all good the http res
		// and req is sending to the next process
		next.ServeHTTP(w, r)

	}

}
