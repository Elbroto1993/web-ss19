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
}

// KarteikastenData Struct
type KarteikastenData struct {
	Id                     string         `json:"id"`
	LoggedIn               string         `json:"loggedin"`
	UserName               string         `json:"username"`
	AnzEigeneKaesten       string         `json:"anzeigenekasten"`
	AnzOeffentlicheKaesten string         `json:"anzoeffentlichekaesten"`
	Kaesten                []Karteikasten `json:"kaesten"`
}

// MeineKarteienData Struct
type MeineKarteienData struct {
	Id                     string         `json:"id"`
	UserName               string         `json:"username"`
	AnzEigeneKaesten       string         `json:"anzeigenekasten"`
	AnzOeffentlicheKaesten string         `json:"anzoeffentlichekaesten"`
	Fortschritt            string         `json:"fortschritt"`
	Kaesten                []Karteikasten `json:"kaesten"`
}

// ViewData Struct
type ViewData struct {
	Kategorie              string        `json:"kategorie"`
	Titel                  string        `json:"titel"`
	Beschreibung           string        `json:"beschreibung"`
	Fortschritt            string        `json:"fortschritt"`
	Private                string        `json:"private"`
	CreatedByUserID        string        `json:"createdByUserId"`
	UserID                 string        `json:"userid"`
	Ueberkategorie         string        `json:"ueberkategorie"`
	AnzKarten              string        `json:"anzkarten"`
	UserName               string        `json:"username"`
	LoggedIn               string        `json:"loggedin"`
	AnzEigeneKaesten       string        `json:"anzeigenekasten"`
	AnzOeffentlicheKaesten string        `json:"anzoeffentlichekaesten"`
	SelectedKarte          Karteikarte   `json:"selectedkarte"`
	Karten                 []Karteikarte `json:"karten"`
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
}

// EditData Struct
type EditData struct {
	UserName               string `json:"username"`
	AnzEigeneKaesten       string `json:"anzeigenekasten"`
	AnzOeffentlicheKaesten string `json:"anzoeffentlichekaesten"`
}

// Edit2Data Struct
type Edit2Data struct {
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
}

// ProfilData Struct
type ProfilData struct {
	UserName               string    `json:"username"`
	Email                  string    `json:"email"`
	AnzEigeneKaesten       string    `json:"anzeigenekasten"`
	AnzOeffentlicheKaesten string    `json:"anzoeffentlichekaesten"`
	AnzEigeneKarten        string    `json:"anzeigenekarten"`
	CreatedAt              time.Time `json:"createdat"`
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
	index := IndexData{
		AnzUser:                strconv.Itoa(len(allUser)),
		AnzKasten:              strconv.Itoa(len(allKasten)),
		AnzKarten:              strconv.Itoa(len(allKarten)),
		AnzOeffentlicheKaesten: strconv.Itoa(len(anzOeffentlicheKaesten)),
		AnzEigeneKaesten:       strconv.Itoa(len(eigeneKaesten)),
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
func GetKarteikastenData(username string) (KarteikastenData, error) {
	// Get all oeffentliche Kaesten and decode them
	alleOeffentlichenKaesten, err := GetAlleOeffentlichenKaesten()
	if err != nil {
		return KarteikastenData{}, err
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
	for i := 0; i < len(alleOeffentlichenKaesten); i++ {
		countKarten := 0
		for j := 0; j < len(decodedKarten); j++ {
			if string(alleOeffentlichenKaesten[i].Id) == decodedKarten[j].KastenID {
				countKarten++
			}
		}
		alleOeffentlichenKaesten[i].AnzKarten = strconv.Itoa(countKarten)
	}
	var retValue KarteikastenData
	retValue.Kaesten = alleOeffentlichenKaesten
	retValue.AnzOeffentlicheKaesten = strconv.Itoa(len(alleOeffentlichenKaesten))
	retValue.AnzEigeneKaesten = strconv.Itoa(len(eigeneKaesten))
	return retValue, nil
}

// GetMeineKarteienData ...
func GetMeineKarteienData(username string) (MeineKarteienData, error) {
	// Get all oeffentliche Kaesten and decode them
	alleOeffentlichenKaesten, err := GetAlleOeffentlichenKaesten()
	if err != nil {
		return MeineKarteienData{}, err
	}

	eigeneKaesten, err := getEigeneKaesten(username)
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
	for i := 0; i < len(eigeneKaesten); i++ {
		var tempKarten []Karteikarte
		countKarten := 0
		for j := 0; j < len(decodedKarten); j++ {
			if string(eigeneKaesten[i].Id) == decodedKarten[j].KastenID {
				tempKarten = append(tempKarten, decodedKarten[j])
				countKarten++
			}
		}
		eigeneKaesten[i].Fortschritt = getFortschritt(tempKarten)
		eigeneKaesten[i].AnzKarten = strconv.Itoa(countKarten)
	}
	var retValue MeineKarteienData
	retValue.Kaesten = eigeneKaesten
	retValue.AnzOeffentlicheKaesten = strconv.Itoa(len(alleOeffentlichenKaesten))
	retValue.AnzEigeneKaesten = strconv.Itoa(len(eigeneKaesten))
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
		Kategorie:              decodeKasten.Kategorie,
		Titel:                  decodeKasten.Titel,
		Beschreibung:           decodeKasten.Beschreibung,
		Fortschritt:            getFortschritt(returnKarten),
		Private:                decodeKasten.Private,
		CreatedByUserID:        decodeKasten.CreatedByUserID,
		UserID:                 decodeKasten.UserID,
		Ueberkategorie:         decodeKasten.Ueberkategorie,
		AnzKarten:              strconv.Itoa(anzKarten),
		AnzOeffentlicheKaesten: strconv.Itoa(len(anzOeffentlicheKaesten)),
		AnzEigeneKaesten:       strconv.Itoa(len(eigeneKaesten)),
		SelectedKarte:          decodeKarte,
		Karten:                 returnKarten,
	}
	return viewData, nil
}

// GetLernData ...
func GetLernData(_id string, username string) (LernData, error) {
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

	////////// CHECK IF KASTEN IS ALREADY BEING LEARNED, ELSE ADD KASTEN TO USER ////////////
	user, _ := GetUserByUsername(username)
	if user.Id != _id {
		kasten, _ := GetKastenById(_id)
		kasten.UserID = user.Id
		newKastenId, _ := kasten.Add()
		for i := 0; i < len(decodeKarten); i++ {
			if decodeKarten[i].KastenID == _id {
				decodeKarten[i].KastenID = newKastenId
				decodeKarten[i].Add()
			}
		}
	}
	/////////////////////////////////////////////////////////////////////////////////////////

	// Calucalate value for random karte
	var retKarte Karteikarte
	if len(returnKarten) != 0 {
		retKarte = returnKarten[rand.Intn(len(returnKarten))]
	}

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
	}
	return lernData, nil
}

// GetEditData ...
func GetEditData() (EditData, error) {
	editData := EditData{}
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

	edit2Data := Edit2Data{
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
