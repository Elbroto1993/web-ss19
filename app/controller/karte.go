package controller

import (
	"encoding/json"
	"fmt"
	"github.com/Elbroto1993/web-ss19-w-template/app/model"
	"net/http"
	"strconv"
)

// AddKarte controller
func AddOrUpdateKarte(w http.ResponseWriter, r *http.Request) {
	karte := model.Karteikarte{}
	err := json.NewDecoder(r.Body).Decode(&karte)
	if err != nil {
		fmt.Println(err)
	}
	_karteid := karte.Id
	_kastenid := karte.KastenID
	// If no _karteid in url/karte create new karte, else update karte
	if _karteid == "" {

		karte.Fach = strconv.Itoa(0)
		karte.KastenID = _kastenid

		karteid, err := karte.Add()
		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(karteid)
	} else {

		tempkarte, err := model.GetKarteById(_karteid)
		if err != nil {
			fmt.Println(err)
		}

		tempkarte.Titel = karte.Titel
		tempkarte.Frage = karte.Frage
		tempkarte.Antwort = karte.Antwort
		tempkarte.Fach = strconv.Itoa(0)

		karteid, err := tempkarte.Update()
		if err != nil {
			fmt.Println(err)
		}
		json.NewEncoder(w).Encode(karteid)
	}
}

// DeleteKarte controller
func DeleteKarte(w http.ResponseWriter, r *http.Request) {
	var _id string
	err := json.NewDecoder(r.Body).Decode(&_id)
	if err != nil {
		fmt.Println(err)
	}
	if _id != "" {
		err := model.DeleteKarte(_id)
		if err != nil {
			fmt.Println(err)
		}
	}
	json.NewEncoder(w).Encode("Kasten geloescht")
}

// AddKarte controller
func KarteRichtigOderFalsch(w http.ResponseWriter, r *http.Request) {
	karte := model.Karteikarte{}
	err := json.NewDecoder(r.Body).Decode(&karte)
	if err != nil {
		fmt.Println(err)
	}
	_karteid := karte.Id

	tempkarte, err := model.GetKarteById(_karteid)
	if err != nil {
		fmt.Println(err)
	}

	tempkarte.Fach = karte.Fach

	karteid, err := tempkarte.Update()
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(karteid)
}
