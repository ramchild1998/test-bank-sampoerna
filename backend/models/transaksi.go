// Package models berisi definisi model data untuk aplikasi bank
package models

import (
	"time"
)

// Transaksi merepresentasikan model data untuk transaksi bank
// Struct ini memiliki beberapa field:
// - ID: Primary key untuk identifikasi unik transaksi
// - NomorRekening: Nomor rekening yang melakukan transaksi
// - JenisTransaksi: Jenis transaksi (misal: setoran, penarikan, transfer)
// - JumlahTransaksi: Jumlah uang yang ditransaksikan
// - TanggalTransaksi: Waktu transaksi dilakukan
type Transaksi struct {
	ID               uint      `gorm:"primaryKey"` // ID unik transaksi
	NomorRekening    string    `gorm:"not null"`   // Nomor rekening, tidak boleh kosong
	JenisTransaksi   string    `gorm:"not null"`   // Jenis transaksi, tidak boleh kosong
	JumlahTransaksi  float64   `gorm:"not null"`   // Jumlah transaksi, tidak boleh kosong
	TanggalTransaksi time.Time `gorm:"not null"`   // Tanggal transaksi, tidak boleh kosong
}
