package handlers

import (
	"backend/database"
	"backend/models"
	"encoding/json"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

// CreateRekening menangani pembuatan rekening baru
// Menerima data rekening dalam format JSON melalui request body
// Mengembalikan data rekening yang berhasil dibuat dalam format JSON
func CreateRekening(w http.ResponseWriter, r *http.Request) {
	var rekening models.Rekening
	json.NewDecoder(r.Body).Decode(&rekening)
	rekening.TanggalPembuatan = time.Now()
	database.DB.Create(&rekening)
	json.NewEncoder(w).Encode(rekening)
}

// ReadRekening menangani pembacaan data rekening berdasarkan ID
// Menerima ID rekening melalui URL parameter
// Mengembalikan data rekening dalam format JSON
func ReadRekening(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var rekening models.Rekening
	database.DB.First(&rekening, params["id"])
	json.NewEncoder(w).Encode(rekening)
}

// UpdateRekening menangani pembaruan data rekening berdasarkan ID
// Menerima ID rekening melalui URL parameter dan data pembaruan dalam format JSON melalui request body
// Mengembalikan data rekening yang telah diperbarui dalam format JSON
func UpdateRekening(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var rekening models.Rekening
	database.DB.First(&rekening, params["id"])
	json.NewDecoder(r.Body).Decode(&rekening)
	database.DB.Save(&rekening)
	json.NewEncoder(w).Encode(rekening)
}

// DeleteRekening menangani penghapusan data rekening berdasarkan ID
// Menerima ID rekening melalui URL parameter
// Mengembalikan status 204 (No Content) jika berhasil
func DeleteRekening(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var rekening models.Rekening
	database.DB.Delete(&rekening, params["id"])
	w.WriteHeader(http.StatusNoContent)
}
