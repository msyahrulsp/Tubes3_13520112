package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strings"
	"time"

	"backend/database"
	"backend/lib"
	"backend/model"
)

func GetHasilByQuery(w http.ResponseWriter, r *http.Request) {
	var hasil model.Hasil
	var arr_hasil []model.Hasil
	var response model.ResponseHasil

	db := database.Connect()
	defer db.Close()

	q := lib.IsValidQuery(r.FormValue("query"))
	if (q == -1) {
		response.Status = 400
		response.Message = "Bad Request"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}
	
	if (q == 0) {
		date, err := time.Parse("02-01-2006", strings.ReplaceAll(r.FormValue("query"), "/", "-"))
		if err != nil {
			panic(err)
		}
		rows, err := db.Query("SELECT Tanggal, NamaPengguna, NamaPenyakit, Persentase, Hasil FROM Hasil WHERE Tanggal = ?", date.Format("2006-01-02"))
		if err != nil {
			log.Print(err)
			response.Status = 404
			response.Message = "Not Found"
		} else {
			for rows.Next() {
				if err := rows.Scan(&hasil.Tanggal, &hasil.NamaPengguna, &hasil.NamaPenyakit, &hasil.Persentase, &hasil.Hasil)
				err != nil {
					log.Fatal(err.Error())
				} else {
					arr_hasil = append(arr_hasil, hasil)
				}
			}
			response.Status = 200
			response.Message = "OK"
			response.Data = arr_hasil
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	if (q == 1) {
		rows, err := db.Query("SELECT Tanggal, NamaPengguna, NamaPenyakit, Persentase, Hasil FROM Hasil WHERE NamaPenyakit = ?", r.FormValue("query"))
		if err != nil {
			log.Print(err)
			response.Status = 404
			response.Message = "Not Found"
		} else {
			for rows.Next() {
				if err := rows.Scan(&hasil.Tanggal, &hasil.NamaPengguna, &hasil.NamaPenyakit, &hasil.Persentase, &hasil.Hasil)
				err != nil {
					log.Fatal(err.Error())
				} else {
					arr_hasil = append(arr_hasil, hasil)
				}
			}
			response.Status = 200
			response.Message = "OK"
			response.Data = arr_hasil
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	queries := strings.SplitN(r.FormValue("query"), " ", 2)
	date, err := time.Parse("02-01-2006", strings.ReplaceAll(queries[0], "/", "-"))
		if err != nil {
			panic(err)
		}
	rows, err := db.Query("SELECT Tanggal, NamaPengguna, NamaPenyakit, Persentase, Hasil FROM Hasil WHERE Tanggal = ? AND NamaPenyakit = ?", date.Format("2006-01-02"), queries[1])
	if err != nil {
		log.Print(err)
		response.Status = 404
		response.Message = "Not Found"
	} else {
		for rows.Next() {
			if err := rows.Scan(&hasil.Tanggal, &hasil.NamaPengguna, &hasil.NamaPenyakit, &hasil.Persentase, &hasil.Hasil)
			err != nil {
				log.Fatal(err.Error())
			} else {
				arr_hasil = append(arr_hasil, hasil)
			}
		}
		response.Status = 200
		response.Message = "OK"
		response.Data = arr_hasil
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func AddHasil(w http.ResponseWriter, r *http.Request) {
	var hasil model.Hasil
	var penyakit model.Penyakit
	var arr_hasil []model.Hasil
	var response model.ResponseHasil
	var persentase float32
	var hasilTes bool

	db := database.Connect()
	defer db.Close()

	if !lib.IsValidDNA(r.FormValue("DNA")) || r.FormValue("namaPengguna") == "" {
		response.Status = 400
		response.Message = "Bad Request"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	tanggal := time.Now().Format("2006-01-02")
	// tanggal := time.Now().Format("02-01-2006")

	rows, err := db.Query("SELECT * FROM Penyakit WHERE NamaPenyakit = ?", r.FormValue("namaPenyakit"))

	if err != nil {
		response.Status = 404
		response.Message = "Not Found"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}
	
	for rows.Next() {
		if err := rows.Scan(&penyakit.NamaPenyakit, &penyakit.SequenceDNA); err != nil {
			log.Fatal(err.Error())
		}
	}

	if (model.Penyakit{}) == penyakit {
		response.Status = 400
		response.Message = "Bad Request"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	// Tes KMP, kalo gagal pakai hamming
	kmp := lib.KmpMatch(r.FormValue("DNA"), penyakit.SequenceDNA)

	if (kmp == -1) {
		_, persentase = lib.HammingDistance(r.FormValue("DNA"), penyakit.SequenceDNA) 
	} else {
		persentase = 100
	}

	if (persentase >= 80) {
		hasilTes = true
	} else {
		hasilTes = false
	}
	
	_, err = db.Exec("INSERT INTO Hasil (Tanggal, NamaPengguna, NamaPenyakit, Persentase, Hasil) values (?,?,?,?,?)",
		tanggal,
		r.FormValue("namaPengguna"),
		r.FormValue("namaPenyakit"),
		persentase,
		hasilTes,
	)

	if err != nil {
		log.Print(err)
		response.Status = 404
		response.Message = "Not Found"
	} else {
		hasil.Tanggal = tanggal
		hasil.NamaPengguna = r.FormValue("namaPengguna")
		hasil.NamaPenyakit = r.FormValue("namaPenyakit")
		hasil.Persentase = persentase
		hasil.Hasil = hasilTes
		arr_hasil = append(arr_hasil, hasil)

		response.Status = 200
		response.Message = "OK"
		response.Data = arr_hasil
		log.Print("Insert data to database")
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}