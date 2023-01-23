package cek_mutasi_service

import (
	"time"
)

type CekMutasiRequest struct {
	NoRek string `json:"norek"`
	StartDate time.Time `json:"start_date"`
	EndDate time.Time `json:"end_date"`
}

type CekMutasiResponse struct {
	Error bool `json:"error"`
	ReffId string `json:"reff_id"`
	Message string `json:"message"`
	Data *[]CekMutasiData `json:"data"`
}

type CekMutasiData struct {
	MutationDate time.Time `json:"date"`
	MutationType string `json:"type"`
	Gram float32 `json:"gram"`
	HargaTopup int `json:"harga_topup"`
	HargaBuyback int `json:"harga_buyback"`
	Saldo float32 `json:"saldo"`
}