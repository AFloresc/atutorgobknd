package main

import (
	"log"
	"net/http"

	"github.com/subosito/gotenv"

	"github.com/atutor/controller"
	"github.com/atutor/domain"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

func init() {
	gotenv.Load()
}

func main() {
	app := controller.Application{
		Client: &domain.Client{},
	}

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
