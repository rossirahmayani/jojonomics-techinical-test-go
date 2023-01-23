package cek_saldo_service

import (
	"encoding/json"
	"github.com/teris-io/shortid"
	db "jojonomics-techinical-test-go/config/database"
	_ "log"
	"net/http"
)


func CekSaldo(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	request := &CekSaldoRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		_ = json.NewEncoder(w).Encode(generateErrorResponse(err))
		return
	}

	response := cek(request.NoRek)
	_ = json.NewEncoder(w).Encode(response)
}

func cek(norek string) *CekSaldoResponse{
	dbCon := db.Connect()
	defer dbCon.Close()

	//var data db.Rekening;
	var data db.Rekening
	err := dbCon.QueryRow(`SELECT * FROM tbl_rekening WHERE norek = $1`, norek).Scan(
			&data.Id, &data.Norek, &data.CreateDate, &data.UpdateDate, &data.Saldo)

	if err != nil {
		return generateErrorResponse(err)
	}

	return &CekSaldoResponse{
		Error:   false,
		Data:    &CekSaldoData{
			NoRek: data.Norek,
			Saldo: data.Saldo,
		},
	}
}

func generateErrorResponse (err error) *CekSaldoResponse{
	var reffId string
	reffId, _ = shortid.Generate()
	
	return &CekSaldoResponse{
		Error:   true,
		ReffId:  reffId,
		Message: err.Error(),
	}
}

