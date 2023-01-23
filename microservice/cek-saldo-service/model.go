package cek_saldo_service

type CekSaldoRequest struct {
	NoRek string `json:"norek"`
}

type CekSaldoResponse struct {
	Error bool `json:"error"`
	ReffId string `json:"reff_id"`
	Message string `json:"message"`
	Data *CekSaldoData `json:"data"`
}

type CekSaldoData struct {
	NoRek string `json:"norek"`
	Saldo float32 `json:"saldo"`
}