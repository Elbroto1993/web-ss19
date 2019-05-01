package model

// Karteikasten Struct (Model)
type Karteikasten struct {
	KastenID        int64  `json:"kastenid,string"`
	Kategorie       string `json:"kategorie"`
	Titel           string `json:"titel"`
	Beschreibung    string `json:"beschreibung"`
	Fortschritt     int64  `json:"fortschritt,string"`
	Private         bool   `json:"private,string"`
	CreatedByUserID int64  `json:"createdByUserId,string"`
	UserID          int64  `json:"userid,string"`
	Ueberkategorie  string `json:"ueberkategorie"`
}
