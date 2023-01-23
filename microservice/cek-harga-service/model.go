package cek_harga_service

type CekHargaResponse struct {
	Error bool `json:"error"`
	Data *CekHargaData `json:"data"`
}

type CekHargaData struct {
	HargaBuyback int `json:"harga_buyback"`
	HargaTopup int `json:"harga_topup"`
}
