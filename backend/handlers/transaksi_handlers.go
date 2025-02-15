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

	// Mulai transaksi database
	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Cari rekening terkait
	var rekening models.Rekening
	if err := tx.Where("nomor_rekening = ?", transaksi.NomorRekening).First(&rekening).Error; err != nil {
		tx.Rollback()
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Rekening tidak ditemukan"})
		return
	}

	// Update saldo rekening
	if transaksi.JenisTransaksi == "penarikan" {
		if rekening.Saldo < transaksi.JumlahTransaksi {
			tx.Rollback()
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Saldo tidak mencukupi"})
			return
		}
		rekening.Saldo -= transaksi.JumlahTransaksi
	} else if transaksi.JenisTransaksi == "setoran" {
		rekening.Saldo += transaksi.JumlahTransaksi
	} else {
		tx.Rollback()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Jenis transaksi tidak valid"})
		return
	}

	// Simpan perubahan
	if err := tx.Save(&rekening).Error; err != nil {
		tx.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Buat transaksi
	transaksi.TanggalTransaksi = time.Now()
	if err := tx.Create(&transaksi).Error; err != nil {
		tx.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tx.Commit()
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

	tx := database.DB.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Ambil transaksi lama
	var oldTransaksi models.Transaksi
	if err := tx.First(&oldTransaksi, params["id"]).Error; err != nil {
		tx.Rollback()
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Transaksi tidak ditemukan"})
		return
	}

	// Decode transaksi baru
	var newTransaksi models.Transaksi
	json.NewDecoder(r.Body).Decode(&newTransaksi)

	// Cari rekening lama
	var oldRekening models.Rekening
	if err := tx.Where("nomor_rekening = ?", oldTransaksi.NomorRekening).First(&oldRekening).Error; err != nil {
		tx.Rollback()
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(map[string]string{"error": "Rekening lama tidak ditemukan"})
		return
	}

	// Rollback perubahan saldo lama
	if oldTransaksi.JenisTransaksi == "penarikan" {
		oldRekening.Saldo += oldTransaksi.JumlahTransaksi
	} else if oldTransaksi.JenisTransaksi == "setoran" {
		oldRekening.Saldo -= oldTransaksi.JumlahTransaksi
	}

	// Cari rekening baru (jika nomor rekening berubah)
	var currentRekening models.Rekening
	if newTransaksi.NomorRekening != oldTransaksi.NomorRekening {
		if err := tx.Where("nomor_rekening = ?", newTransaksi.NomorRekening).First(&currentRekening).Error; err != nil {
			tx.Rollback()
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(map[string]string{"error": "Rekening baru tidak ditemukan"})
			return
		}
	} else {
		currentRekening = oldRekening
	}

	// Validasi transaksi baru
	if newTransaksi.JenisTransaksi == "penarikan" {
		if currentRekening.Saldo < newTransaksi.JumlahTransaksi {
			tx.Rollback()
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(map[string]string{"error": "Saldo tidak mencukupi"})
			return
		}
		currentRekening.Saldo -= newTransaksi.JumlahTransaksi
	} else if newTransaksi.JenisTransaksi == "setoran" {
		currentRekening.Saldo += newTransaksi.JumlahTransaksi
	} else {
		tx.Rollback()
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Jenis transaksi tidak valid"})
		return
	}

	// Simpan semua perubahan
	if err := tx.Save(&oldRekening).Error; err != nil {
		tx.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := tx.Save(&currentRekening).Error; err != nil {
		tx.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Update data transaksi
	oldTransaksi.NomorRekening = newTransaksi.NomorRekening
	oldTransaksi.JenisTransaksi = newTransaksi.JenisTransaksi
	oldTransaksi.JumlahTransaksi = newTransaksi.JumlahTransaksi

	if err := tx.Save(&oldTransaksi).Error; err != nil {
		tx.Rollback()
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	tx.Commit()
	json.NewEncoder(w).Encode(oldTransaksi)
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
