package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	cek_harga_service "jojonomics-techinical-test-go/microservice/cek-harga-service"
	cek_saldo_service "jojonomics-techinical-test-go/microservice/cek-saldo-service"
	"log"
	"net/http"
)

func main() {
	if err:= godotenv.Load("misc/.env"); err != nil {
		log.Fatal(err)
	}


	r := mux.NewRouter();
	r.HandleFunc("/api/check-harga", cek_harga_service.CekHarga).Methods("GET")
	r.HandleFunc("/api/saldo", cek_saldo_service.CekSaldo).Methods("GET")

	log.Fatal(http.ListenAndServe(":8000", r))
}
