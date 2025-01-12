package models

import (
    "time"
)

type Transaksi struct {
    ID             uint      `gorm:"primaryKey"`
    NomorRekening  string    `gorm:"not null"`
    JenisTransaksi string    `gorm:"not null"`
    JumlahTransaksi float64  `gorm:"not null"`
    TanggalTransaksi time.Time `gorm:"not null"`
}
