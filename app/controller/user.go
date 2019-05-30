package controller

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/Elbroto1993/web-ss19-w-template/app/model"
	"net/http"
	"time"

	"github.com/gorilla/sessions"
	"golang.org/x/crypto/bcrypt"
)

var store *sessions.CookieStore

func init() {
	// key must be 16, 24 or 32 bytes long (AES-128, AES-192 or AES-256)
	key := make([]byte, 32)
	rand.Read(key)
	store = sessions.NewCookieStore(key)
}

// AddUser controller
func AddUser(w http.ResponseWriter, r *http.Request) {
	username := r.FormValue("username")
	password := r.FormValue("password")
	email := r.FormValue("email")

	user := model.User{}
	user.Username = username
	user.Password = password
	user.Email = email
	user.CreatedAt = time.Now()

	err := user.Add()
	if err != nil {
		data := struct {
			ErrorMsg string
		}{
			ErrorMsg: err.Error(),
		}
		tmpl.ExecuteTemplate(w, "register.tmpl", data)
	} else {
		Index(w, r)
	}
}

// DeleteUser controller
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	username := session.Values["username"].(string)

	err := model.DeleteUser(username)

	if err != nil {
		fmt.Println(err)
	}

	Index(w, r)
}

// UpdateUser controller
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")
	username := session.Values["username"].(string)
	password := r.FormValue("password")
	email := r.FormValue("email")

	user := model.User{}
	user.Username = username
	user.Password = password
	user.Email = email

	err := user.Update()
	if err != nil {
		data := struct {
			ErrorMsg string
		}{
			ErrorMsg: err.Error(),
		}
		tmpl.ExecuteTemplate(w, "profil.tmpl", data)
	} else {
		tmpl.ExecuteTemplate(w, "profil.tmpl", nil)
	}
}

// Login controller
func Login(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "index.tmpl", nil)
}

// AuthenticateUser controller
func AuthenticateUser(w http.ResponseWriter, r *http.Request) {
	var err error
	var user = model.User{}
	errorMsg := "Username and/or password wrong!"
	username := r.FormValue("username")
	password := r.FormValue("password")

	// Authentication
	user, err = model.GetUserByUsername(username)
	if err == nil {
		// decode base64 String to []byte
		passwordDB, _ := base64.StdEncoding.DecodeString(user.Password)
		err = bcrypt.CompareHashAndPassword(passwordDB, []byte(password))
		if err == nil {
			session, _ := store.Get(r, "session")

			// Set user as authenticated
			session.Values["authenticated"] = true
			session.Values["username"] = username
			session.Save(r, w)
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			data, err := model.GetIndexData(username)
			if err != nil {
				fmt.Println(err)
			}
			data.ErrorMsg = errorMsg
			tmpl.ExecuteTemplate(w, "index.tmpl", data)
		}
	} else {
		Index(w, r)
	}
}

// Logout controller
func Logout(w http.ResponseWriter, r *http.Request) {
	session, _ := store.Get(r, "session")

	session.Values["authenticated"] = false
	session.Values["username"] = ""
	session.Save(r, w)

	Index(w, r)
}

// Auth is an authentication handler
func Auth(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session, _ := store.Get(r, "session")

		// Check if user is authenticated
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			http.Redirect(w, r, "/", http.StatusFound)
		} else {
			h(w, r)
		}
	}
}
