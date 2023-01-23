package cek_harga_service

import (
	"encoding/json"
	"log"
	db "jojonomics-techinical-test-go/config/database"
	"net/http"
)

func CekHarga(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := cek()
	json.NewEncoder(w).Encode(response)
}

func cek() *CekHargaResponse{
	dbCon := db.Connect()
	defer dbCon.Close()
	var data db.Harga
	err := dbCon.QueryRow(`SELECT * FROM tbl_harga WHERE is_active = true ORDER by id desc limit 1`).Scan(
		&data.Id, &data.HargaTopup, &data.HargaBuyback, &data.ReffId, &data.IsActive, &data.AdminId, &data.CreateDate, &data.UpdateDate)

	if err != nil {
		log.Print(err)
		return &CekHargaResponse{Error: true}
	}

	return &CekHargaResponse{
				Error: false,
				Data: &CekHargaData{
					data.HargaBuyback, data.HargaTopup}}
}

