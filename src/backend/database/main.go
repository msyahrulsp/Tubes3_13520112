package main

import (
	"fmt"
	"log"
	"net/http"

	"db/controllers"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

// func main() {
// 	initDB()
// 	log.Println("Starting the HTTP server on port 8090")

// 	router := mux.NewRouter().StrictSlash(true)
// 	initaliseHandlers(router)
// 	log.Fatal(http.ListenAndServe(":8090", router))
// }

func initaliseHandlers(router *mux.Router) {
	router.HandleFunc("/penyakit", controllers.AddPenyakit).Methods("POST")
	router.HandleFunc("/penyakit", controllers.GetAllPenyakit).Methods("GET")
	router.HandleFunc("/penyakit/penyakit={namaPenyakit}", controllers.GetPenyakitByNamaPenyakit).Methods("GET")

	router.HandleFunc("/hasil", controllers.AddPenyakit).Methods("POST")
	router.HandleFunc("/hasil", controllers.GetAllHasil).Methods("GET")
	router.HandleFunc("/hasil/tanggal={tanggal}", controllers.GetHasilByTanggal).Methods("GET")
	router.HandleFunc("/hasil/namapenyakit={namaPenyakit}", controllers.GetHasilByNamaPenyakit).Methods("GET")
	router.HandleFunc("/hasil/tanggal={tanggal}/namapenyakit={namaPenyakit}", controllers.GetHasilByTanggalAndNamaPenyakit).Methods("GET")
}

func main() {

	router := mux.NewRouter()
	// router.HandleFunc("/penyakit", getAllPenyakit).Methods("GET")
	initaliseHandlers(router)
	http.Handle("/", router)
	fmt.Println("Connected to port 8090")
	log.Fatal(http.ListenAndServe(":8090", router))
}