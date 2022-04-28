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
	router.HandleFunc("/penyakit", controllers.GetAllPenyakit).Methods("GET")
	router.HandleFunc("/penyakit/penyakit={namaPenyakit}", controllers.GetPenyakitByNamaPenyakit).Methods("GET")

	router.HandleFunc("/hasil", controllers.AddHasil).Methods("POST")
	router.HandleFunc("/hasil", controllers.GetAllHasil).Methods("GET")
	router.HandleFunc("/hasil/tanggal={tanggal}", controllers.GetHasilByTanggal).Methods("GET")
	router.HandleFunc("/hasil/namapenyakit={namaPenyakit}", controllers.GetHasilByNamaPenyakit).Methods("GET")
	router.HandleFunc("/hasil/tanggal={tanggal}/namapenyakit={namaPenyakit}", controllers.GetHasilByTanggalAndNamaPenyakit).Methods("GET")
}

func main() {
	r := mux.NewRouter()
	initaliseHandlers(r)

	// Use default options
	handler := cors.Default().Handler(r)

	http.ListenAndServe(":8090", handler)
}