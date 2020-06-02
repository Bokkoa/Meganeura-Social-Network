package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/digitalHanzo/Meganeura-Social-Network/middlew"
	"github.com/digitalHanzo/Meganeura-Social-Network/routers"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Handlers - func for list routes
func Handlers() {

	router := mux.NewRouter()

	//routes
	router.HandleFunc("/register", middlew.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckDB(routers.Login)).Methods("POST")
	// nested middleware funcs as two middlewares
	router.HandleFunc("/showprofile", middlew.CheckDB(middlew.ValidateJWT(routers.ShowProfile))).Methods("GET")

	router.HandleFunc("/editprofile", middlew.CheckDB(middlew.ValidateJWT(routers.EditProfile))).Methods("PUT")

	router.HandleFunc("/nugget", middlew.CheckDB(middlew.ValidateJWT(routers.InsertNugget))).Methods("POST")
	router.HandleFunc("/readnuggets", middlew.CheckDB(middlew.ValidateJWT(routers.ShowNuggets))).Methods("GET")

	//get by env
	PORT := os.Getenv("PORT")

	if PORT == "" {
		PORT = "9090"
	}

	//setting cors
	handler := cors.AllowAll().Handler(router)

	//server mount on port "localhost:9090"
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
