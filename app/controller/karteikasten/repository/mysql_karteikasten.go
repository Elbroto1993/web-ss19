package repository

import (
	"database/sql"

	"github.com/Elbroto1993/web-ss19/app/controller/karteikasten"
	"github.com/Elbroto1993/web-ss19/app/model"
)

type mysqlKarteikastenRepository struct {
	Conn *sql.DB
}

// NewMysqlKarteikastenRepository will create an object that represent the karteikasten.Repository interface
func NewMysqlKarteikastenRepository(Conn *sql.DB) karteikasten.Repository {
	return &mysqlKarteikastenRepository{Conn}
}

func (m *mysqlKarteikastenRepository) Fetch() ([]model.Karteikasten, error) {
	var allKasten []model.Karteikasten
	// Create Prepared Statement
	stmt, err := m.Conn.Prepare("SELECT kasten_id, user_id, created_by_user_id, kategorie, titel, beschreibung, fortschritt, private, ueberkategorie FROM karteikasten")
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
		var kasten model.Karteikasten
		err := rows.Scan(&kasten.KastenID, &kasten.UserID, &kasten.CreatedByUserID, &kasten.Kategorie, &kasten.Titel, &kasten.Beschreibung, &kasten.Fortschritt, &kasten.Private, &kasten.Ueberkategorie)
		if err != nil {
			return nil, err
		}
		allKasten = append(allKasten, kasten)
	}
	return allKasten, nil
}

func (m *mysqlKarteikastenRepository) GetByID(id int64) (*model.Karteikasten, error) {
	kasten := &model.Karteikasten{}
	// Create Prepared Statement
	stmt, err := m.Conn.Prepare("SELECT kasten_id, user_id, created_by_user_id, kategorie, titel, beschreibung, fortschritt, private, ueberkategorie FROM karteikasten WHERE kasten_id=?")
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
	// Loop through results and save in karteikasten
	for rows.Next() {
		err := rows.Scan(&kasten.KastenID, &kasten.UserID, &kasten.CreatedByUserID, &kasten.Kategorie, &kasten.Titel, &kasten.Beschreibung, &kasten.Fortschritt, &kasten.Private, &kasten.Ueberkategorie)
		if err != nil {
			return nil, err
		}
	}
	return kasten, nil
}

func (m *mysqlKarteikastenRepository) GetByUserID(id int64) ([]model.Karteikasten, error) {
	var allKasten []model.Karteikasten
	// Create Prepared Statement
	stmt, err := m.Conn.Prepare("SELECT kasten_id, user_id, created_by_user_id, kategorie, titel, beschreibung, fortschritt, private, ueberkategorie FROM karteikasten WHERE user_id=?")
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
	// Loop through results and save in karteikasten
	for rows.Next() {
		var kasten model.Karteikasten
		err := rows.Scan(&kasten.KastenID, &kasten.UserID, &kasten.CreatedByUserID, &kasten.Kategorie, &kasten.Titel, &kasten.Beschreibung, &kasten.Fortschritt, &kasten.Private, &kasten.Ueberkategorie)
		if err != nil {
			return nil, err
		}
		allKasten = append(allKasten, kasten)
	}
	return allKasten, nil
}

func (m *mysqlKarteikastenRepository) Store(k *model.Karteikasten) error {
	stmt, err := m.Conn.Prepare("INSERT karteikasten SET user_id=?, created_by_user_id=?, kategorie=?, titel=?, beschreibung=?, fortschritt=?, private=?, ueberkategorie=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Fill prepStmt with values and execute query
	res, err := stmt.Exec(k.UserID, k.CreatedByUserID, k.Kategorie, k.Titel, k.Beschreibung, k.Fortschritt, k.Private, k.Ueberkategorie)
	if err != nil {
		return err
	}
	// Get last inserted id
	lastID, err := res.LastInsertId()
	k.KastenID = lastID
	return nil
}

func (m *mysqlKarteikastenRepository) Delete(id int64) error {
	stmt, err := m.Conn.Prepare("DELETE FROM karteikasten WHERE kasten_id=?")
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

func (m *mysqlKarteikastenRepository) Update(k *model.Karteikasten) error {
	stmt, err := m.Conn.Prepare("UPDATE karteikasten SET kategorie=?, titel=?, beschreibung=?, private=?, ueberkategorie=?, fortschritt=? WHERE kasten_id=?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	// Fill prepStmt with values and execute query
	_, err = stmt.Exec(k.Kategorie, k.Titel, k.Beschreibung, k.Private, k.Ueberkategorie, k.Fortschritt, k.KastenID)
	if err != nil {
		return err
	}
	return nil
}
