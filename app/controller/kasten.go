package controller

import (
	"fmt"
	"github.com/Elbroto1993/web-ss19-w-template/app/model"
	"net/http"
)

// AddKasten controller
func Addkasten(w http.ResponseWriter, r *http.Request) {
	titel := r.FormValue("titel")
	kategorie := r.FormValue("kategorie")
	ueberkategorie := r.FormValue("ueberkategorie")
	beschreibung := r.FormValue("beschreibung")
	private := r.FormValue("private")

	// Add username from session to struct
	session, err := store.Get(r, "session")
	if err != nil {
		fmt.Println(err)
	}
	userName := session.Values["username"].(string)
	actUser, err := model.GetUserByUsername(userName)
	if err != nil {
		fmt.Println(err)
	}
	createdByUserID := actUser.Id

	kasten := model.Karteikasten{}
	kasten.Titel = titel
	kasten.Kategorie = kategorie
	kasten.Ueberkategorie = ueberkategorie
	kasten.Beschreibung = beschreibung
	kasten.Private = private
	kasten.CreatedByUserID = createdByUserID
	kasten.UserID = createdByUserID

	err := kasten.Add()
	if err != nil {
		data := struct {
			ErrorMsg string
		}{
			ErrorMsg: err.Error(),
		}
		tmpl.ExecuteTemplate(w, "register.tmpl", data)
	} else {
		tmpl.ExecuteTemplate(w, "edit2.tmpl", nil)
	}
}
