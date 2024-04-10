package main

import (
	"log"

	"github.com/paolochmprojects/Apibasego/storage"
)

func main() {

	// Init env
	if err := loadEnv(); err != nil {
		log.Fatal(err)
	}

	// Init database connection
	storage.New()

	// instance an app Fiber
	app := newServer()

	app.Listen(":3000")
}
