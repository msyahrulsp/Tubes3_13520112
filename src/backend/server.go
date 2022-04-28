package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"backend/controllers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"

	"github.com/rs/cors"
)

func initaliseHandlers(router *mux.Router) {
	router.HandleFunc("/", running)

	router.HandleFunc("/penyakit", controllers.AddPenyakit).Methods("POST")

	router.HandleFunc("/hasil", controllers.AddHasil).Methods("POST")
	router.HandleFunc("/hasil/find", controllers.GetHasilByQuery).Methods("POST")
}

func running (w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Server is running...")
}

func main() {
	r := mux.NewRouter()
	initaliseHandlers(r)

	// Use default options
	handler := cors.Default().Handler(r)

	port := os.Getenv("PORT")
	// http.ListenAndServe(":8090", handler)
	log.Print("Listening on :" + port)
	log.Fatal(http.ListenAndServe(":" + port, handler))
}