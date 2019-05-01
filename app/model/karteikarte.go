package model

// Karteikarte Struct (Model)
type Karteikarte struct {
	KarteID  int64  `json:"karteid,string"`
	KastenID int64  `json:"kastenid,string"`
	Titel    string `json:"titel"`
	Frage    string `json:"frage"`
	Antwort  string `json:"antwort"`
	Fach     int64  `json:"fach,string"`
}
