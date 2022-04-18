package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"db/database"
	"db/model"

	"github.com/gorilla/mux"
)

func GetAllPenyakit(w http.ResponseWriter, r *http.Request) {
	var penyakit model.Penyakit
	var arr_penyakit []model.Penyakit
	var response model.ResponsePenyakit

	db := database.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT NamaPenyakit, SequenceDNA FROM Penyakit")
	if err != nil {
		log.Print(err)
		response.Status = 404
		response.Message = "Not Found"
	} else {
		for rows.Next() {
			if err := rows.Scan(&penyakit.NamaPenyakit, &penyakit.SequenceDNA); err != nil {
				log.Fatal(err.Error())
			} else {
				arr_penyakit = append(arr_penyakit, penyakit)
			}
		}
		response.Status = 200
		response.Message = "OK"
		response.Data = arr_penyakit
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetPenyakitByNamaPenyakit(w http.ResponseWriter, r *http.Request) {
	var penyakit model.Penyakit
	var arr_penyakit []model.Penyakit
	var response model.ResponsePenyakit

	db := database.Connect()
	defer db.Close()

	vars := mux.Vars(r)

	rows, err := db.Query("SELECT NamaPenyakit, SequenceDNA FROM Penyakit WHERE NamaPenyakit = ?", vars["namaPenyakit"])
	if err != nil {
		log.Print(err)
		response.Status = 404
		response.Message = "Not Found"
	} else {
		for rows.Next() {
			if err := rows.Scan(&penyakit.NamaPenyakit, &penyakit.SequenceDNA); err != nil {
				log.Fatal(err.Error())
			} else {
				arr_penyakit = append(arr_penyakit, penyakit)
			}
		}
		response.Status = 200
		response.Message = "OK"
		response.Data = arr_penyakit
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func AddPenyakit(w http.ResponseWriter, r *http.Request) {
	var penyakit model.Penyakit
	var arr_penyakit []model.Penyakit
	var response model.ResponsePenyakit

	db := database.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	  if err != nil {
		  panic(err)
	  }

	_, err = db.Exec("INSERT INTO Penyakit (NamaPenyakit, SequenceDNA) values (?,?)",
		r.FormValue("namaPenyakit"),
		r.FormValue("sequenceDNA"),
	)

	if err != nil {
		log.Print(err)
		response.Status = 404
		response.Message = "Not Found"
	} else {
		penyakit.NamaPenyakit = r.FormValue("namaPenyakit")
		penyakit.SequenceDNA = r.FormValue("sequenceDNA")
		arr_penyakit = append(arr_penyakit, penyakit)

		response.Status = 200
		response.Message = "OK"
		response.Data = arr_penyakit
		log.Print("Insert data to database")
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}