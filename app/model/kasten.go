package model

import (
	"encoding/json"
	"fmt"

	couchdb "github.com/leesper/couchdb-golang"
)

// Karteikasten Struct
type Karteikasten struct {
	Id              string `json:"_id"`
	Rev             string `json:"_rev"`
	Type            string `json:"type"`
	Kategorie       string `json:"kategorie"`
	Titel           string `json:"titel"`
	Beschreibung    string `json:"beschreibung"`
	Private         string `json:"private"`
	CreatedByUserID string `json:"createdByUserId"`
	UserID          string `json:"userid"`
	Ueberkategorie  string `json:"ueberkategorie"`
	AnzKarten       string `json:"anzkarten"`
	Fortschritt     string `json:"fortschritt"`
	couchdb.Document
}

// Add Kasten
func (karteikasten Karteikasten) Add() (kastenid string, err error) {

	karteikasten.Type = "Karteikasten"

	// Convert Karteikasten struct to map[string]interface as required by Save() method
	k, err := karteikasten2Map(karteikasten)

	// Delete _id and _rev from map, otherwise DB access will be denied (unauthorized)
	delete(k, "_id")
	delete(k, "_rev")

	// Add karteikasten to DB

	kastenid, _, err = btDB.Save(k, nil)

	if err != nil {
		fmt.Printf("[Add] error: %s", err)
	}

	return kastenid, err
}

// Update Kasten
func (karteikasten Karteikasten) Update() (string, error) {

	karteikasten.Type = "Karteikasten"

	// Convert Karteikasten struct to map[string]interface as required by Save() method
	k, err := karteikasten2Map(karteikasten)

	// Add karteikasten to DB

	err = btDB.Set(karteikasten.Id, k)

	if err != nil {
		fmt.Printf("[Add] error: %s", err)
	}

	return karteikasten.Id, err
}

func GetKastenById(kastenid string) (Karteikasten, error) {
	query := `
	{
		"selector": {
			 "type": "Karteikasten",
			 "_id": "%s"
		}
	}`
	k, err := btDB.QueryJSON(fmt.Sprintf(query, kastenid))
	if err != nil || len(k) != 1 {
		return Karteikasten{}, err
	}

	kasten, err := map2Karteikasten(k[0])
	if err != nil {
		return Karteikasten{}, err
	}

	return kasten, nil
}

// ---------------------------------------------------------------------------
// Internal helper functions
// ---------------------------------------------------------------------------

// Convert from User struct to map[string]interface{} as required by golang-couchdb methods
func karteikasten2Map(k Karteikasten) (karteikasten map[string]interface{}, err error) {
	kJSON, err := json.Marshal(k)
	json.Unmarshal(kJSON, &karteikasten)

	return karteikasten, err
}

// Convert from map[string]interface{} to User struct as required by golang-couchdb methods
func map2Karteikasten(karteikasten map[string]interface{}) (k Karteikasten, err error) {
	kJSON, err := json.Marshal(karteikasten)
	json.Unmarshal(kJSON, &k)

	return k, err
}
