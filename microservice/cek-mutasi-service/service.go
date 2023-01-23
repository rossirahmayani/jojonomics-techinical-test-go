package cek_mutasi_service

import (
	"encoding/json"
	"github.com/teris-io/shortid"
	db "jojonomics-techinical-test-go/config/database"
	"net/http"
)

func CekMutasi(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	request := &CekMutasiRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	if err != nil {
		_ = json.NewEncoder(w).Encode(generateErrorResponse(err))
		return
	}

	response := cek(request)
	_ = json.NewEncoder(w).Encode(response)
}

func cek(request *CekMutasiRequest) *CekMutasiResponse{
	dbCon := db.Connect()
	defer dbCon.Close()
	var mutations []CekMutasiData
	//var rekening db.Rekening
	//err := dbCon.QueryRow(`SELECT * FROM tbl_rekening WHERE norek = $1`, request.NoRek).Scan(
	//	&rekening.Id, &rekening.Norek, &rekening.CreateDate, &rekening.UpdateDate, &rekening.Saldo)
	//
	//if err != nil {
	//	return generateErrorResponse(err)
	//}

	rows, err := dbCon.Query(`SELECT tbl_transaksi.* from tbl_transaksi join tbl_rekening on tbl_rekening.id = tbl_transaksi.rekening_id 
					where tbl_rekening.norek = $1 and tbl_transaksi.transaction_date between $2 and $3 
					order by tbl_transaksi.id desc`, request.NoRek, request.StartDate, request.EndDate)
	if err != nil {
		return generateErrorResponse(err)
	}



	for rows.Next() {
		var trx db.Transaksi
		err = rows.Scan(&trx.Id, &trx.RekeningId, &trx.HargaBuyback, &trx.Gram, &trx.TransactionType,
			&trx.ReffId, &trx.SaldoAkhir, &trx.TransactionDate, &trx.CreateDate, &trx.UpdateDate, &trx.HargaTopup, &trx.TransactionAmount)

		if err != nil {
			return generateErrorResponse(err)
		}

		mutation := CekMutasiData{
			MutationDate: trx.TransactionDate,
			MutationType: trx.TransactionType,
			Gram:         trx.Gram,
			HargaTopup:   trx.HargaTopup,
			HargaBuyback: trx.HargaBuyback,
			Saldo:        trx.SaldoAkhir,
		}

		mutations = append(mutations, mutation)
	}

	return &CekMutasiResponse{
		Error:   false,
		Data:    &mutations,
	}
}

func generateErrorResponse (err error) *CekMutasiResponse{
	var reffId string
	reffId, _ = shortid.Generate()

	return &CekMutasiResponse{
		Error:   true,
		ReffId:  reffId,
		Message: err.Error(),
	}
}
