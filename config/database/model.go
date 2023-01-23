package database

import (
	"math/big"
	"time"
)

type Rekening struct {
	Id int `db:"id"`
	Norek string `db:"norek"`
	Saldo int `db:"saldo"`
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
	Harga int `db:"harga"`
	Gram big.Float `db:"gram"`
	Type string `db:"type"`
	SaldoAkhir big.Float `db:"saldo_akhir"`
	TransactionDate time.Time `db:"transaction_date"`
	CreateDate time.Time `db:"create_date"`
	UpdateDate time.Time `db:"update_date"`
}