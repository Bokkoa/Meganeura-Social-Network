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

	//nuggets
	router.HandleFunc("/nugget", middlew.CheckDB(middlew.ValidateJWT(routers.InsertNugget))).Methods("POST")
	router.HandleFunc("/readnuggets", middlew.CheckDB(middlew.ValidateJWT(routers.ShowNuggets))).Methods("GET")
	router.HandleFunc("/erasenugget", middlew.CheckDB(middlew.ValidateJWT(routers.EraseNugget))).Methods("DELETE")

	//images
	//upload
	router.HandleFunc("/uploadavatar", middlew.CheckDB(middlew.ValidateJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/uploadbanner", middlew.CheckDB(middlew.ValidateJWT(routers.UploadBanner))).Methods("POST")
	//get
	router.HandleFunc("/getavatar", middlew.CheckDB(middlew.ValidateJWT(routers.GetAvatar))).Methods("GET")
	router.HandleFunc("/getbanner", middlew.CheckDB(middlew.ValidateJWT(routers.GetBanner))).Methods("GET")

	router.HandleFunc("/uprelation", middlew.CheckDB(middlew.ValidateJWT(routers.UpRelation))).Methods("POST")
	router.HandleFunc("/downrelation", middlew.CheckDB(middlew.ValidateJWT(routers.DownRelation))).Methods("DELETE")
	router.HandleFunc("/showrelation", middlew.CheckDB(middlew.ValidateJWT(routers.ShowRelation))).Methods("GET")

	router.HandleFunc("/listusers", middlew.CheckDB(middlew.ValidateJWT(routers.ListUsers))).Methods("GET")
	router.HandleFunc("/followingnuggets", middlew.CheckDB(middlew.ValidateJWT(routers.ShowNuggetsFollowing))).Methods("GET")

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
