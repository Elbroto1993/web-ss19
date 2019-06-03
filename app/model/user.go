package model

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	couchdb "github.com/leesper/couchdb-golang"
	"golang.org/x/crypto/bcrypt"
)

// User Struct
type User struct {
	Id        string    `json:"_id"`
	Rev       string    `json:"_rev"`
	Type      string    `json:"type"`
	Username  string    `json:"username"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"createdat"`
	couchdb.Document
}

// Add User
func (user User) Add() (err error) {
	// Check wether username already exists
	userInDB, err := GetUserByUsername(user.Username)
	if err == nil && userInDB.Username == user.Username {
		return errors.New("username exists already")
	}

	// Check wether email already exists
	userInDB, err = GetUserByEmail(user.Email)
	if err == nil && userInDB.Email == user.Email {
		return errors.New("email exists already")
	}

	// Hash password
	hashedPwd, err := bcrypt.GenerateFromPassword([]byte(user.Password), 14)
	b64HashedPwd := base64.StdEncoding.EncodeToString(hashedPwd)

	user.Password = b64HashedPwd
	user.Type = "User"

	// Convert Todo struct to map[string]interface as required by Save() method
	u, err := user2Map(user)

	// Delete _id and _rev from map, otherwise DB access will be denied (unauthorized)
	delete(u, "_id")
	delete(u, "_rev")

	// Add todo to DB
	_, _, err = btDB.Save(u, nil)

	if err != nil {
		fmt.Printf("[Add] error: %s", err)
	}

	return err
}

// Delete User by username
func DeleteUser(username string) (err error) {
	user, err := GetUserByUsername(username)
	if err != nil {
		return err
	}

	karten, err := GetEigeneKarten(username)
	if err != nil {
		return err
	}
	for i := 0; i < len(karten); i++ {
		btDB.Delete(karten[i].Id)
	}

	// Delete all kaesten from user
	kaesten, err := getEigeneKaesten(username)
	if err != nil {
		return err
	}
	for i := 0; i < len(kaesten); i++ {
		btDB.Delete(kaesten[i].Id)
	}

	// Delete user from DB
	err = btDB.Delete(user.Id)

	if err != nil {
		fmt.Println(err)
	}

	return err
}

// Update User
func (user User) Update() (err error) {

	// Convert Todo struct to map[string]interface as required by Save() method
	u, err := user2Map(user)

	// Add todo to DB
	err = btDB.Set(user.Id, u)

	if err != nil {
		fmt.Printf("[Update] error: %s", err)
	}

	return err
}

// GetUserByUsername retrieve User by username
func GetUserByUsername(username string) (user User, err error) {
	if username == "" {
		return User{}, errors.New("no username provided")
	}

	query := `
	{
		"selector": {
			 "type": "User",
			 "username": "%s"
		}
	}`
	u, err := btDB.QueryJSON(fmt.Sprintf(query, username))
	if err != nil || len(u) != 1 {
		return User{}, err
	}

	user, err = map2User(u[0])
	if err != nil {
		return User{}, err
	}

	return user, nil
}

// GetUserByEmail retrieve User by Email
func GetUserByEmail(email string) (user User, err error) {
	if email == "" {
		return User{}, errors.New("no email provided")
	}

	query := `
	{
		"selector": {
			 "type": "User",
			 "email": "%s"
		}
	}`
	u, err := btDB.QueryJSON(fmt.Sprintf(query, email))
	if err != nil || len(u) != 1 {
		return User{}, err
	}

	user, err = map2User(u[0])
	if err != nil {
		return User{}, err
	}

	return user, nil
}

// ---------------------------------------------------------------------------
// Internal helper functions
// ---------------------------------------------------------------------------

// Convert from User struct to map[string]interface{} as required by golang-couchdb methods
func user2Map(u User) (user map[string]interface{}, err error) {
	uJSON, err := json.Marshal(u)
	json.Unmarshal(uJSON, &user)

	return user, err
}

// Convert from map[string]interface{} to User struct as required by golang-couchdb methods
func map2User(user map[string]interface{}) (u User, err error) {
	uJSON, err := json.Marshal(user)
	json.Unmarshal(uJSON, &u)

	return u, err
}
