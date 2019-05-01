package repository

import (
	"database/sql"

	"github.com/Elbroto1993/web-ss19/app/controller/user"
	"github.com/Elbroto1993/web-ss19/app/model"
)

type mysqlUserRepository struct {
	Conn *sql.DB
}

// NewMysqlUserRepository will create an object that represent the user.Repository interface
func NewMysqlUserRepository(Conn *sql.DB) user.Repository {
	return &mysqlUserRepository{Conn}
}

func (m *mysqlUserRepository) Fetch() ([]model.User, error) {
	var allUser []model.User
	// Create Prepared Statement
	stmt, err := m.Conn.Prepare("SELECT user_id, username, password, e_mail, created_at FROM user")
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
		var user model.User
		err := rows.Scan(&user.UserID, &user.Username, &user.Password, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		allUser = append(allUser, user)
	}
	return allUser, nil
}

func (m *mysqlUserRepository) GetByID(id int64) (*model.User, error) {
	user := &model.User{}
	// Create Prepared Statement
	stmt, err := m.Conn.Prepare("SELECT user_id, username, password, e_mail, created_at FROM user WHERE user_id=?")
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
		err := rows.Scan(&user.UserID, &user.Username, &user.Password, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
	}
	return user, nil
}

func (m *mysqlUserRepository) GetByUsername(name string) (*model.User, error) {
	user := &model.User{}
	// Create Prepared Statement
	stmt, err := m.Conn.Prepare("SELECT user_id, username, password, e_mail, created_at FROM user WHERE username=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	// Query the stmt and save results in rows
	rows, err := stmt.Query(name)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// Loop through results and save in user
	for rows.Next() {
		err := rows.Scan(&user.UserID, &user.Username, &user.Password, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
	}
	return user, nil
}

func (m *mysqlUserRepository) GetByMail(mail string) (*model.User, error) {
	user := &model.User{}
	// Create Prepared Statement
	stmt, err := m.Conn.Prepare("SELECT user_id, username, password, e_mail, created_at FROM user WHERE e_mail=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	// Query the stmt and save results in rows
	rows, err := stmt.Query(mail)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	// Loop through results and save in user
	for rows.Next() {
		err := rows.Scan(&user.UserID, &user.Username, &user.Password, &user.Email, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
	}
	return user, nil
}

func (m *mysqlUserRepository) Store(u *model.User) error {

	// PreparedStatement for insert user
	stmt, err := m.Conn.Prepare("INSERT user SET username=?, password=?, e_mail=?, created_at=?")
	if err != nil {
		return err
	}
	defer stmt.Close()

	// Fill prepStmt with values and execute query
	res, err := stmt.Exec(u.Username, u.Password, u.Email, u.CreatedAt)
	if err != nil {
		return err
	}
	// Get last inserted id
	lastID, err := res.LastInsertId()
	u.UserID = lastID
	return nil
}

func (m *mysqlUserRepository) Delete(id int64) error {

	stmt, err := m.Conn.Prepare("DELETE FROM user WHERE user_id=?")
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

func (m *mysqlUserRepository) Update(u *model.User) error {
	// PreparedStatement for Update
	stmt, err := m.Conn.Prepare("UPDATE user set e_mail = ?, password = ? where user_id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	// Fill prepStmt with values and execute query
	_, err = stmt.Exec(u.Email, u.Password, u.UserID)
	if err != nil {
		return err
	}
	return nil
}
