package main

import (
	"database/sql"
	"fmt"
	"net/http"

	"github.com/IceWreck/RequestBin/config"
	"github.com/IceWreck/RequestBin/controllers"
	"github.com/IceWreck/RequestBin/db"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	_ "github.com/mattn/go-sqlite3"
)

const filepath = "./test.db"

func main() {
	// Initialize the Database
	var err error
	db.GlobalDB, err = sql.Open("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
	defer db.GlobalDB.Close()
	db.CreateTable()

	// Settings
	settings := config.LoadSettings()
	credentials := map[string]string{
		settings.Username: settings.Password,
	}
	fmt.Println("Using credentials: ")
	fmt.Println(credentials)

	// Router
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Not letting chi handle it because I want it to respond to all possible methods
	r.HandleFunc("/request", controllers.RequestView)

	// Used a subrouter cause I wanted to use the basic httpauth middleware
	r.Route("/", func(r chi.Router) {
		r.Use(middleware.BasicAuth("Credentials:", credentials))
		r.Get("/", controllers.ShowView)
	})

	fmt.Println("Listing on port " + settings.Port + " ....")
	http.ListenAndServe(settings.Port, r)
}
