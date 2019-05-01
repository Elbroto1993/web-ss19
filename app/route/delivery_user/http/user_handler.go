package http

/****** @TODO - ADD USEFUL ERROR HANDLING *******/
import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/Elbroto1993/web-ss19/app/model"
	// "github.com/Elbroto1993/web-ss19/app/session"
	"github.com/Elbroto1993/web-ss19/app/controller/user"
	validator "gopkg.in/go-playground/validator.v9"

	"github.com/Elbroto1993/web-ss19/app/controller/session"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

// HTTPUserHandler represent the httphandler for user
type HTTPUserHandler struct {
	UserUsecase user.Usecase
}

func NewUserHttpHandler(r *mux.Router, us user.Usecase) {
	handler := &HTTPUserHandler{
		UserUsecase: us,
	}
	r.HandleFunc("/id", handler.GetByCookieID).Methods("GET")
	r.HandleFunc("/id/{id}", handler.GetByID).Methods("GET")
	r.HandleFunc("/getAllUsersLength", handler.GetAllUsersLength).Methods("GET")
	r.HandleFunc("/", handler.Store).Methods("POST")
	r.HandleFunc("/update", handler.Update).Methods("POST")
	r.HandleFunc("/", handler.Delete).Methods("DELETE")
}

// GetByID without password and email
func (u *HTTPUserHandler) GetByID(w http.ResponseWriter, r *http.Request) {

	params := mux.Vars(r) // Get all params
	id, _ := strconv.ParseInt(params["id"], 10, 64)
	user, err := u.UserUsecase.GetByID(id)
	if err != nil {
		fmt.Println(err)
	} else {
		user.Password = ""
		user.Email = ""
		json.NewEncoder(w).Encode(user)
	}
}

// GetByCookieID only for logged in users
func (u *HTTPUserHandler) GetByCookieID(w http.ResponseWriter, r *http.Request) {
	valid := session.GetCookieAuthenticated(w, r)
	if valid {
		id := session.GetCookieUserID(w, r)
		user, err := u.UserUsecase.GetByID(id)
		if err != nil {
			fmt.Println(err)
		} else {
			json.NewEncoder(w).Encode(user)
		}
	} else {
		fmt.Println("Nicht eingeloggt")
	}
}

// Store to register new user
func (u *HTTPUserHandler) Store(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		fmt.Println(err)
	}
	// If struct isn't valid don't store it
	if ok, err := isRequestValid(&user); !ok {
		fmt.Println(err)
	} else {
		err = u.UserUsecase.Store(&user)
		if err != nil {
			json.NewEncoder(w).Encode(err.Error())
		} else {
			json.NewEncoder(w).Encode(user)
		}
	}
}

// Delete user where user id = cookie user id
func (u *HTTPUserHandler) Delete(w http.ResponseWriter, r *http.Request) {
	valid := session.GetCookieAuthenticated(w, r)
	if valid {
		id := session.GetCookieUserID(w, r)
		err := u.UserUsecase.Delete(int64(id))
		if err != nil {
			fmt.Println(err)
		} else {
			json.NewEncoder(w).Encode(http.StatusNoContent)
		}
	} else {
		fmt.Println("Nicht eingeloggt")
	}
}

// Update user where user id = cookie user id
func (u *HTTPUserHandler) Update(w http.ResponseWriter, r *http.Request) {
	valid := session.GetCookieAuthenticated(w, r)
	if valid {
		id := session.GetCookieUserID(w, r)
		var user model.User
		err := json.NewDecoder(r.Body).Decode(&user)
		if err != nil {
			fmt.Println(err)
		}
		user.UserID = id
		// If struct isn't valid don't store it
		if ok, err := isRequestValid(&user); !ok {
			fmt.Println(err)
		} else {
			err = u.UserUsecase.Update(&user)
			if err != nil {
				json.NewEncoder(w).Encode(err.Error())
			} else {
				json.NewEncoder(w).Encode(user)
			}
		}
	} else {
		fmt.Println("Nicht eingloggt")
	}
}

func isRequestValid(m *model.User) (bool, error) {

	validate := validator.New()

	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}

//HELPER FUNCTIONS
func (u *HTTPUserHandler) GetAllUsersLength(w http.ResponseWriter, r *http.Request) {
	res, err := u.UserUsecase.GetAllUsersLength()
	if err != nil {
		fmt.Println(err)
	}
	json.NewEncoder(w).Encode(res)
}
