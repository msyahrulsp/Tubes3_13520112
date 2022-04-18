package controllers

import (
	"encoding/json"
	"log"
	"net/http"

	"db/database"
	"db/model"

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
	var arr_hasil []model.Hasil
	var response model.ResponseHasil

	db := database.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	  if err != nil {
		  panic(err)
	  }

	// Panggil algoritma pattern matching
	// DNA := r.FormValue("DNA")
	
	// persentase = algo pattern matching 
	persentase := 100 

	// hasil = algo pattern matching
	hasilTes := true
	
	_, err = db.Exec("INSERT INTO Hasil (Tanggal, NamaPengguna, NamaPenyakit, Persentase, Hasil) values (?,?)",
		r.FormValue("tanggal"),
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
		hasil.Tanggal = r.FormValue("tanggal")
		hasil.NamaPengguna = r.FormValue("namaPengguna")
		hasil.NamaPenyakit = r.FormValue("namaPenyakit")
		hasil.Persentase = persentase
		hasil.Hasil = hasilTes
		hasil.DNA = r.FormValue("DNA")
		arr_hasil = append(arr_hasil, hasil)

		response.Status = 200
		response.Message = "OK"
		response.Data = arr_hasil
		log.Print("Insert data to database")
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}