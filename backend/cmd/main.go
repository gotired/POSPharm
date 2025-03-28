package main

import (
	"log"

	"github.com/gotired/POSPharm/config"
	"github.com/gotired/POSPharm/infrastructure"
)

func main() {

	// load env
	config.LoadEnv()

	// setup database
	db := config.InitDB()

	app := infrastructure.NewFiberApp(db)

	log.Fatal(app.Listen(":8080"))
}
