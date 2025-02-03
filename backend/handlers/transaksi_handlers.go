package handlers

import (
	"backend/database"
	"backend/models"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// CreateTransaksi menangani pembuatan transaksi baru
// Menerima data transaksi dalam format JSON melalui request body
// Mengembalikan data transaksi yang berhasil dibuat dalam format JSON
func CreateTransaksi(w http.ResponseWriter, r *http.Request) {
	var transaksi models.Transaksi
	json.NewDecoder(r.Body).Decode(&transaksi)
	transaksi.TanggalTransaksi = time.Now()
	database.DB.Create(&transaksi)
	json.NewEncoder(w).Encode(transaksi)
}

// ReadTransaksi menangani pembacaan data transaksi berdasarkan ID
// Menerima ID transaksi melalui URL parameter
// Mengembalikan data transaksi dalam format JSON
func ReadTransaksi(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var transaksi models.Transaksi
	database.DB.First(&transaksi, params["id"])
	json.NewEncoder(w).Encode(transaksi)
}

// UpdateTransaksi menangani pembaruan data transaksi berdasarkan ID
// Menerima ID transaksi melalui URL parameter dan data pembaruan dalam format JSON melalui request body
// Mengembalikan data transaksi yang telah diperbarui dalam format JSON
func UpdateTransaksi(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var transaksi models.Transaksi
	database.DB.First(&transaksi, params["id"])
	json.NewDecoder(r.Body).Decode(&transaksi)
	database.DB.Save(&transaksi)
	json.NewEncoder(w).Encode(transaksi)
}

// DeleteTransaksi menangani penghapusan data transaksi berdasarkan ID
// Menerima ID transaksi melalui URL parameter
// Mengembalikan status 204 (No Content) jika berhasil
func DeleteTransaksi(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var transaksi models.Transaksi
	database.DB.Delete(&transaksi, params["id"])
	w.WriteHeader(http.StatusNoContent)
}
