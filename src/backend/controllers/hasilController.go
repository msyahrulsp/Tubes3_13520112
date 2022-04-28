package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"backend/database"
	"backend/lib"
	"backend/model"

	"github.com/gorilla/mux"
)

func GetAllHasil(w http.ResponseWriter, r *http.Request) {
	var hasil model.Hasil
	var arr_hasil []model.Hasil
	var response model.ResponseHasil

	db := database.Connect()
	defer db.Close()

	rows, err := db.Query("SELECT Tanggal, NamaPengguna, NamaPenyakit, Persentase, Hasil FROM Hasil")
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
				response.Status = 200
				response.Message = "OK"
				response.Data = arr_hasil
			}
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetHasilByNamaPenyakit(w http.ResponseWriter, r *http.Request) {
	var hasil model.Hasil
	var arr_hasil []model.Hasil
	var response model.ResponseHasil

	db := database.Connect()
	defer db.Close()

	vars := mux.Vars(r)

	rows, err := db.Query("SELECT Tanggal, NamaPengguna, NamaPenyakit, Persentase, Hasil FROM Hasil WHERE NamaPenyakit = ?", vars["namaPenyakit"])
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

func GetHasilByTanggal(w http.ResponseWriter, r *http.Request) {
	var hasil model.Hasil
	var arr_hasil []model.Hasil
	var response model.ResponseHasil

	db := database.Connect()
	defer db.Close()

	vars := mux.Vars(r)

	// To Do: Manipulate string tanggal

	rows, err := db.Query("SELECT Tanggal, NamaPengguna, NamaPenyakit, Persentase, Hasil FROM Hasil WHERE Tanggal = ?", vars["tanggal"])
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

func GetHasilByTanggalAndNamaPenyakit(w http.ResponseWriter, r *http.Request) {
	var hasil model.Hasil
	var arr_hasil []model.Hasil
	var response model.ResponseHasil

	db := database.Connect()
	defer db.Close()

	vars := mux.Vars(r)

	// To Do: Manipulate string tanggal and nama penyakit

	rows, err := db.Query("SELECT Tanggal, NamaPengguna, NamaPenyakit, Persentase, Hasil FROM Hasil WHERE Tanggal = ? AND NamaPenyakit = ?", vars["tanggal"], vars["namaPenyakit"])
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

	if !lib.IsValidDNA(r.FormValue("DNA")) {
		response.Status = 400
		response.Message = "Bad Request"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}

	tanggal := time.Now().Format("2006-01-02")

	rows, err := db.Query("SELECT * FROM Penyakit WHERE NamaPenyakit = ?", r.FormValue("namaPenyakit"))

	if err != nil {
		log.Print(err)
		response.Status = 404
		response.Message = "Not Found"
	}

	for rows.Next() {
		if err := rows.Scan(&penyakit.NamaPenyakit, &penyakit.SequenceDNA); err != nil {
			log.Fatal(err.Error())
		}
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