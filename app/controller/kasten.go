package controller

import (
	"encoding/json"
	"fmt"
	"github.com/Elbroto1993/web-ss19-w-template/app/model"
	"net/http"
)

// AddKasten controller
func AddKasten(w http.ResponseWriter, r *http.Request) {

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

	err = json.NewDecoder(r.Body).Decode(&kasten)
	if err != nil {
		fmt.Println(err)
	}
	kasten.CreatedByUserID = createdByUserID
	kasten.UserID = createdByUserID

	err = kasten.Add()
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
