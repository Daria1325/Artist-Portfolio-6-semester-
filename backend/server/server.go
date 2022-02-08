package server

import (
	"fmt"
	cnfg "github.com/daria/Portfolio/backend/config"
	"github.com/daria/Portfolio/backend/database"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
)

const layoutISO = "2006-01-02"

type Server struct {
	Repo       *database.Repo
	StatusUser bool
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
	// series, err := MainServer.Repo.GetSeries()
	// if err != nil {
	// 	fmt.Fprintf(w, err.Error())
	// 	return
	// }
	// data := struct {
	// 	Title string
	// 	Items []database.Series
	// }{
	// 	Title: "My page",
	// 	Items: series,
	// }

	t.ExecuteTemplate(w, "work", nil)
}
func showSeries(w http.ResponseWriter, r *http.Request) {

	t, err := template.ParseFiles("frontend/templates/client/show_series.html", "frontend/templates/client/header.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	vars := mux.Vars(r)
	id := vars["id"]
	pictures := []database.Picture{}
	series, err := MainServer.Repo.GetSeriesById(id)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	if series.ID != -1 {
		pictures, err = MainServer.Repo.GetPictureBySeries(id)
		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}

		data := struct {
			Title    string
			Series   database.Series
			Pictures []database.Picture
		}{
			Title:    "My page",
			Series:   series,
			Pictures: pictures,
		}

		t.ExecuteTemplate(w, "show_series", data)
	} else {
		t.Parse("<div>404 page not found</div>")
		t.Execute(w, nil)
	}

}
func showPicture(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("frontend/templates/client/show_picture.html", "frontend/templates/client/header.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	vars := mux.Vars(r)
	idP := vars["id_p"]
	idS := vars["id_s"]
	picture, err := MainServer.Repo.GetPictureById(idP)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	if picture.ID != -1 && strconv.Itoa(int(picture.SeriesId.Int32)) == idS {
		t.ExecuteTemplate(w, "show_picture", picture)
	} else {
		t.Parse("<div>404 page not found</div>")
		t.Execute(w, nil)
	}

}

func login(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, err := template.ParseFiles("frontend/templates/admin/login.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}

		t.ExecuteTemplate(w, "login", nil)
	}
	if r.Method == "POST" {
		err := godotenv.Load("Variables.env")
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		user := os.Getenv("USERNAME")
		pass := os.Getenv("PASSWORD")
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		username := r.FormValue("username")
		password := r.FormValue("password")
		if username == user && password == pass {
			http.Redirect(w, r, "/admin", http.StatusSeeOther)
			MainServer.StatusUser = true
		} else {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
		}
	}

}
func admin(w http.ResponseWriter, r *http.Request) {
	if MainServer.StatusUser {
		t, err := template.ParseFiles("frontend/templates/admin/admin.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}
		t.ExecuteTemplate(w, "admin", nil)
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

}

func adminSeries(w http.ResponseWriter, r *http.Request) {
	if MainServer.StatusUser {
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
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

}
func addSeriesHandler(w http.ResponseWriter, r *http.Request) {
	if MainServer.StatusUser {
		series := database.Series{}
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		series.Name = r.FormValue("add_series_name")
		series.Description.String = r.FormValue("add_series_description")
		err = MainServer.Repo.AddSeries(series)
		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}
		http.Redirect(w, r, "/admin/series", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

}
func editSeriesHandler(w http.ResponseWriter, r *http.Request) {
	if MainServer.StatusUser {
		series := database.Series{}
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			log.Println(err)
		}
		series.ID = id
		err = r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		series.Name = r.FormValue("edit_series_name")
		series.Description.String = r.FormValue("edit_series_description")

		err = MainServer.Repo.UpdateSeries(series)
		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}
		http.Redirect(w, r, "/admin/series", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

}
func deleteSeriesHandler(w http.ResponseWriter, r *http.Request) {
	if MainServer.StatusUser {
		vars := mux.Vars(r)
		id := vars["id"]
		err := MainServer.Repo.DeleteSeries(id)
		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}
		http.Redirect(w, r, "/admin/series", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

}

func adminPictures(w http.ResponseWriter, r *http.Request) {
	if MainServer.StatusUser {
		t, err := template.ParseFiles("frontend/templates/admin/admin_pictures.html")
		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}
		pictures, err := MainServer.Repo.GetPictures()
		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}
		data := struct {
			Title string
			Items []database.Picture
		}{
			Title: "Pictures",
			Items: pictures,
		}

		t.ExecuteTemplate(w, "admin_pictures", data)
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

}
func addPicturesHandler(w http.ResponseWriter, r *http.Request) {
	if MainServer.StatusUser {
		picture := database.Picture{}
		err := r.ParseForm()
		if err != nil {
			log.Println(err)
		}
		picture.Name = r.FormValue("add_series_name")
		picture.Description.String = r.FormValue("add_series_description")

		err = MainServer.Repo.AddPicture(picture)
		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}
		http.Redirect(w, r, "/admin/pictures", http.StatusSeeOther)
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

}
func editPicturesHandler(w http.ResponseWriter, r *http.Request) {
	if MainServer.StatusUser {

	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}
}
func deletePicturesHandler(w http.ResponseWriter, r *http.Request) {
	if MainServer.StatusUser {
		vars := mux.Vars(r)
		id := vars["id"]
		err := MainServer.Repo.DeletePictures(id)
		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}
		http.Redirect(w, r, "/admin/pictures", 301)
	} else {
		http.Redirect(w, r, "/login", http.StatusSeeOther)
	}

}

func Start(config *cnfg.Config) error {
	rtr := mux.NewRouter()
	rtr.HandleFunc("/about", about).Methods("GET")
	rtr.HandleFunc("/contact", contact).Methods("GET")
	rtr.HandleFunc("/clients", clients).Methods("GET")
	rtr.HandleFunc("/work", work).Methods("GET")
	rtr.HandleFunc("/series/{id:[0-9]+}", showSeries).Methods("GET")
	rtr.HandleFunc("/series/{id_s:[0-9]+}/{id_p:[0-9]+}", showPicture).Methods("GET")

	rtr.HandleFunc("/login", login)
	rtr.HandleFunc("/admin", admin).Methods("GET")
	rtr.HandleFunc("/admin/series", adminSeries)
	rtr.HandleFunc("/admin/pictures", adminPictures)
	rtr.HandleFunc("/admin/series/edit/{id:[0-9]+}", editSeriesHandler).Methods("POST")
	rtr.HandleFunc("/admin/series/delete/{id:[0-9]+}", deleteSeriesHandler)
	rtr.HandleFunc("/admin/series/add", addSeriesHandler).Methods("POST")
	rtr.HandleFunc("/admin/pictures/edit/{id:[0-9]+}", editPicturesHandler).Methods("POST")
	rtr.HandleFunc("/admin/pictures/delete/{id:[0-9]+}", deletePicturesHandler)
	rtr.HandleFunc("/admin/pictures/add", addPicturesHandler)

	http.Handle("/", rtr)
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./frontend/static/"))))
	http.Handle("/image/", http.StripPrefix("/image/", http.FileServer(http.Dir("./data/image/"))))

	err := http.ListenAndServe(config.BindAddr, nil)
	if err != nil {
		return err
	}
	return nil
}
