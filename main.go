package main

import (
	"database/sql"
	"fmt"

	_karteikarteRepo "github.com/Elbroto1993/web-ss19/app/controller/karteikarte/repository"
	_karteikarteUcase "github.com/Elbroto1993/web-ss19/app/controller/karteikarte/usecase"
	_karteikarteHttpDeliver "github.com/Elbroto1993/web-ss19/app/route/delivery_karte/http"

	_karteikastenRepo "github.com/Elbroto1993/web-ss19/app/controller/karteikasten/repository"
	_karteikastenUcase "github.com/Elbroto1993/web-ss19/app/controller/karteikasten/usecase"
	_karteikastenHttpDeliver "github.com/Elbroto1993/web-ss19/app/route/delivery_kasten/http"

	_userRepo "github.com/Elbroto1993/web-ss19/app/controller/user/repository"
	_userUcase "github.com/Elbroto1993/web-ss19/app/controller/user/usecase"
	_userHttpDeliver "github.com/Elbroto1993/web-ss19/app/route/delivery_user/http"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	///PROBIEREN
	_sessionHttpDeliver "github.com/Elbroto1993/web-ss19/app/controller/session"
	////
)

func main() {

	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/braintrain?parseTime=true")
	if err != nil {
		fmt.Println(err)
	}
	db.SetMaxIdleConns(10)
	defer db.Close()

	router := mux.NewRouter()

	// Subrouter setup for user api
	userRepo := _userRepo.NewMysqlUserRepository(db)
	userUsecase := _userUcase.NewUserUsecase(userRepo)
	userPrefix := router.PathPrefix("/user").Subrouter()
	_userHttpDeliver.NewUserHttpHandler(userPrefix, userUsecase)

	// Subrouter setup for karteikasten api
	karteikastenRepo := _karteikastenRepo.NewMysqlKarteikastenRepository(db)
	karteikastenUsecase := _karteikastenUcase.NewKarteikastenUsecase(karteikastenRepo)
	karteikastenPrefix := router.PathPrefix("/karteikasten").Subrouter()
	_karteikastenHttpDeliver.NewKarteikastenHttpHandler(karteikastenPrefix, karteikastenUsecase)

	// Subrouter setup for karteikarte api
	karteikarteRepo := _karteikarteRepo.NewMysqlKarteikarteRepository(db)
	karteikarteUsecase := _karteikarteUcase.NewKarteikarteUsecase(karteikarteRepo)
	karteikartePrefix := router.PathPrefix("/karteikarte").Subrouter()
	_karteikarteHttpDeliver.NewKarteikarteHttpHandler(karteikartePrefix, karteikarteUsecase)

	///PROBIEREN
	userPrefix = router.PathPrefix("/login").Subrouter()
	_sessionHttpDeliver.NewSessionHttpHandler(userPrefix, userUsecase)
	/////

	/* ONLY FOR TESTING */

	router.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
	/********************/

	log.Fatal(http.ListenAndServe("localhost:8080", router))
}
