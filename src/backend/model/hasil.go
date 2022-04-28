package model

type Hasil struct {
	Tanggal   	  string `form:"tanggal" json:"tanggal"`
	NamaPengguna  string `form:"namaPengguna" json:"namaPengguna"`
	NamaPenyakit  string `form:"namaPenyakit" json:"namaPenyakit"`
	Persentase    float32    `json:"persentase"`
	Hasil         bool   `json:"hasil"`
}

type ResponseHasil struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Hasil
}