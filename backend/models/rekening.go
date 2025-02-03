// Package models berisi definisi model data untuk aplikasi bank
package models

import (
	"time"
)

// Rekening merepresentasikan model data untuk rekening bank
// Struct ini memiliki beberapa field:
// - ID: Primary key untuk identifikasi unik rekening
// - NamaPemilik: Nama pemilik rekening
// - NomorRekening: Nomor rekening yang unik
// - Saldo: Jumlah uang dalam rekening
// - TanggalPembuatan: Waktu pembuatan rekening
type Rekening struct {
	ID               uint      `gorm:"primaryKey"`      // ID unik rekening
	NamaPemilik      string    `gorm:"not null"`        // Nama pemilik rekening, tidak boleh kosong
	NomorRekening    string    `gorm:"unique;not null"` // Nomor rekening unik, tidak boleh kosong
	Saldo            float64   `gorm:"not null"`        // Saldo rekening, tidak boleh kosong
	TanggalPembuatan time.Time `gorm:"not null"`        // Tanggal pembuatan rekening, tidak boleh kosong
}
