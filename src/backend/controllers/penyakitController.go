package controllers

import (
	"backend/lib"
	"encoding/json"
	"log"
	"net/http"

	"backend/database"
	"backend/model"
)

func AddPenyakit(w http.ResponseWriter, r *http.Request) {
	var penyakit model.Penyakit
	var arr_penyakit []model.Penyakit
	var response model.ResponsePenyakit

	db := database.Connect()
	defer db.Close()

	if lib.IsValidDNA(r.FormValue("sequenceDNA")) {
		_, err := db.Exec("INSERT INTO Penyakit (NamaPenyakit, SequenceDNA) values (?,?)",
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
	} else {
		response.Status = 400
		response.Message = "Bad Request"
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}