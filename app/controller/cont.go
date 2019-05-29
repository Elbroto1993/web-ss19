package controller

import (
	"fmt"
	"github.com/Elbroto1993/web-ss19-w-template/app/model"
	"html/template"
	"net/http"
	"strconv"
)

var tmpl *template.Template

func init() {
	tmpl = template.Must(template.ParseGlob("templates/*.tmpl"))
}

func Index(w http.ResponseWriter, r *http.Request) {
	data, err := model.GetIndexData()
	if err != nil {
		fmt.Println(err)
	}

	// If user is logged in add loggedin and username to struct
	session, err := store.Get(r, "session")
	if session.Values["authenticated"] != nil && session.Values["username"] != nil {
		data.LoggedIn = strconv.FormatBool(session.Values["authenticated"].(bool))
		data.UserName = session.Values["username"].(string)
	}

	tmpl.ExecuteTemplate(w, "index.tmpl", data)
}
func Edit(w http.ResponseWriter, r *http.Request) {
	data, err := model.GetEditData()
	if err != nil {
		fmt.Println(err)
	}

	// Add username from session to struct
	session, err := store.Get(r, "session")
	if err != nil {
		fmt.Println(err)
	}
	data.UserName = session.Values["username"].(string)

	tmpl.ExecuteTemplate(w, "edit.tmpl", data)
}
func Edit2(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "edit2.tmpl", nil)
}
func Karteikasten(w http.ResponseWriter, r *http.Request) {
	kaesten, err := model.GetKarteikastenData()
	if err != nil {
		fmt.Println(err)
	}

	// If user is logged in add loggedin and username to struct
	session, err := store.Get(r, "session")
	if session.Values["authenticated"] != nil && session.Values["username"] != nil {
		kaesten.LoggedIn = strconv.FormatBool(session.Values["authenticated"].(bool))
		kaesten.UserName = session.Values["username"].(string)
	}

	tmpl.ExecuteTemplate(w, "karteikasten.tmpl", kaesten)
}
func Lern(w http.ResponseWriter, r *http.Request) {
	_id := r.FormValue("_kastenid")
	data, err := model.GetLernData(_id)
	if err != nil {
		fmt.Println(err)
	}

	// Add username from session to struct
	session, err := store.Get(r, "session")
	if err != nil {
		fmt.Println(err)
	}
	data.UserName = session.Values["username"].(string)

	tmpl.ExecuteTemplate(w, "lern.tmpl", data)
}
func Lern2(w http.ResponseWriter, r *http.Request) {
	karteid := r.FormValue("_karteid")
	kastenid := r.FormValue("_kastenid")
	data, err := model.GetLern2Data(kastenid, karteid)
	if err != nil {
		fmt.Println(err)
	}

	// Add username from session to struct
	session, err := store.Get(r, "session")
	if err != nil {
		fmt.Println(err)
	}
	data.UserName = session.Values["username"].(string)

	tmpl.ExecuteTemplate(w, "lern2.tmpl", data)
}
func Meinekarteien(w http.ResponseWriter, r *http.Request) {
	kaesten, err := model.GetKarteikastenData()
	if err != nil {
		fmt.Println(err)
	}
	tmpl.ExecuteTemplate(w, "meinekarteien.tmpl", kaesten)
}
func Profil(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "profil.tmpl", nil)
}
func Register(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "register.tmpl", nil)
}
func View(w http.ResponseWriter, r *http.Request) {
	kastenid := r.FormValue("_kastenid")
	karteid := r.FormValue("_karteid")
	viewData, err := model.GetViewData(kastenid, karteid)
	if err != nil {
		fmt.Println(err)
	}

	// If user is logged in add loggedin and username to struct
	session, err := store.Get(r, "session")
	if session.Values["authenticated"] != nil && session.Values["username"] != nil {
		viewData.LoggedIn = strconv.FormatBool(session.Values["authenticated"].(bool))
		viewData.UserName = session.Values["username"].(string)
	}

	tmpl.ExecuteTemplate(w, "view.tmpl", viewData)
}
