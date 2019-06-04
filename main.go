package main

import (
	"github.com/Elbroto1993/web-ss19-w-template/app/controller"
	"net/http"
)

func main() {

	server := http.Server{Addr: ":8080"}

	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/edit", controller.Auth(controller.Edit))
	http.HandleFunc("/edit2", controller.Auth(controller.Edit2))
	http.HandleFunc("/karteikasten", controller.Karteikasten)
	http.HandleFunc("/lern", controller.Auth(controller.Lern))
	http.HandleFunc("/lern2", controller.Auth(controller.Lern2))
	http.HandleFunc("/meinekarteien", controller.Auth(controller.Meinekarteien))
	http.HandleFunc("/profil", controller.Auth(controller.Profil))
	http.HandleFunc("/register", controller.Register)
	http.HandleFunc("/view", controller.View)

	http.HandleFunc("/add-user", controller.AddUser)
	http.HandleFunc("/delete-user", controller.DeleteUser)
	http.HandleFunc("/update-user", controller.UpdateUser)
	http.HandleFunc("/authenticate-user", controller.AuthenticateUser)
	http.HandleFunc("/logout", controller.Logout)

	http.HandleFunc("/add-or-update-kasten", controller.AddOrUpdateKasten)
	http.HandleFunc("/delete-kasten", controller.DeleteKasten)

	http.HandleFunc("/add-or-update-karte", controller.AddOrUpdateKarte)
	http.HandleFunc("/update-karte-lern", controller.KarteRichtigOderFalsch)
	http.HandleFunc("/delete-karte", controller.DeleteKarte)

	server.ListenAndServe()
}
