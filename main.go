package main

import (
	"log"

	"github.com/digitalHanzo/Meganeura-Social-Network/db"
	"github.com/digitalHanzo/Meganeura-Social-Network/handlers"
)

func main() {

	if db.CheckConnection() == 0 {
		log.Fatal("Sin conexión a la DB")
		return
	}

	handlers.Handlers()

}
