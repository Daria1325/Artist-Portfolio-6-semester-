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

var MainServer = Server{}

func about(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("frontend/templates/about.html", "frontend/templates/header.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	t.ExecuteTemplate(w, "about", nil)
}
func contact(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("frontend/templates/contact.html", "frontend/templates/header.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	t.ExecuteTemplate(w, "contact", nil)
}
func clients(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("frontend/templates/clients.html", "frontend/templates/header.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	t.ExecuteTemplate(w, "clients", nil)
}
func work(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{
		"mod": func(i, j int) int {
			return i % j
		},
	}
	t, err := template.New("test").Funcs(funcMap).ParseFiles("frontend/templates/work.html", "frontend/templates/header.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	series, err := MainServer.Repo.GetSeries()
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	data := struct {
		Title string
		Items []database.Series
	}{
		Title: "My page",
		Items: series,
	}

	t.ExecuteTemplate(w, "work", data)
}
func show_series(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("frontend/templates/show_series.html", "frontend/templates/header.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	t.ExecuteTemplate(w, "show_series", nil)
}
func show_picture(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("frontend/templates/show_picture.html", "frontend/templates/header.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	t.ExecuteTemplate(w, "show_picture", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("frontend/templates/login.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	t.ExecuteTemplate(w, "login", nil)
}
func edit(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("frontend/templates/edit_admin.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	t.ExecuteTemplate(w, "edit", nil)
}

func Start(config *cnfg.Config) error {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/about", about).Methods("GET")
	rtr.HandleFunc("/contact", contact).Methods("GET")
	rtr.HandleFunc("/clients", clients).Methods("GET")
	rtr.HandleFunc("/work", work).Methods("GET")
	rtr.HandleFunc("/series/{id:[0-9]+}", show_series).Methods("GET")
	rtr.HandleFunc("/series/{id:[0-9]+}/{id:[0-9]+}", show_picture).Methods("GET")

	rtr.HandleFunc("/login", login).Methods("GET")
	rtr.HandleFunc("/edit", edit).Methods("GET")

	http.Handle("/", rtr)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./frontend/static/"))))

	err := http.ListenAndServe(config.BindAddr, nil)
	if err != nil {
		return err
	}
	return nil
}
