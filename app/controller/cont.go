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
	userName := ""
	var loggedIn string
	// If user is logged in add loggedin and username to struct
	session, err := store.Get(r, "session")
	if session.Values["authenticated"] != nil && session.Values["username"] != nil {
		loggedIn = strconv.FormatBool(session.Values["authenticated"].(bool))
		userName = session.Values["username"].(string)
	}

	data, err := model.GetIndexData(userName)
	if err != nil {
		fmt.Println(err)
	}

	data.LoggedIn = loggedIn
	data.UserName = userName

	tmpl.ExecuteTemplate(w, "index.tmpl", data)
}
func Edit(w http.ResponseWriter, r *http.Request) {
	kastenid := r.FormValue("_kastenid")

	// Add username from session to struct
	session, err := store.Get(r, "session")
	if err != nil {
		fmt.Println(err)
	}
	username := session.Values["username"].(string)
	data, err := model.GetEditData(username, kastenid)
	if err != nil {
		fmt.Println(err)
	}
	data.UserName = session.Values["username"].(string)

	tmpl.ExecuteTemplate(w, "edit.tmpl", data)
}
func Edit2(w http.ResponseWriter, r *http.Request) {
	kastenid := r.FormValue("_kastenid")
	karteid := r.FormValue("_karteid")

	session, err := store.Get(r, "session")
	if err != nil {
		fmt.Println(err)
	}
	userName := session.Values["username"].(string)

	data, err := model.GetEdit2Data(kastenid, karteid, userName)
	if err != nil {
		fmt.Println(err)
	}

	tmpl.ExecuteTemplate(w, "edit2.tmpl", data)
}
func Karteikasten(w http.ResponseWriter, r *http.Request) {
	kategorie := r.FormValue("_kategorie")

	userName := ""
	var loggedIn string
	// If user is logged in add loggedin and username to struct
	session, err := store.Get(r, "session")
	if session.Values["authenticated"] != nil && session.Values["username"] != nil {
		loggedIn = strconv.FormatBool(session.Values["authenticated"].(bool))
		userName = session.Values["username"].(string)
	}

	kaesten, err := model.GetKarteikastenData(userName, kategorie)
	if err != nil {
		fmt.Println(err)
	}

	kaesten.UserName = userName
	kaesten.LoggedIn = loggedIn

	tmpl.ExecuteTemplate(w, "karteikasten.tmpl", kaesten)
}
func Lern(w http.ResponseWriter, r *http.Request) {
	// Add username from session to struct
	session, err := store.Get(r, "session")
	if err != nil {
		fmt.Println(err)
	}
	userName := session.Values["username"].(string)

	_kastenid := r.FormValue("_kastenid")
	_karteid := r.FormValue("_karteid")
	data, err := model.GetLernData(_kastenid, _karteid, userName)
	if err != nil {
		fmt.Println(err)
	}

	data.UserName = userName

	tmpl.ExecuteTemplate(w, "lern.tmpl", data)
}
func Lern2(w http.ResponseWriter, r *http.Request) {
	karteid := r.FormValue("_karteid")
	kastenid := r.FormValue("_kastenid")

	// Add username from session to struct
	session, err := store.Get(r, "session")
	if err != nil {
		fmt.Println(err)
	}
	userName := session.Values["username"].(string)

	data, err := model.GetLern2Data(kastenid, karteid, userName)
	if err != nil {
		fmt.Println(err)
	}

	data.UserName = userName

	tmpl.ExecuteTemplate(w, "lern2.tmpl", data)
}
func Meinekarteien(w http.ResponseWriter, r *http.Request) {
	kategorie := r.FormValue("_kategorie")

	// Add username from session to struct
	session, err := store.Get(r, "session")
	if err != nil {
		fmt.Println(err)
	}
	userName := session.Values["username"].(string)

	kaesten, err := model.GetMeineKarteienData(userName, kategorie)
	if err != nil {
		fmt.Println(err)
	}

	kaesten.UserName = userName

	tmpl.ExecuteTemplate(w, "meinekarteien.tmpl", kaesten)
}
func Profil(w http.ResponseWriter, r *http.Request) {

	session, err := store.Get(r, "session")
	if err != nil {
		fmt.Println(err)
	}
	userName := session.Values["username"].(string)

	data, err := model.GetProfilData(userName)
	if err != nil {
		fmt.Println(err)
	}

	tmpl.ExecuteTemplate(w, "profil.tmpl", data)
}
func Register(w http.ResponseWriter, r *http.Request) {
	data, err := model.GetRegisterData()
	if err != nil {
		fmt.Println(err)
	}

	tmpl.ExecuteTemplate(w, "register.tmpl", data)
}
func View(w http.ResponseWriter, r *http.Request) {
	kastenid := r.FormValue("_kastenid")
	karteid := r.FormValue("_karteid")

	userName := ""
	var loggedIn string
	// If user is logged in add loggedin and username to struct
	session, err := store.Get(r, "session")
	if session.Values["authenticated"] != nil && session.Values["username"] != nil {
		loggedIn = strconv.FormatBool(session.Values["authenticated"].(bool))
		userName = session.Values["username"].(string)
	}

	viewData, err := model.GetViewData(kastenid, karteid, userName)
	if err != nil {
		fmt.Println(err)
	}

	viewData.LoggedIn = loggedIn
	viewData.UserName = userName

	tmpl.ExecuteTemplate(w, "view.tmpl", viewData)
}
