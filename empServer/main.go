package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joshsoftware/emp/empServer/db"
	"github.com/joshsoftware/emp/empServer/user"
)

func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "PONG")
}

func init() {
	db.DBConnetion()
}

func startHttp() {
	var router = mux.NewRouter()

	router.HandleFunc("/ping", pingHandler)
	router.HandleFunc("/getemps", user.ReadAll).Methods(http.MethodGet)
	router.HandleFunc("/createemp", user.Create).Methods(http.MethodPost)
	router.HandleFunc("/updateemp/{id}", user.Update).Methods(http.MethodPut)
	router.HandleFunc("/deleteemp/{id}", user.Delete).Methods(http.MethodDelete)
	router.HandleFunc("/getemp/{id}", user.Read).Methods(http.MethodGet)

	http.Handle("/", router)

	http.ListenAndServe(":3000", nil)
}
func main() {
	fmt.Println("hallo server")
	defer db.GetDB().Close()
	startHttp()
}
