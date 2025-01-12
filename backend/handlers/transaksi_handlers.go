package handlers

import (
	"backend/database"
	"backend/models"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func CreateTransaksi(w http.ResponseWriter, r *http.Request) {
	var transaksi models.Transaksi
	json.NewDecoder(r.Body).Decode(&transaksi)
	transaksi.TanggalTransaksi = time.Now()
	database.DB.Create(&transaksi)
	json.NewEncoder(w).Encode(transaksi)
}

func ReadTransaksi(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var transaksi models.Transaksi
	database.DB.First(&transaksi, params["id"])
	json.NewEncoder(w).Encode(transaksi)
}

func UpdateTransaksi(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var transaksi models.Transaksi
	database.DB.First(&transaksi, params["id"])
	json.NewDecoder(r.Body).Decode(&transaksi)
	database.DB.Save(&transaksi)
	json.NewEncoder(w).Encode(transaksi)
}

func DeleteTransaksi(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var transaksi models.Transaksi
	database.DB.Delete(&transaksi, params["id"])
	w.WriteHeader(http.StatusNoContent)
}
