package controllers

import (
	"fmt"
	"html/template"
	"net/http"

	"github.com/IceWreck/RequestBin/db"
)

//ShowView - View all your request logs here
func ShowView(w http.ResponseWriter, r *http.Request) {
	// r.SetBasicAuth("admin", "admin")

	t, err := template.ParseFiles("views/index.html")
	if err != nil {
		fmt.Println(err)
	}
	t.Execute(w, db.GetShowList())
}
