package model

import (
	couchdb "github.com/leesper/couchdb-golang"
	"github.com/mitchellh/mapstructure"
	"math"
	"math/rand"
	"strconv"
	"time"
)

// IndexData Struct
type IndexData struct {
	AnzUser                string `json:"anzuser"`
	AnzKasten              string `json:"anzkasten"`
	AnzKarten              string `json:"anzkarten"`
	LoggedIn               string `json:"loggedin"`
	UserName               string `json:"username"`
	AnzEigeneKaesten       string `json:"anzeigenekasten"`
	AnzOeffentlicheKaesten string `json:"anzoeffentlichekaesten"`
	ErrorMsg               string `json:"errormsg"`
	Image                  string `json:"image"`
}

// KarteikastenData Struct
type KarteikastenData struct {
	Id                     string         `json:"id"`
	LoggedIn               string         `json:"loggedin"`
	UserName               string         `json:"username"`
	AnzEigeneKaesten       string         `json:"anzeigenekasten"`
	AnzOeffentlicheKaesten string         `json:"anzoeffentlichekaesten"`
	Kaesten                []Karteikasten `json:"kaesten"`
	Image                  string         `json:"image"`
}

// MeineKarteienData Struct
type MeineKarteienData struct {
	Id                     string         `json:"id"`
	UserName               string         `json:"username"`
	AnzEigeneKaesten       string         `json:"anzeigenekasten"`
	AnzOeffentlicheKaesten string         `json:"anzoeffentlichekaesten"`
	Fortschritt            string         `json:"fortschritt"`
	MeineKaesten           []Karteikasten `json:"meinekaesten"`
	AndereKaesten          []Karteikasten `json:"anderekaesten"`
	Image                  string         `json:"image"`
}

// ViewData Struct
type ViewData struct {
	Kategorie              string        `json:"kategorie"`
	Titel                  string        `json:"titel"`
	Beschreibung           string        `json:"beschreibung"`
	Fortschritt            string        `json:"fortschritt"`
	Private                string        `json:"private"`
	CreatedByUserID        string        `json:"createdByUserId"`
	CreatedByUsername      string        `json:"createdbyusername"`
	UserID                 string        `json:"userid"`
	Ueberkategorie         string        `json:"ueberkategorie"`
	AnzKarten              string        `json:"anzkarten"`
	UserName               string        `json:"username"`
	LoggedIn               string        `json:"loggedin"`
	AnzEigeneKaesten       string        `json:"anzeigenekasten"`
	AnzOeffentlicheKaesten string        `json:"anzoeffentlichekaesten"`
	SelectedKarte          Karteikarte   `json:"selectedkarte"`
	Karten                 []Karteikarte `json:"karten"`
	Image                  string        `json:"image"`
}

// LernData Struct
type LernData struct {
	Kategorie              string      `json:"kategorie"`
	Titel                  string      `json:"titel"`
	Beschreibung           string      `json:"beschreibung"`
	Fortschritt            string      `json:"fortschritt"`
	Private                string      `json:"private"`
	CreatedByUserID        string      `json:"createdByUserId"`
	UserID                 string      `json:"userid"`
	Ueberkategorie         string      `json:"ueberkategorie"`
	AnzKarten              string      `json:"anzkarten"`
	AnzFachZero            string      `json:"anzfachzero"`
	AnzFachOne             string      `json:"anzfachone"`
	AnzFachTwo             string      `json:"anzfachtwo"`
	AnzFachThree           string      `json:"anzfachthree"`
	AnzFachFour            string      `json:"anzfachfour"`
	UserName               string      `json:"username"`
	AnzEigeneKaesten       string      `json:"anzeigenekasten"`
	AnzOeffentlicheKaesten string      `json:"anzoeffentlichekaesten"`
	Karte                  Karteikarte `json:"karte"`
	NewKastenID            string      `json:"kastenid"`
	Image                  string      `json:"image"`
}

// EditData Struct
type EditData struct {
	UserName               string `json:"username"`
	AnzEigeneKaesten       string `json:"anzeigenekasten"`
	AnzOeffentlicheKaesten string `json:"anzoeffentlichekaesten"`
	Image                  string `json:"image"`
}

// Edit2Data Struct
type Edit2Data struct {
	Id                     string        `json:"id"`
	Kategorie              string        `json:"kategorie"`
	Titel                  string        `json:"titel"`
	Fortschritt            string        `json:"fortschritt"`
	CreatedByUserID        string        `json:"createdByUserId"`
	UserID                 string        `json:"userid"`
	Ueberkategorie         string        `json:"ueberkategorie"`
	AnzKarten              string        `json:"anzkarten"`
	UserName               string        `json:"username"`
	AnzEigeneKaesten       string        `json:"anzeigenekasten"`
	AnzOeffentlicheKaesten string        `json:"anzoeffentlichekaesten"`
	SelectedKarte          Karteikarte   `json:"selectedkarte"`
	Karten                 []Karteikarte `json:"karten"`
	Image                  string        `json:"image"`
}

// ProfilData Struct
type ProfilData struct {
	UserName               string    `json:"username"`
	Email                  string    `json:"email"`
	AnzEigeneKaesten       string    `json:"anzeigenekasten"`
	AnzOeffentlicheKaesten string    `json:"anzoeffentlichekaesten"`
	AnzEigeneKarten        string    `json:"anzeigenekarten"`
	CreatedAt              time.Time `json:"createdat"`
	Image                  string    `json:"image"`
}

// RegisterData struct
type RegisterData struct {
	AnzOeffentlicheKaesten string `json:"anzoeffentlichekaesten"`
	ErrorMsg               string `json:"errormsg"`
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
func GetIndexData(username string) (IndexData, error) {
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
	anzOeffentlicheKaesten, err := GetAlleOeffentlichenKaesten()
	if err != nil {
		return IndexData{}, err
	}
	var eigeneKaesten []Karteikasten
	if username != "" {
		eigeneKaesten, err = getEigeneKaesten(username)
		if err != nil {
			return IndexData{}, err
		}
	}
	var user User
	if username != "" {
		user, err = GetUserByUsername(username)
		if err != nil {
			return IndexData{}, err
		}
	}
	index := IndexData{
		AnzUser:                strconv.Itoa(len(allUser)),
		AnzKasten:              strconv.Itoa(len(allKasten)),
		AnzKarten:              strconv.Itoa(len(allKarten)),
		AnzOeffentlicheKaesten: strconv.Itoa(len(anzOeffentlicheKaesten)),
		AnzEigeneKaesten:       strconv.Itoa(len(eigeneKaesten)),
		Image:                  user.Image,
	}
	return index, nil
}

// GetRegisterData ...
func GetRegisterData() (RegisterData, error) {
	anzOeffentlicheKaesten, err := GetAlleOeffentlichenKaesten()
	if err != nil {
		return RegisterData{}, err
	}

	registerData := RegisterData{
		AnzOeffentlicheKaesten: strconv.Itoa(len(anzOeffentlicheKaesten)),
		ErrorMsg:               "",
	}

	return registerData, nil
}

// GetKarteikastenData ...
func GetKarteikastenData(username string, kategorie string) (KarteikastenData, error) {
	// Get all oeffentliche Kaesten and decode them
	alleOeffentlichenKaesten, err := GetAlleOeffentlichenKaesten()
	if err != nil {
		return KarteikastenData{}, err
	}

	var filterKaesten []Karteikasten
	if kategorie != "" {
		for i := 0; i < len(alleOeffentlichenKaesten); i++ {
			if alleOeffentlichenKaesten[i].Kategorie == kategorie {
				filterKaesten = append(filterKaesten, alleOeffentlichenKaesten[i])
			}
		}
		// Get all eigene kaesten if logged in
		var eigeneKaesten []Karteikasten
		if username != "" {
			eigeneKaesten, err = getEigeneKaesten(username)
			if err != nil {
				return KarteikastenData{}, err
			}
		}

		// Get all Karten and decode them
		allKarten, err := GetAllKarten()
		if err != nil {
			return KarteikastenData{}, err
		}
		var decodedKarten []Karteikarte
		mapstructure.Decode(allKarten, &decodedKarten)
		// Fill AnzKarten from Karteikasten with values
		for i := 0; i < len(filterKaesten); i++ {
			countKarten := 0
			for j := 0; j < len(decodedKarten); j++ {
				if string(filterKaesten[i].Id) == decodedKarten[j].KastenID {
					countKarten++
				}
			}
			filterKaesten[i].AnzKarten = strconv.Itoa(countKarten)
		}
		var user User
		if username != "" {
			user, err = GetUserByUsername(username)
			if err != nil {
				return KarteikastenData{}, err
			}
		}
		var retValue KarteikastenData
		retValue.Kaesten = filterKaesten
		retValue.AnzOeffentlicheKaesten = strconv.Itoa(len(alleOeffentlichenKaesten))
		retValue.AnzEigeneKaesten = strconv.Itoa(len(eigeneKaesten))
		retValue.Image = user.Image
		return retValue, nil
	} else {

		// Get all eigene kaesten if logged in
		var eigeneKaesten []Karteikasten
		if username != "" {
			eigeneKaesten, err = getEigeneKaesten(username)
			if err != nil {
				return KarteikastenData{}, err
			}
		}

		// Get all Karten and decode them
		allKarten, err := GetAllKarten()
		if err != nil {
			return KarteikastenData{}, err
		}
		var decodedKarten []Karteikarte
		mapstructure.Decode(allKarten, &decodedKarten)
		// Fill AnzKarten from Karteikasten with values
		for i := 0; i < len(alleOeffentlichenKaesten); i++ {
			countKarten := 0
			for j := 0; j < len(decodedKarten); j++ {
				if string(alleOeffentlichenKaesten[i].Id) == decodedKarten[j].KastenID {
					countKarten++
				}
			}
			alleOeffentlichenKaesten[i].AnzKarten = strconv.Itoa(countKarten)
		}
		var user User
		if username != "" {
			user, err = GetUserByUsername(username)
			if err != nil {
				return KarteikastenData{}, err
			}
		}
		var retValue KarteikastenData
		retValue.Kaesten = alleOeffentlichenKaesten
		retValue.AnzOeffentlicheKaesten = strconv.Itoa(len(alleOeffentlichenKaesten))
		retValue.AnzEigeneKaesten = strconv.Itoa(len(eigeneKaesten))
		retValue.Image = user.Image
		return retValue, nil
	}
}

// GetMeineKarteienData ...
func GetMeineKarteienData(username string) (MeineKarteienData, error) {
	user, err := GetUserByUsername(username)
	if err != nil {
		return MeineKarteienData{}, err
	}

	// Get all oeffentliche Kaesten and decode them
	alleOeffentlichenKaesten, err := GetAlleOeffentlichenKaesten()
	if err != nil {
		return MeineKarteienData{}, err
	}

	allKaestenFromUser, err := getEigeneKaesten(username)
	if err != nil {
		return MeineKarteienData{}, err
	}

	// Get all Karten and decode them
	allKarten, err := GetAllKarten()
	if err != nil {
		return MeineKarteienData{}, err
	}
	var decodedKarten []Karteikarte
	mapstructure.Decode(allKarten, &decodedKarten)
	// Fill AnzKarten from Karteikasten with values
	for i := 0; i < len(allKaestenFromUser); i++ {
		var tempKarten []Karteikarte
		countKarten := 0
		for j := 0; j < len(decodedKarten); j++ {
			if string(allKaestenFromUser[i].Id) == decodedKarten[j].KastenID {
				tempKarten = append(tempKarten, decodedKarten[j])
				countKarten++
			}
		}
		allKaestenFromUser[i].Fortschritt = getFortschritt(tempKarten)
		allKaestenFromUser[i].AnzKarten = strconv.Itoa(countKarten)
	}

	var meineKaesten []Karteikasten
	var andereKaesten []Karteikasten
	// Split allKaestenFromUser in eigenen und andere kaesten
	for i := 0; i < len(allKaestenFromUser); i++ {
		if allKaestenFromUser[i].CreatedByUserID == user.Id {
			meineKaesten = append(meineKaesten, allKaestenFromUser[i])
		} else {
			andereKaesten = append(andereKaesten, allKaestenFromUser[i])
		}
	}
	var retValue MeineKarteienData
	retValue.MeineKaesten = meineKaesten
	retValue.AndereKaesten = andereKaesten
	retValue.AnzOeffentlicheKaesten = strconv.Itoa(len(alleOeffentlichenKaesten))
	retValue.AnzEigeneKaesten = strconv.Itoa(len(allKaestenFromUser))
	retValue.Image = user.Image
	return retValue, nil
}

// GetViewData ...
func GetViewData(kastenid string, karteid string, username string) (ViewData, error) {
	// Data for sidebar
	anzOeffentlicheKaesten, err := GetAlleOeffentlichenKaesten()
	if err != nil {
		return ViewData{}, err
	}
	var eigeneKaesten []Karteikasten
	if username != "" {
		eigeneKaesten, err = getEigeneKaesten(username)
		if err != nil {
			return ViewData{}, err
		}
	}

	kasten, err := btDB.Get(kastenid, nil)
	if err != nil {
		return ViewData{}, err
	}
	var decodeKasten Karteikasten
	mapstructure.Decode(kasten, &decodeKasten)

	// Get username from createdByUserId
	createdByUsername, err := btDB.Get(decodeKasten.UserID, nil)
	if err != nil {
		return ViewData{}, err
	}

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

	// Check if kasten is being owned by current user for fortschritt
	var fortschritt string
	fortschritt = strconv.Itoa(0)
	if username != "" {
		currentUser, err := GetUserByUsername(username)
		if err != nil {
			return ViewData{}, err
		}
		if currentUser.Id == decodeKasten.UserID {
			fortschritt = getFortschritt(returnKarten)
		}
	}
	var user User
	if username != "" {
		user, err = GetUserByUsername(username)
		if err != nil {
			return ViewData{}, err
		}
	}
	viewData := ViewData{
		Kategorie:              decodeKasten.Kategorie,
		Titel:                  decodeKasten.Titel,
		Beschreibung:           decodeKasten.Beschreibung,
		Fortschritt:            fortschritt,
		Private:                decodeKasten.Private,
		CreatedByUsername:      createdByUsername["username"].(string),
		CreatedByUserID:        decodeKasten.CreatedByUserID,
		UserID:                 decodeKasten.UserID,
		Ueberkategorie:         decodeKasten.Ueberkategorie,
		AnzKarten:              strconv.Itoa(anzKarten),
		AnzOeffentlicheKaesten: strconv.Itoa(len(anzOeffentlicheKaesten)),
		AnzEigeneKaesten:       strconv.Itoa(len(eigeneKaesten)),
		SelectedKarte:          decodeKarte,
		Karten:                 returnKarten,
		Image:                  user.Image,
	}
	return viewData, nil
}

// GetLernData ...
func GetLernData(_kastenid string, _karteid string, username string) (LernData, error) {
	var fach0, fach1, fach2, fach3, fach4 int
	var retKarte Karteikarte
	// Data for sidebar
	anzOeffentlicheKaesten, err := GetAlleOeffentlichenKaesten()
	if err != nil {
		return LernData{}, err
	}
	var eigeneKaesten []Karteikasten
	if username != "" {
		eigeneKaesten, err = getEigeneKaesten(username)
		if err != nil {
			return LernData{}, err
		}
	}

	kasten, err := btDB.Get(_kastenid, nil)
	if err != nil {
		return LernData{}, err
	}
	var decodeKasten Karteikasten
	mapstructure.Decode(kasten, &decodeKasten)
	decodeKasten.Id = kasten["_id"].(string)

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

	var newKastenId string
	newKastenId = ""
	////////// CHECK IF KASTEN IS ALREADY BEING LEARNED, ELSE ADD KASTEN TO USER ////////////
	user, _ := GetUserByUsername(username)
	if user.Id != decodeKasten.UserID {
		kasten, _ := GetKastenById(_kastenid)
		kasten.Private = "true"
		kasten.UserID = user.Id
		newKastenId, _ = kasten.Add()
		for i := 0; i < len(decodeKarten); i++ {
			if decodeKarten[i].KastenID == _kastenid {
				decodeKarten[i].KastenID = newKastenId
				decodeKarten[i].Add()
			}
		}
		// Filter allKarten and only get the karten that belong to the kasten
		for i := 0; i < len(decodeKarten); i++ {
			if decodeKarten[i].KastenID == newKastenId {
				returnKarten = append(returnKarten, decodeKarten[i])
				anzKarten++
			}
		}
		// Calculate count for each fach
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
		karteGefunden := false
		if len(returnKarten)-2 > 0 {
			for !karteGefunden {
				zufallsFach := getZufallsFach()
				for i := 0; i < len(returnKarten); i++ {
					if returnKarten[i].Fach == strconv.Itoa(zufallsFach) && returnKarten[i].Id != _karteid {
						retKarte = returnKarten[i]
						karteGefunden = true
					}
				}
			}
		} else if len(returnKarten) > 0 {
			retKarte = returnKarten[0]
		}
	} else {
		// Filter allKarten and only get the karten that belong to the kasten
		for i := 0; i < len(decodeKarten); i++ {
			if decodeKarten[i].KastenID == _kastenid {
				returnKarten = append(returnKarten, decodeKarten[i])
				anzKarten++
			}
		}
		// Calculate count for each fach
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
		karteGefunden := false
		if len(returnKarten)-2 > 0 {
			for !karteGefunden {
				zufallsFach := getZufallsFach()
				for i := 0; i < len(returnKarten); i++ {
					if returnKarten[i].Fach == strconv.Itoa(zufallsFach) && returnKarten[i].Id != _karteid {
						retKarte = returnKarten[i]
						karteGefunden = true
					}
				}
			}
		} else if len(returnKarten) > 0 {
			retKarte = returnKarten[0]
		}
	}
	////////////////////////////////////////////////////////////////////////////////////////
	lernData := LernData{
		Kategorie:              decodeKasten.Kategorie,
		Titel:                  decodeKasten.Titel,
		Beschreibung:           decodeKasten.Beschreibung,
		Fortschritt:            getFortschritt(returnKarten),
		Private:                decodeKasten.Private,
		CreatedByUserID:        decodeKasten.CreatedByUserID,
		UserID:                 decodeKasten.UserID,
		Ueberkategorie:         decodeKasten.Ueberkategorie,
		AnzKarten:              strconv.Itoa(anzKarten),
		AnzFachZero:            strconv.Itoa(fach0),
		AnzFachOne:             strconv.Itoa(fach1),
		AnzFachTwo:             strconv.Itoa(fach2),
		AnzFachThree:           strconv.Itoa(fach3),
		AnzFachFour:            strconv.Itoa(fach4),
		AnzOeffentlicheKaesten: strconv.Itoa(len(anzOeffentlicheKaesten)),
		AnzEigeneKaesten:       strconv.Itoa(len(eigeneKaesten)),
		Karte:                  retKarte,
		NewKastenID:            newKastenId,
		Image:                  user.Image,
	}
	return lernData, nil
}

// GetLern2Data ...
func GetLern2Data(kastenid string, karteid string, username string) (LernData, error) {
	// Data for sidebar
	anzOeffentlicheKaesten, err := GetAlleOeffentlichenKaesten()
	if err != nil {
		return LernData{}, err
	}
	var eigeneKaesten []Karteikasten
	if username != "" {
		eigeneKaesten, err = getEigeneKaesten(username)
		if err != nil {
			return LernData{}, err
		}
	}

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
	for i := 0; i < len(tempKarten); i++ {
		switch tempKarten[i].Fach {
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
	var user User
	if username != "" {
		user, _ = GetUserByUsername(username)
	}
	lernData := LernData{
		Kategorie:              decodeKasten.Kategorie,
		Titel:                  decodeKasten.Titel,
		Beschreibung:           decodeKasten.Beschreibung,
		Fortschritt:            getFortschritt(tempKarten),
		Private:                decodeKasten.Private,
		CreatedByUserID:        decodeKasten.CreatedByUserID,
		UserID:                 decodeKasten.UserID,
		Ueberkategorie:         decodeKasten.Ueberkategorie,
		AnzKarten:              strconv.Itoa(anzKarten),
		AnzFachZero:            strconv.Itoa(fach0),
		AnzFachOne:             strconv.Itoa(fach1),
		AnzFachTwo:             strconv.Itoa(fach2),
		AnzFachThree:           strconv.Itoa(fach3),
		AnzFachFour:            strconv.Itoa(fach4),
		AnzOeffentlicheKaesten: strconv.Itoa(len(anzOeffentlicheKaesten)),
		AnzEigeneKaesten:       strconv.Itoa(len(eigeneKaesten)),
		Karte:                  decodeKarte,
		NewKastenID:            decodeKarte.KastenID,
		Image:                  user.Image,
	}
	return lernData, nil
}

// GetEditData ...
func GetEditData(username string) (EditData, error) {
	editData := EditData{}
	user, err := GetUserByUsername(username)
	if err != nil {
		return EditData{}, err
	}
	editData.Image = user.Image
	return editData, nil
}

// GetEdit2Data ...
func GetEdit2Data(kastenid string, karteid string, username string) (Edit2Data, error) {
	// Data for sidebar
	anzOeffentlicheKaesten, err := GetAlleOeffentlichenKaesten()
	if err != nil {
		return Edit2Data{}, err
	}
	var eigeneKaesten []Karteikasten
	if username != "" {
		eigeneKaesten, err = getEigeneKaesten(username)
		if err != nil {
			return Edit2Data{}, err
		}
	}

	kasten, err := btDB.Get(kastenid, nil)
	if err != nil {
		return Edit2Data{}, err
	}
	var decodeKasten Karteikasten
	mapstructure.Decode(kasten, &decodeKasten)
	decodeKasten.Id = kasten["_id"].(string)

	// Only if a karte was selected
	var decodeKarte Karteikarte
	if karteid != "" {
		karte, err := btDB.Get(karteid, nil)
		if err != nil {
			return Edit2Data{}, err
		}
		mapstructure.Decode(karte, &decodeKarte)
		decodeKarte.Id = karte["_id"].(string)
	} else {
		decodeKarte = Karteikarte{}
	}

	allKarten, err := GetAllKarten()
	if err != nil {
		return Edit2Data{}, err
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
	var user User
	if username != "" {
		user, _ = GetUserByUsername(username)
	}
	edit2Data := Edit2Data{
		Id:                     decodeKasten.Id,
		UserName:               username,
		Kategorie:              decodeKasten.Kategorie,
		Titel:                  decodeKasten.Titel,
		Fortschritt:            getFortschritt(returnKarten),
		CreatedByUserID:        decodeKasten.CreatedByUserID,
		UserID:                 decodeKasten.UserID,
		Ueberkategorie:         decodeKasten.Ueberkategorie,
		AnzKarten:              strconv.Itoa(anzKarten),
		AnzOeffentlicheKaesten: strconv.Itoa(len(anzOeffentlicheKaesten)),
		AnzEigeneKaesten:       strconv.Itoa(len(eigeneKaesten)),
		SelectedKarte:          decodeKarte,
		Karten:                 returnKarten,
		Image:                  user.Image,
	}
	return edit2Data, nil
}

func GetProfilData(username string) (ProfilData, error) {
	user, err := GetUserByUsername(username)
	if err != nil {
		return ProfilData{}, err
	}

	anzEigeneKaesten, err := getEigeneKaesten(username)
	if err != nil {
		return ProfilData{}, err
	}

	eigeneKarten, err := GetEigeneKarten(username)
	if err != nil {
		return ProfilData{}, err
	}

	anzOeffentlicheKaesten, err := GetAlleOeffentlichenKaesten()
	if err != nil {
		return ProfilData{}, err
	}

	profilData := ProfilData{
		UserName:               user.Username,
		Email:                  user.Email,
		AnzEigeneKaesten:       strconv.Itoa(len(anzEigeneKaesten)),
		AnzEigeneKarten:        strconv.Itoa(len(eigeneKarten)),
		AnzOeffentlicheKaesten: strconv.Itoa(len(anzOeffentlicheKaesten)),
		CreatedAt:              user.CreatedAt,
		Image:                  user.Image,
	}

	return profilData, nil
}

// ---------------------------------------------------------------------------
// Internal helper functions
// ---------------------------------------------------------------------------

func getFortschritt(karten []Karteikarte) string {
	if len(karten) == 0 {
		return strconv.Itoa(0)
	}
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
	var retWert float64
	retWert = math.Floor(float64((float64(temp) / (float64(4) * float64(len(karten))) * float64(100))))
	return strconv.Itoa(int(retWert))
}

func getEigeneKaesten(username string) ([]Karteikasten, error) {
	userInDB, err := GetUserByUsername(username)
	if err != nil {
		return []Karteikasten{}, err
	}

	kaesten, err := GetAllKasten()
	if err != nil {
		return []Karteikasten{}, err
	}

	var decodedKaesten []Karteikasten
	mapstructure.Decode(kaesten, &decodedKaesten)
	// Add _id to decodeKarten, because mapstructure.Decode doesn't do it
	index := 0
	for _, v := range kaesten {
		decodedKaesten[index].Id = v["_id"].(string)
		index++
	}

	var retKaesten []Karteikasten
	mapstructure.Decode(kaesten, &decodedKaesten)
	for i := 0; i < len(kaesten); i++ {
		if decodedKaesten[i].UserID == userInDB.Id {
			retKaesten = append(retKaesten, decodedKaesten[i])
		}
	}
	return retKaesten, nil
}

func GetEigeneKarten(username string) ([]Karteikarte, error) {
	userInDB, err := GetUserByUsername(username)
	if err != nil {
		return []Karteikarte{}, err
	}

	kaesten, err := GetAllKasten()
	if err != nil {
		return []Karteikarte{}, err
	}
	var decodedKaesten []Karteikasten
	mapstructure.Decode(kaesten, &decodedKaesten)
	// Add _id to decodedKaesten, because mapstructure.Decode doesn't do it
	index := 0
	for _, v := range kaesten {
		decodedKaesten[index].Id = v["_id"].(string)
		index++
	}

	var kaestenFromUser []Karteikasten
	for i := 0; i < len(kaesten); i++ {
		if decodedKaesten[i].UserID == userInDB.Id {
			kaestenFromUser = append(kaestenFromUser, decodedKaesten[i])
		}
	}

	karten, err := GetAllKarten()
	if err != nil {
		return []Karteikarte{}, err
	}

	var decodedKarten []Karteikarte
	mapstructure.Decode(karten, &decodedKarten)
	// Add _id to decodedKarten, because mapstructure.Decode doesn't do it
	index = 0
	for _, v := range karten {
		decodedKarten[index].Id = v["_id"].(string)
		index++
	}

	var retKarten []Karteikarte
	for i := 0; i < len(decodedKarten); i++ {
		for j := 0; j < len(kaestenFromUser); j++ {
			if kaestenFromUser[j].Id == decodedKarten[i].KastenID {
				retKarten = append(retKarten, decodedKarten[i])
			}
		}
	}
	return retKarten, nil
}

func GetAlleOeffentlichenKaesten() ([]Karteikasten, error) {

	kaesten, err := GetAllKasten()
	if err != nil {
		return []Karteikasten{}, err
	}

	var decodedKaesten []Karteikasten
	mapstructure.Decode(kaesten, &decodedKaesten)
	// Add _id to decodeKarten, because mapstructure.Decode doesn't do it
	index := 0
	for _, v := range kaesten {
		decodedKaesten[index].Id = v["_id"].(string)
		index++
	}

	var retKaesten []Karteikasten
	for i := 0; i < len(kaesten); i++ {
		if decodedKaesten[i].Private == "false" {
			retKaesten = append(retKaesten, decodedKaesten[i])
		}
	}

	return retKaesten, nil
}

// Function to decide which fach will be shown to the user
func getZufallsFach() int {
	r := math.Floor((float64(rand.Intn(14-0) + 0)))
	var f int

	switch r {
	case 0:
		f = 4
		break
	case 1:
	case 2:
		f = 3
		break
	case 3:
	case 4:
	case 5:
		f = 2
		break
	case 6:
	case 7:
	case 8:
	case 9:
		f = 1
		break
	case 10:
	case 11:
	case 12:
	case 13:
	case 14:
		f = 0
		break
	}
	return f
}
