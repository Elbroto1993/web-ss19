package model

import (
	"encoding/json"
	"fmt"

	couchdb "github.com/leesper/couchdb-golang"
)

// Karteikarte Struct
type Karteikarte struct {
	Id       string `json:"_id"`
	Rev      string `json:"_rev"`
	Type     string `json:"type"`
	KastenID string `json:"kastenid"`
	Titel    string `json:"titel"`
	Frage    string `json:"frage"`
	Antwort  string `json:"antwort"`
	Fach     string `json:"fach"`
	couchdb.Document
}

// Add Karte
func (karteikarte Karteikarte) Add() (string, error) {

	karteikarte.Type = "Karteikarte"

	// Convert Karteikasten struct to map[string]interface as required by Save() method
	k, err := karteikarte2Map(karteikarte)

	// Delete _id and _rev from map, otherwise DB access will be denied (unauthorized)
	delete(k, "_id")
	delete(k, "_rev")

	// Add karteikasten to DB

	karteid, _, err := btDB.Save(k, nil)

	if err != nil {
		fmt.Printf("[Add] error: %s", err)
	}

	return karteid, err
}

// Update Karte
func (karteikarte Karteikarte) Update() (string, error) {

	karteikarte.Type = "Karteikarte"

	// Convert karteikarte struct to map[string]interface as required by Save() method
	k, err := karteikarte2Map(karteikarte)

	// Add karteikarte to DB

	err = btDB.Set(karteikarte.Id, k)

	if err != nil {
		fmt.Printf("[Add] error: %s", err)
	}

	return karteikarte.Id, err
}

// Delete Karte by Id
func DeleteKarte(_id string) (err error) {

	// Delete karte from DB
	err = btDB.Delete(_id)
	if err != nil {
		fmt.Println(err)
	}

	return err
}

func GetKarteById(karteid string) (Karteikarte, error) {
	query := `
	{
		"selector": {
			 "type": "Karteikarte",
			 "_id": "%s"
		}
	}`
	k, err := btDB.QueryJSON(fmt.Sprintf(query, karteid))
	if err != nil || len(k) != 1 {
		return Karteikarte{}, err
	}

	karte, err := map2Karteikarte(k[0])
	if err != nil {
		return Karteikarte{}, err
	}

	return karte, nil
}

// ---------------------------------------------------------------------------
// Internal helper functions
// ---------------------------------------------------------------------------

// Convert from User struct to map[string]interface{} as required by golang-couchdb methods
func karteikarte2Map(k Karteikarte) (karteikarte map[string]interface{}, err error) {
	kJSON, err := json.Marshal(k)
	json.Unmarshal(kJSON, &karteikarte)

	return karteikarte, err
}

// Convert from map[string]interface{} to User struct as required by golang-couchdb methods
func map2Karteikarte(karteikarte map[string]interface{}) (k Karteikarte, err error) {
	kJSON, err := json.Marshal(karteikarte)
	json.Unmarshal(kJSON, &k)

	return k, err
}
