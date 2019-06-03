package controller

import (
	"encoding/json"
	"fmt"
	"github.com/Elbroto1993/web-ss19-w-template/app/model"
	"net/http"
)

// AddKasten controller
func AddOrUpdateKasten(w http.ResponseWriter, r *http.Request) {
	kasten := model.Karteikasten{}

	err := json.NewDecoder(r.Body).Decode(&kasten)
	if err != nil {
		fmt.Println(err)
	}

	_kastenid := kasten.Id
	// If no _kastenid in url create new kasten, else update kasten
	if _kastenid == "" {
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

		kasten.CreatedByUserID = createdByUserID
		kasten.UserID = createdByUserID

		kastenid, err := kasten.Add()
		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(kastenid)
	} else {

		tempKasten, err := model.GetKastenById(_kastenid)
		if err != nil {
			fmt.Println(err)
		}

		tempKasten.Titel = kasten.Titel
		tempKasten.Kategorie = kasten.Kategorie
		tempKasten.Ueberkategorie = kasten.Ueberkategorie
		tempKasten.Beschreibung = kasten.Beschreibung
		tempKasten.Private = kasten.Private

		kastenid, err := tempKasten.Update()
		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(kastenid)
	}

}
