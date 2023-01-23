package input_harga_service

type InputHargaRequest struct {
	AdminId string `json:"admin_id"`
	HargaBuyback int `json:"harga_buyback"`
	HargaTopup int `json:"harga_topup"`
}
