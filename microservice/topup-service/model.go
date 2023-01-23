package topup_service

type TopupRequest struct {
	Gram float32 `json:"gram"`
	HargaBuyback int `json:"harga"`
	Norek string `json:"norek"`
}
