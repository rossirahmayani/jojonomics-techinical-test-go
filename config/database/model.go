package database

import (
	"time"
)

type Rekening struct {
	Id int `db:"id"`
	Norek string `db:"norek"`
	Saldo float32 `db:"saldo"`
	CreateDate time.Time `db:"create_date"`
	UpdateDate time.Time `db:"update_date"`
}

type Harga struct {
	Id int `db:"id"`
	HargaTopup int `db:"harga_topup"`
	HargaBuyback int `db:"harga_buyback"`
	ReffId string `db:"reff_id"`
	IsActive bool `db:"is_active"`
	AdminId string `db:"admin_id"`
	CreateDate time.Time `db:"create_date"`
	UpdateDate time.Time `db:"update_date"`
}

type Transaksi struct {
	Id int `db:"id"`
	RekeningId int `db:"rekening_id"`
	HargaTopup int `db:"harga_topup"`
	HargaBuyback int `db:"harga_buyback"`
	Gram float32 `db:"gram"`
	TransactionType string `db:"type"`
	TransactionAmount int `db:"transaction_amount"`
	ReffId string `db:"reff_id"`
	SaldoAkhir float32 `db:"saldo_akhir"`
	TransactionDate time.Time `db:"transaction_date"`
	CreateDate time.Time `db:"create_date"`
	UpdateDate time.Time `db:"update_date"`
}