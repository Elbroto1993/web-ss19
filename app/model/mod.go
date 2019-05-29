package model

import (
	couchdb "github.com/leesper/couchdb-golang"
	"github.com/mitchellh/mapstructure"
	"math"
	"math/rand"
	"strconv"
	// "time"
)

// User Struct
// type User struct {
// 	Id        string    `json:"_id"`
// 	Rev       string    `json:"_rev"`
// 	Username  string    `json:"username"`
// 	Password  string    `json:"password"`
// 	LoggedIn  string    `json:"loggedin"`
// 	Email     string    `json:"email"`
// 	CreatedAt time.Time `json:"createdat"`
// 	couchdb.Document
// }

// Karteikarte Struct
type Karteikarte struct {
	Id       string `json:"_id"`
	Rev      string `json:"_rev"`
	KastenID string `json:"kastenid"`
	Titel    string `json:"titel"`
	Frage    string `json:"frage"`
	Antwort  string `json:"antwort"`
	Fach     string `json:"fach"`
	couchdb.Document
}

// IndexData Struct
type IndexData struct {
	AnzUser   string `json:"anzuser"`
	AnzKasten string `json:"anzkasten"`
	AnzKarten string `json:"anzkarten"`
	LoggedIn  string `json:"loggedin"`
	UserName  string `json:"username"`
}

// KarteikastenData Struct
type KarteikastenData struct {
	Id        string         `json:"id"`
	LoggedIn  string         `json:"loggedin"`
	UserName  string         `json:"username"`
	AnzKarten string         `json:"anzkarten"`
	Kaesten   []Karteikasten `json:"kaesten"`
	// Kategorie       string `json:"kategorie"`
	// Titel           string `json:"titel"`
	// Beschreibung    string `json:"beschreibung"`
	// Private         string `json:"private"`
	// CreatedByUserID string `json:"createdByUserId"`
	// UserID          string `json:"userid"`
	// Ueberkategorie  string `json:"ueberkategorie"`
}

// ViewData Struct
type ViewData struct {
	Kategorie       string        `json:"kategorie"`
	Titel           string        `json:"titel"`
	Beschreibung    string        `json:"beschreibung"`
	Fortschritt     string        `json:"fortschritt"`
	Private         string        `json:"private"`
	CreatedByUserID string        `json:"createdByUserId"`
	UserID          string        `json:"userid"`
	Ueberkategorie  string        `json:"ueberkategorie"`
	AnzKarten       string        `json:"anzkarten"`
	UserName        string        `json:"username"`
	LoggedIn        string        `json:"loggedin"`
	SelectedKarte   Karteikarte   `json:"selectedkarte"`
	Karten          []Karteikarte `json:"karten"`
}

// LernData Struct
type LernData struct {
	Kategorie       string      `json:"kategorie"`
	Titel           string      `json:"titel"`
	Beschreibung    string      `json:"beschreibung"`
	Fortschritt     string      `json:"fortschritt"`
	Private         string      `json:"private"`
	CreatedByUserID string      `json:"createdByUserId"`
	UserID          string      `json:"userid"`
	Ueberkategorie  string      `json:"ueberkategorie"`
	AnzKarten       string      `json:"anzkarten"`
	AnzFachZero     string      `json:"anzfachzero"`
	AnzFachOne      string      `json:"anzfachone"`
	AnzFachTwo      string      `json:"anzfachtwo"`
	AnzFachThree    string      `json:"anzfachthree"`
	AnzFachFour     string      `json:"anzfachfour"`
	UserName        string      `json:"username"`
	Karte           Karteikarte `json:"karte"`
}

// EditData Struct
type EditData struct {
	UserName string `json:"username"`
}

// CouchDB Connection
var btDB *couchdb.Database

func init() {
	var err error
	btDB, err = couchdb.NewDatabase("http://localhost:5984/braintrain")
	if err != nil {
		panic(err)
	}
}

// GetAllKasten , helper function
func GetAllKasten() ([]map[string]interface{}, error) {
	allKasten, err := btDB.QueryJSON(`
	{
		"selector": {
			"type": {
				"$eq": "Karteikasten"
			}
		}
	}
	`)
	if err != nil {
		return nil, err
	}
	return allKasten, nil
}

// GetAllKarten , helper function
func GetAllKarten() ([]map[string]interface{}, error) {
	allKarten, err := btDB.QueryJSON(`
	{
		"selector": {
			"type": {
				"$eq": "Karteikarte"
			}
		}
	}
	`)
	if err != nil {
		return nil, err
	}
	return allKarten, nil
}

// GetAllUser , helper function
func GetAllUser() ([]map[string]interface{}, error) {
	allUser, err := btDB.QueryJSON(`
	{
		"selector": {
			"type": {
				"$eq": "User"
			}
		}
	}
	`)
	if err != nil {
		return nil, err
	}
	return allUser, nil
}

// GetIndexData ...
func GetIndexData() (IndexData, error) {
	allUser, err := GetAllUser()
	if err != nil {
		return IndexData{}, err
	}
	allKasten, err := GetAllKasten()
	if err != nil {
		return IndexData{}, err
	}
	allKarten, err := GetAllKarten()
	if err != nil {
		return IndexData{}, err
	}
	index := IndexData{
		AnzUser:   strconv.Itoa(len(allUser)),
		AnzKasten: strconv.Itoa(len(allKasten)),
		AnzKarten: strconv.Itoa(len(allKarten)),
	}
	return index, nil
}

// GetKarteikastenData ...
func GetKarteikastenData() (KarteikastenData, error) {
	// Get all Kaesten and decode them
	allKasten, err := GetAllKasten()
	if err != nil {
		return KarteikastenData{}, err
	}
	var decodedKaesten []Karteikasten
	mapstructure.Decode(allKasten, &decodedKaesten)
	// Add _id to decodedKaesten, because mapstructure.Decode doesn't do it
	index := 0
	for _, v := range allKasten {
		decodedKaesten[index].Id = v["_id"].(string)
		index++
	}

	// Get all Karten and decode them
	allKarten, err := GetAllKarten()
	if err != nil {
		return KarteikastenData{}, err
	}
	var decodedKarten []Karteikarte
	mapstructure.Decode(allKarten, &decodedKarten)
	// Fill AnzKarten from Karteikasten with values
	for i := 0; i < len(decodedKaesten); i++ {
		countKarten := 0
		for j := 0; j < len(decodedKarten); j++ {
			if string(decodedKaesten[i].Id) == decodedKarten[j].KastenID {
				countKarten++
			}
		}
		decodedKaesten[i].AnzKarten = strconv.Itoa(countKarten)
	}
	var retValue KarteikastenData
	retValue.Kaesten = decodedKaesten
	return retValue, nil
}

// GetViewData ...
func GetViewData(kastenid string, karteid string) (ViewData, error) {
	kasten, err := btDB.Get(kastenid, nil)
	if err != nil {
		return ViewData{}, err
	}
	var decodeKasten Karteikasten
	mapstructure.Decode(kasten, &decodeKasten)

	// Only if a karte was selected
	var decodeKarte Karteikarte
	if karteid != "" {
		karte, err := btDB.Get(karteid, nil)
		if err != nil {
			return ViewData{}, err
		}
		mapstructure.Decode(karte, &decodeKarte)
		decodeKarte.Id = karte["_id"].(string)
	} else {
		decodeKarte = Karteikarte{}
	}

	allKarten, err := GetAllKarten()
	if err != nil {
		return ViewData{}, err
	}
	var decodeKarten []Karteikarte
	var returnKarten []Karteikarte

	anzKarten := 0
	mapstructure.Decode(allKarten, &decodeKarten)
	// Add _id to decodeKarten, because mapstructure.Decode doesn't do it
	index := 0
	for _, v := range allKarten {
		decodeKarten[index].Id = v["_id"].(string)
		index++
	}
	// Filter all karten by kastenid
	for i := 0; i < len(decodeKarten); i++ {
		if decodeKarten[i].KastenID == kastenid {
			returnKarten = append(returnKarten, decodeKarten[i])
			anzKarten++
		}
	}

	viewData := ViewData{
		Kategorie:       decodeKasten.Kategorie,
		Titel:           decodeKasten.Titel,
		Beschreibung:    decodeKasten.Beschreibung,
		Fortschritt:     getFortschritt(returnKarten),
		Private:         decodeKasten.Private,
		CreatedByUserID: decodeKasten.CreatedByUserID,
		UserID:          decodeKasten.UserID,
		Ueberkategorie:  decodeKasten.Ueberkategorie,
		AnzKarten:       strconv.Itoa(anzKarten),
		SelectedKarte:   decodeKarte,
		Karten:          returnKarten,
	}
	return viewData, nil
}

// GetLernData ...
func GetLernData(_id string) (LernData, error) {
	kasten, err := btDB.Get(_id, nil)
	if err != nil {
		return LernData{}, err
	}
	var decodeKasten Karteikasten
	mapstructure.Decode(kasten, &decodeKasten)
	allKarten, err := GetAllKarten()
	if err != nil {
		return LernData{}, err
	}
	var decodeKarten []Karteikarte
	var returnKarten []Karteikarte
	anzKarten := 0
	mapstructure.Decode(allKarten, &decodeKarten)
	// Add _id to decodeKarten, because mapstructure.Decode doesn't do it
	index := 0
	for _, v := range allKarten {
		decodeKarten[index].Id = v["_id"].(string)
		index++
	}
	// Filter allKarten and only get the karten that belong to the kasten
	for i := 0; i < len(decodeKarten); i++ {
		if decodeKarten[i].KastenID == _id {
			returnKarten = append(returnKarten, decodeKarten[i])
			anzKarten++
		}
	}
	// Calculate count for each fach
	var fach0, fach1, fach2, fach3, fach4 int
	fach1 = 0
	for i := 0; i < len(returnKarten); i++ {
		switch returnKarten[i].Fach {
		case "0":
			fach0++
		case "1":
			fach1++
		case "2":
			fach2++
		case "3":
			fach3++
		case "4":
			fach4++

		}
	}

	// Calucalate value for random karte
	var retKarte Karteikarte
	if len(returnKarten) != 0 {
		retKarte = returnKarten[rand.Intn(len(returnKarten))]
	}

	lernData := LernData{
		Kategorie:       decodeKasten.Kategorie,
		Titel:           decodeKasten.Titel,
		Beschreibung:    decodeKasten.Beschreibung,
		Fortschritt:     getFortschritt(returnKarten),
		Private:         decodeKasten.Private,
		CreatedByUserID: decodeKasten.CreatedByUserID,
		UserID:          decodeKasten.UserID,
		Ueberkategorie:  decodeKasten.Ueberkategorie,
		AnzKarten:       strconv.Itoa(anzKarten),
		AnzFachZero:     strconv.Itoa(fach0),
		AnzFachOne:      strconv.Itoa(fach1),
		AnzFachTwo:      strconv.Itoa(fach2),
		AnzFachThree:    strconv.Itoa(fach3),
		AnzFachFour:     strconv.Itoa(fach4),
		Karte:           retKarte,
	}
	return lernData, nil
}

// GetLern2Data ...
func GetLern2Data(kastenid string, karteid string) (LernData, error) {
	kasten, err := btDB.Get(kastenid, nil)
	if err != nil {
		return LernData{}, err
	}
	var decodeKasten Karteikasten
	mapstructure.Decode(kasten, &decodeKasten)
	decodeKasten.Id = kasten["_id"].(string)
	karte, err := btDB.Get(karteid, nil)
	if err != nil {
		return LernData{}, err
	}
	var decodeKarte Karteikarte
	mapstructure.Decode(karte, &decodeKarte)
	decodeKarte.Id = karte["_id"].(string)

	allKarten, err := GetAllKarten()
	if err != nil {
		return LernData{}, err
	}
	var decodeKarten []Karteikarte
	var tempKarten []Karteikarte
	anzKarten := 0
	mapstructure.Decode(allKarten, &decodeKarten)
	// Filter allKarten and only get the karten that belong to the kasten
	for i := 0; i < len(decodeKarten); i++ {
		if decodeKarten[i].KastenID == kastenid {
			tempKarten = append(tempKarten, decodeKarten[i])
			anzKarten++
		}
	}
	// Calculate count for each fach
	var fach0, fach1, fach2, fach3, fach4 int
	fach1 = 0
	for i := 0; i < len(decodeKarten); i++ {
		switch decodeKarten[i].Fach {
		case "0":
			fach0++
		case "1":
			fach1++
		case "2":
			fach2++
		case "3":
			fach3++
		case "4":
			fach4++
		}
	}

	lernData := LernData{
		Kategorie:       decodeKasten.Kategorie,
		Titel:           decodeKasten.Titel,
		Beschreibung:    decodeKasten.Beschreibung,
		Fortschritt:     getFortschritt(tempKarten),
		Private:         decodeKasten.Private,
		CreatedByUserID: decodeKasten.CreatedByUserID,
		UserID:          decodeKasten.UserID,
		Ueberkategorie:  decodeKasten.Ueberkategorie,
		AnzKarten:       strconv.Itoa(anzKarten),
		AnzFachZero:     strconv.Itoa(fach0),
		AnzFachOne:      strconv.Itoa(fach1),
		AnzFachTwo:      strconv.Itoa(fach2),
		AnzFachThree:    strconv.Itoa(fach3),
		AnzFachFour:     strconv.Itoa(fach4),
		Karte:           decodeKarte,
	}
	return lernData, nil
}

// GetEditData ...
func GetEditData() (EditData, error) {
	editData := EditData{}
	return editData, nil
}

// ---------------------------------------------------------------------------
// Internal helper functions
// ---------------------------------------------------------------------------

func getFortschritt(karten []Karteikarte) string {
	// Loop through all karten and count how many karten are in each fach
	fach0 := 0
	fach1 := 0
	fach2 := 0
	fach3 := 0
	fach4 := 0
	for _, s := range karten {
		switch s.Fach {
		case "0":
			fach0++
			break
		case "1":
			fach1++
			break
		case "2":
			fach2++
			break
		case "3":
			fach3++
			break
		case "4":
			fach4++
			break
		}
	}
	// Calculate progress for kasten
	temp := 0
	temp += 1 * fach1
	temp += 2 * fach2
	temp += 3 * fach3
	temp += 4 * fach4
	retWert := math.Round(float64((temp / (4 * len(karten))) * 100))

	return strconv.Itoa(int(retWert))
}
