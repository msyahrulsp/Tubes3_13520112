package model

type Penyakit struct {
	NamaPenyakit string `form:"namaPenyakit" json:"namaPenyakit"`
	SequenceDNA  string `form:"sequenceDNA" json:"sequenceDNA"`
}

type ResponsePenyakit struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Penyakit
}