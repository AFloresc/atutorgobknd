package main

import (
	"log"
	"net/http"
	"testing"

	"github.com/gorilla/mux"
)

func TestMain(t *testing.T) {
	app := Application{}

	//client := domain.Client{}

	app.Initialize()
	err := app.Client.AutoMigrate()
	if err != nil {
		log.Fatal(err)
	}
	//ctx := context.Background()

	router := mux.NewRouter()
	app.InitializeRoutes(router)
	log.Println("Listening on port 8000...")

	log.Fatal(http.ListenAndServe(":8000", router))

}
