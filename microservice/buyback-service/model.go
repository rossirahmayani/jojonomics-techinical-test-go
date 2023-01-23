package buyback_service

type BuybackRequest struct {
	Gram float32 `json:"gram"`
	HargaBuyback int `json:"harga"`
	Norek string `json:"norek"`
}
