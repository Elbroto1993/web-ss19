package model

import (
	"encoding/json"
	"fmt"
	"github.com/mitchellh/mapstructure"

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

// Delete Kasten by Id
func DeleteKasten(_id string) (err error) {

	allKarten, err := GetAllKarten()
	if err != nil {
		return err
	}
	var decodeKarten []Karteikarte
	mapstructure.Decode(allKarten, &decodeKarten)
	// Add _id to decodeKarten, because mapstructure.Decode doesn't do it
	index := 0
	for _, v := range allKarten {
		decodeKarten[index].Id = v["_id"].(string)
		index++
	}
	// Delete all karten from kasten
	for i := 0; i < len(decodeKarten); i++ {
		if decodeKarten[i].KastenID == _id {
			btDB.Delete(decodeKarten[i].Id)
		}
	}

	// Delete kasten from DB
	err = btDB.Delete(_id)
	if err != nil {
		fmt.Println(err)
	}

	return err
}

// Delete Kasten by Id
func DeleteKastenWithProfile(username string) (err error) {

	user, err := GetUserByUsername(username)
	if err != nil {
		return err
	}
	createdByUserId := user.Id

	allKasten, err := GetAllKasten()
	if err != nil {
		return err
	}

	var decodeKasten []Karteikasten
	mapstructure.Decode(allKasten, &decodeKasten)
	// Add _id to decodeKasten, because mapstructure.Decode doesn't do it
	index := 0
	for _, v := range allKasten {
		decodeKasten[index].Id = v["_id"].(string)
		index++
	}

	for j := 0; j < len(decodeKasten); j++ {
		if createdByUserId == decodeKasten[j].CreatedByUserID {
			allKarten, err := GetAllKarten()
			if err != nil {
				return err
			}
			var decodeKarten []Karteikarte
			mapstructure.Decode(allKarten, &decodeKarten)
			// Add _id to decodeKarten, because mapstructure.Decode doesn't do it
			index = 0
			for _, v := range allKarten {
				decodeKarten[index].Id = v["_id"].(string)
				index++
			}
			// Delete all karten from kasten
			for i := 0; i < len(decodeKarten); i++ {
				if decodeKarten[i].KastenID == decodeKasten[j].Id {
					btDB.Delete(decodeKarten[i].Id)
				}
			}

			// Delete kasten from DB
			err = btDB.Delete(decodeKasten[j].Id)
			if err != nil {
				fmt.Println(err)
			}
		}
	}

	return err
}

// GetKastenById ...
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
