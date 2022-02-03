package server

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	cnfg "github.com/daria/Portfolio/backend/config"
	"github.com/daria/Portfolio/backend/database"
	"github.com/gorilla/mux"
)

type Server struct {
	Repo *database.Repo
}

var MainServer = Server{}

func about(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("frontend/templates/client/about.html", "frontend/templates/client/header.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	t.ExecuteTemplate(w, "about", nil)
}
func contact(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("frontend/templates/client/contact.html", "frontend/templates/client/header.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	t.ExecuteTemplate(w, "contact", nil)
}
func clients(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("frontend/templates/client/clients.html", "frontend/templates/client/header.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	clients, err := MainServer.Repo.GetClients(6)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	t.ExecuteTemplate(w, "clients", clients)
}
func work(w http.ResponseWriter, r *http.Request) {
	funcMap := template.FuncMap{
		"mod": func(i, j int) int {
			return i % j
		},
	}
	t, err := template.New("test").Funcs(funcMap).ParseFiles("frontend/templates/client/work.html", "frontend/templates/client/header.html")
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
	t, err := template.ParseFiles("frontend/templates/client/show_series.html", "frontend/templates/client/header.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	t.ExecuteTemplate(w, "show_series", nil)
}
func show_picture(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("frontend/templates/client/show_picture.html", "frontend/templates/client/header.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	t.ExecuteTemplate(w, "show_picture", nil)
}

func login(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("frontend/templates/admin/login.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	t.ExecuteTemplate(w, "login", nil)
}
func admin(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("frontend/templates/admin/admin.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	t.ExecuteTemplate(w, "admin", nil)
}
func adminSeries(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		editType := r.FormValue("edit_type")
		fmt.Println(editType)

		http.Redirect(w, r, "/edit", 301)
	} else {
		t, err := template.ParseFiles("frontend/templates/admin/admin_series.html")
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
			Title: "Series",
			Items: series,
		}

		t.ExecuteTemplate(w, "admin_series", data)
	}
}
func adminPictures(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("frontend/templates/admin/admin_pictures.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	t.ExecuteTemplate(w, "admin_pictures", nil)
}
func editSeriesHandler(w http.ResponseWriter, r *http.Request) {

}
func deleteSeriesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	err := MainServer.Repo.DeleteSeries(id)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	http.Redirect(w, r, "/admin/series", 301)
}
func editPicturesHandler(w http.ResponseWriter, r *http.Request) {

}
func deletePicturesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	err := MainServer.Repo.DeletePictures(id)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	http.Redirect(w, r, "/admin/pictures", 301)
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
	rtr.HandleFunc("/admin", admin)
	rtr.HandleFunc("/admin/series", adminSeries)
	rtr.HandleFunc("/admin/pictures", adminPictures)
	rtr.HandleFunc("admin/series/edit/{id:[0-9]+}", editSeriesHandler)
	rtr.HandleFunc("admin/series/delete/{id:[0-9]+}", deleteSeriesHandler)
	rtr.HandleFunc("admin/pictures/edit/{id:[0-9]+}", editPicturesHandler)
	rtr.HandleFunc("admin/pictures/delete/{id:[0-9]+}", deletePicturesHandler)

	http.Handle("/", rtr)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./frontend/static/"))))

	err := http.ListenAndServe(config.BindAddr, nil)
	if err != nil {
		return err
	}
	return nil
}
