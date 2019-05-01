package http

/****** @TODO - ADD USEFUL ERROR HANDLING *******/
import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/Elbroto1993/web-ss19/app/controller/karteikasten"
	"github.com/Elbroto1993/web-ss19/app/model"
	validator "gopkg.in/go-playground/validator.v9"

	"github.com/Elbroto1993/web-ss19/app/controller/session"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

// HTTPKarteikastenHandler  represent the httphandler for karteikasten
type HTTPKarteikastenHandler struct {
	KarteikastenUsecase karteikasten.Usecase
}

func NewKarteikastenHttpHandler(r *mux.Router, us karteikasten.Usecase) {
	handler := &HTTPKarteikastenHandler{
		KarteikastenUsecase: us,
	}
	//r.HandleFunc("/", handler.GetByID).Methods("GET")
	r.HandleFunc("/user", handler.GetByUserID).Methods("GET")
	r.HandleFunc("/createdbyuser", handler.GetByCreatedUserID).Methods("GET")
	r.HandleFunc("/", handler.FetchKaesten).Methods("GET")
	r.HandleFunc("/kasten/{id}", handler.GetByID).Methods("GET")
	r.HandleFunc("/", handler.Store).Methods("POST")
	r.HandleFunc("/storeExistKasten", handler.StoreExistKasten).Methods("POST")
	r.HandleFunc("/update", handler.Update).Methods("POST")
	r.HandleFunc("/{id}", handler.Delete).Methods("DELETE")
}

// FetchKaesten get all Karteikaesten from all users
func (u *HTTPKarteikastenHandler) FetchKaesten(w http.ResponseWriter, r *http.Request) {
	listKaesten, err := u.KarteikastenUsecase.Fetch()
	if err != nil {
		fmt.Println(err)
	} else {
		json.NewEncoder(w).Encode(listKaesten)
	}
}

// GetByID one kasten by id
func (u *HTTPKarteikastenHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Get all params
	id, _ := strconv.ParseInt(params["id"], 10, 64)
	kasten, err := u.KarteikastenUsecase.GetByID(id)
	if err != nil {
		fmt.Println(err)
	} else {
		json.NewEncoder(w).Encode(kasten)
	}
}

// GetByUserID all karteikaesten from logged in user by userid
func (u *HTTPKarteikastenHandler) GetByUserID(w http.ResponseWriter, r *http.Request) {
	valid := session.GetCookieAuthenticated(w, r)
	if valid {
		userid := session.GetCookieUserID(w, r)
		listKaesten, err := u.KarteikastenUsecase.GetByUserID(userid)
		if err != nil {
		} else {
			json.NewEncoder(w).Encode(listKaesten)
		}
	} else {
		fmt.Println("Nicht eingeloggt")
	}
}

// GetByCreatedUserID all karteikaesten from logged in user by created_by_userid
func (u *HTTPKarteikastenHandler) GetByCreatedUserID(w http.ResponseWriter, r *http.Request) {
	valid := session.GetCookieAuthenticated(w, r)
	if valid {
		userid := session.GetCookieUserID(w, r)
		listKaesten, err := u.KarteikastenUsecase.GetByCreatedUserID(userid)
		if err != nil {
		} else {
			json.NewEncoder(w).Encode(listKaesten)
		}
	} else {
		fmt.Println("Nicht eingeloggt")
	}
}

// Store new karteikasten for logged in user
func (u *HTTPKarteikastenHandler) Store(w http.ResponseWriter, r *http.Request) {
	valid := session.GetCookieAuthenticated(w, r)
	if valid {
		// Get userid from cookie for foreign key from kasten
		userid := session.GetCookieUserID(w, r)
		var kasten model.Karteikasten
		kasten.UserID = userid
		kasten.CreatedByUserID = userid
		err := json.NewDecoder(r.Body).Decode(&kasten)
		if err != nil {
			fmt.Println(err)
		}
		// If struct isn't valid don't store it
		if ok, err := isRequestValid(&kasten); !ok {
			fmt.Println(err)
		} else {
			err = u.KarteikastenUsecase.Store(&kasten)
			if err != nil {
				fmt.Println(err)
			} else {
				json.NewEncoder(w).Encode(kasten)
			}
		}
	} else {
		fmt.Println("Nicht eingeloggt")
	}
}

// StoreExistKasten new karteikasten for logged in user from another user
func (u *HTTPKarteikastenHandler) StoreExistKasten(w http.ResponseWriter, r *http.Request) {
	valid := session.GetCookieAuthenticated(w, r)
	if valid {
		// Get userid from cookie for foreign key from kasten
		userid := session.GetCookieUserID(w, r)
		var kasten model.Karteikasten
		kasten.UserID = userid
		err := json.NewDecoder(r.Body).Decode(&kasten)
		if err != nil {
			fmt.Println(err)
		}
		// If struct isn't valid don't store it
		if ok, err := isRequestValid(&kasten); !ok {
			fmt.Println(err)
		} else {
			err = u.KarteikastenUsecase.Store(&kasten)
			if err != nil {
				fmt.Println(err)
			} else {
				json.NewEncoder(w).Encode(kasten)
			}
		}
	} else {
		fmt.Println("Nicht eingeloggt")
	}
}

// Delete Karteikasten by id for logged in user
func (u *HTTPKarteikastenHandler) Delete(w http.ResponseWriter, r *http.Request) {
	valid := session.GetCookieAuthenticated(w, r)
	if valid {
		params := mux.Vars(r) // Get all params
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			fmt.Println(err)
		} else {
			err = u.KarteikastenUsecase.Delete(int64(id))
			if err != nil {
				fmt.Println(err)
			} else {
				json.NewEncoder(w).Encode(http.StatusNoContent)
			}
		}
	} else {
		fmt.Println("Nicht eingeloggt")
	}
}

// Update karteikasten for logged in user, kasten_id to update is expected in struct
func (u *HTTPKarteikastenHandler) Update(w http.ResponseWriter, r *http.Request) {
	valid := session.GetCookieAuthenticated(w, r)
	if valid {
		// Get userid from cookie for foreign key from kasten
		userid := session.GetCookieUserID(w, r)
		var kasten model.Karteikasten
		kasten.UserID = userid
		err := json.NewDecoder(r.Body).Decode(&kasten)
		if err != nil {
			fmt.Println(err)
		}
		// If struct isn't valid don't store it
		if ok, err := isRequestValid(&kasten); !ok {
			fmt.Println(err)
		} else {
			err = u.KarteikastenUsecase.Update(&kasten)
			if err != nil {
				fmt.Println(err)
			} else {
				json.NewEncoder(w).Encode(kasten)
			}
		}
	} else {
		fmt.Println("Nicht eingeloggt")
	}
}

func isRequestValid(m *model.Karteikasten) (bool, error) {

	validate := validator.New()

	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}
