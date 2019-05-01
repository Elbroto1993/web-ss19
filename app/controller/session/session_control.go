package session

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Elbroto1993/web-ss19/app/controller/user"
	"github.com/Elbroto1993/web-ss19/app/model"

	"github.com/gorilla/mux"
	"github.com/gorilla/securecookie"
	"github.com/gorilla/sessions"
)

// HTTPSessionHandler represent the httphandler for sessions
type HTTPSessionHandler struct {
	UserUsecase user.Usecase
}

func NewSessionHttpHandler(r *mux.Router, us user.Usecase) {
	handler := &HTTPSessionHandler{
		UserUsecase: us,
	}
	r.HandleFunc("/", handler.Login).Methods("POST")
	r.HandleFunc("/logout", handler.Logout).Methods("GET")
	r.HandleFunc("/checkLoggedIn", handler.CheckLoggedIn).Methods("GET")
}

var Store = sessions.NewCookieStore(securecookie.GenerateRandomKey(32))

func init() {
	Store.Options = &sessions.Options{
		Path:     "/",
		Domain:   "localhost",
		HttpOnly: true,
	}
}

// Login user and set cookie's authenticated = true
func (u *HTTPSessionHandler) Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	session, err := Store.Get(r, "session")
	if err != nil {
		fmt.Println(err)
	}
	// Save data into user struct
	user := &model.User{}
	_ = json.NewDecoder(r.Body).Decode(&user)
	user, err = u.UserUsecase.GetByUsername(user.Username)
	if err != nil {
		fmt.Println(err)
	}
	checkValid := u.UserUsecase.Login(user)
	if checkValid {
		// Set user id and authenticated for this session
		fmt.Println(user.UserID)
		session.Values["user_id"] = user.UserID
		session.Values["username"] = user.Username
		session.Values["authenticated"] = true
		session.Save(r, w)
		json.NewEncoder(w).Encode(true)
	} else {
		json.NewEncoder(w).Encode(false)
		fmt.Println("Username oder Passwort ist falsch.")
	}
}

func (u *HTTPSessionHandler) Logout(w http.ResponseWriter, r *http.Request) {
	session, err := Store.Get(r, "session")
	if err != nil {
		fmt.Println(err)
	}
	session.Values["authenticated"] = false
	session.Save(r, w)
}

func (u *HTTPSessionHandler) CheckLoggedIn(w http.ResponseWriter, r *http.Request) {
	valid := GetCookieAuthenticated(w, r)
	json.NewEncoder(w).Encode(valid)
}

func GetCookieUserID(w http.ResponseWriter, r *http.Request) int64 {
	session, err := Store.Get(r, "session")
	if err != nil {
		fmt.Println(err)
	}
	return session.Values["user_id"].(int64)
}

func GetCookieUsername(w http.ResponseWriter, r *http.Request) string {
	session, err := Store.Get(r, "session")
	if err != nil {
		fmt.Println(err)
	}
	return session.Values["username"].(string)
}

func GetCookieAuthenticated(w http.ResponseWriter, r *http.Request) bool {
	session, err := Store.Get(r, "session")
	if err != nil {
		fmt.Println(err)
	}
	if session.Values["authenticated"] == nil || session.Values["authenticated"] == false {
		return false
	}
	return true
}
