package main

import (
	"net/http"

	"backend/controllers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"github.com/rs/cors"
)

func initaliseHandlers(router *mux.Router) {
	router.HandleFunc("/penyakit", controllers.AddPenyakit).Methods("POST")

	router.HandleFunc("/hasil", controllers.AddHasil).Methods("POST")
	router.HandleFunc("/hasil/find", controllers.GetHasilByQuery).Methods("POST")
}

func main() {
	r := mux.NewRouter()
	initaliseHandlers(r)

	// Use default options
	handler := cors.Default().Handler(r)

	http.ListenAndServe(":8090", handler)
}