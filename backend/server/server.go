package server

import (
	"fmt"
	cnfg "github.com/daria/Portfolio/backend/config"
	"github.com/daria/Portfolio/backend/database"
	"github.com/gorilla/mux"
	"html/template"
	"net/http"
)

type Server struct {
	Repo *database.Repo
}

func about(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("frontend/templates/about.html", "frontend/templates/header.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	t.ExecuteTemplate(w, "about", nil)
}
func clients(w http.ResponseWriter, r *http.Request) {

}
func work(w http.ResponseWriter, r *http.Request) {

}
func contact(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("frontend/templates/contact.html", "frontend/templates/header.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	t.ExecuteTemplate(w, "contact", nil)
}
func show_series(w http.ResponseWriter, r *http.Request) {

}
func show_picture(w http.ResponseWriter, r *http.Request) {

}

func Start(config *cnfg.Config) error {
	rtr := mux.NewRouter()

	rtr.HandleFunc("/about", about).Methods("GET")
	rtr.HandleFunc("/clients", clients).Methods("GET")
	rtr.HandleFunc("/series", work).Methods("GET")
	rtr.HandleFunc("/contact", contact).Methods("GET")
	rtr.HandleFunc("/series/{id:[0-9]+}", show_series).Methods("GET")
	rtr.HandleFunc("/series/{id:[0-9]+}/{id:[0-9]+}", show_picture).Methods("GET")

	http.Handle("/", rtr)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("../../../frontend/static/"))))

	err := http.ListenAndServe(config.BindAddr, rtr)
	if err != nil {
		return err
	}
	return nil
}
