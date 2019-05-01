package http

/****** @TODO - ADD USEFUL ERROR HANDLING *******/
import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/Elbroto1993/web-ss19/app/controller/karteikarte"
	"github.com/Elbroto1993/web-ss19/app/controller/session"
	"github.com/Elbroto1993/web-ss19/app/model"
	validator "gopkg.in/go-playground/validator.v9"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

// HTTPUserHandler  represent the httphandler for user
type HTTPKarteikarteHandler struct {
	KarteikarteUsecase karteikarte.Usecase
}

func NewKarteikarteHttpHandler(r *mux.Router, us karteikarte.Usecase) {
	handler := &HTTPKarteikarteHandler{
		KarteikarteUsecase: us,
	}
	r.HandleFunc("/{id}", handler.GetByID).Methods("GET")
	r.HandleFunc("/kasten/{id}", handler.GetByKastenID).Methods("GET")
	r.HandleFunc("/", handler.FetchKarten).Methods("GET")
	r.HandleFunc("/", handler.Store).Methods("POST")
	r.HandleFunc("/update", handler.Update).Methods("POST")
	r.HandleFunc("/{id}", handler.Delete).Methods("DELETE")
}

// FetchKarten get all karten from all kaesten from all users
func (u *HTTPKarteikarteHandler) FetchKarten(w http.ResponseWriter, r *http.Request) {

	listKarten, err := u.KarteikarteUsecase.Fetch()
	if err != nil {
		fmt.Println(err)
	} else {
		json.NewEncoder(w).Encode(listKarten)
	}
}

// GetByID get karteikarte by id
func (u *HTTPKarteikarteHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Get all params
	id, _ := strconv.ParseInt(params["id"], 10, 64)
	karte, err := u.KarteikarteUsecase.GetByID(id)
	if err != nil {
		fmt.Println(err)
	} else {
		json.NewEncoder(w).Encode(karte)
	}
}

// GetByKastenID get all karteikarten from one kasten by kasten_id
func (u *HTTPKarteikarteHandler) GetByKastenID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r) // Get all params
	id, _ := strconv.ParseInt(params["id"], 10, 64)
	listKarten, err := u.KarteikarteUsecase.GetByKastenID(id)
	if err != nil {
		fmt.Println(err)
	} else {
		json.NewEncoder(w).Encode(listKarten)
	}
}

// Store karteikarte into kasten for logged in user, kasten_id is expected in struct
func (u *HTTPKarteikarteHandler) Store(w http.ResponseWriter, r *http.Request) {
	valid := session.GetCookieAuthenticated(w, r)
	if valid {
		var karte model.Karteikarte
		err := json.NewDecoder(r.Body).Decode(&karte)
		if err != nil {
			fmt.Println(err)
		}
		// If struct isn't valid don't store it
		if ok, err := isRequestValid(&karte); !ok {
			fmt.Println(err)
		} else {
			err = u.KarteikarteUsecase.Store(&karte)
			if err != nil {
				fmt.Println(err)
			} else {
				json.NewEncoder(w).Encode(karte)
			}
		}
	} else {
		fmt.Println("Nicht eingeloggt.")
	}
}

// Delete karteikarte from kasten for logged in user
func (u *HTTPKarteikarteHandler) Delete(w http.ResponseWriter, r *http.Request) {
	valid := session.GetCookieAuthenticated(w, r)
	if valid {
		params := mux.Vars(r) // Get all params
		id, err := strconv.Atoi(params["id"])
		if err != nil {
			fmt.Println(err)
		} else {
			err = u.KarteikarteUsecase.Delete(int64(id))
			if err != nil {
				fmt.Println(err)
			} else {
				json.NewEncoder(w).Encode(http.StatusNoContent)
			}
		}
	} else {
		fmt.Println("Nicht eingeloggt.")
	}
}

// Update karteikarte for logged in user, kasten_id is expected in struct
func (u *HTTPKarteikarteHandler) Update(w http.ResponseWriter, r *http.Request) {
	valid := session.GetCookieAuthenticated(w, r)
	if valid {
		var karte model.Karteikarte
		err := json.NewDecoder(r.Body).Decode(&karte)
		if err != nil {
			fmt.Println(err)
		}
		// If struct isn't valid don't store it
		if ok, err := isRequestValid(&karte); !ok {
			fmt.Println(err)
		} else {
			err = u.KarteikarteUsecase.Update(&karte)
			if err != nil {
				fmt.Println(err)
			} else {
				json.NewEncoder(w).Encode(karte)
			}
		}
	} else {
		fmt.Println("Nicht eingeloggt.")
	}
}

func isRequestValid(m *model.Karteikarte) (bool, error) {

	validate := validator.New()

	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}
