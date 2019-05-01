package repository

import (
	"database/sql"

	"github.com/Elbroto1993/web-ss19/app/controller/karteikarte"
	"github.com/Elbroto1993/web-ss19/app/model"
)

type mysqlKarteikarteRepository struct {
	Conn *sql.DB
}

// NewMysqlUserRepository will create an object that represent the user.Repository interface
func NewMysqlKarteikarteRepository(Conn *sql.DB) karteikarte.Repository {
	return &mysqlKarteikarteRepository{Conn}
}

func (m *mysqlKarteikarteRepository) Fetch() ([]model.Karteikarte, error) {
	var allKarten []model.Karteikarte
	// Create Prepared Statement
	stmt, err := m.Conn.Prepare("SELECT karte_id, kasten_id, titel, frage, antwort, fach FROM karteikarte")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	// Query the stmt and save results in rows
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// Loop through results and save in book/author
	for rows.Next() {
		var karte model.Karteikarte
		err := rows.Scan(&karte.KarteID, &karte.KastenID, &karte.Titel, &karte.Frage, &karte.Antwort, &karte.Fach)
		if err != nil {
			return nil, err
		}
		allKarten = append(allKarten, karte)
	}
	return allKarten, nil
}

func (m *mysqlKarteikarteRepository) GetByID(id int64) (*model.Karteikarte, error) {
	karte := &model.Karteikarte{}
	// Create Prepared Statement
	stmt, err := m.Conn.Prepare("SELECT karte_id, kasten_id, titel, frage, antwort, fach FROM karteikarte WHERE karte_id=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	// Query the stmt and save results in rows
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// Loop through results and save in user
	for rows.Next() {
		err := rows.Scan(&karte.KarteID, &karte.KastenID, &karte.Titel, &karte.Frage, &karte.Antwort, &karte.Fach)
		if err != nil {
			return nil, err
		}
	}
	return karte, nil
}

func (m *mysqlKarteikarteRepository) GetByKastenID(id int64) ([]model.Karteikarte, error) {
	var allKarten []model.Karteikarte
	// Create Prepared Statement
	stmt, err := m.Conn.Prepare("SELECT karte_id, kasten_id, titel, frage, antwort, fach FROM karteikarte WHERE kasten_id=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	// Query the stmt and save results in rows
	rows, err := stmt.Query(id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// Loop through results and save in user
	for rows.Next() {
		var karte model.Karteikarte
		err := rows.Scan(&karte.KarteID, &karte.KastenID, &karte.Titel, &karte.Frage, &karte.Antwort, &karte.Fach)
		if err != nil {
			return nil, err
		}
		allKarten = append(allKarten, karte)
	}
	return allKarten, nil
}

func (m *mysqlKarteikarteRepository) Store(k *model.Karteikarte) error {
	// PreparedStatement for insert user
	stmt, err := m.Conn.Prepare("INSERT karteikarte SET kasten_id=?, titel=?, frage=?, antwort=?, fach=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Fill prepStmt with values and execute query
	res, err := stmt.Exec(k.KastenID, k.Titel, k.Frage, k.Antwort, k.Fach)
	if err != nil {
		return err
	}
	// Get last inserted id
	lastID, err := res.LastInsertId()
	k.KarteID = lastID
	return nil
}

func (m *mysqlKarteikarteRepository) Delete(id int64) error {
	stmt, err := m.Conn.Prepare("DELETE FROM karteikarte WHERE karte_id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}
	return nil
}

func (m *mysqlKarteikarteRepository) Update(k *model.Karteikarte) error {
	// PreparedStatement for Update
	stmt, err := m.Conn.Prepare("UPDATE karteikarte set titel=?, frage=?, antwort=?, fach=? WHERE karte_id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	// Fill prepStmt with values and execute query
	_, err = stmt.Exec(k.Titel, k.Frage, k.Antwort, k.Fach, k.KarteID)
	if err != nil {
		return err
	}
	return nil
}
