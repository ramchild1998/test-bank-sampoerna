package models

import (
    "time"
)

type Rekening struct {
    ID              uint      `gorm:"primaryKey"`
    NamaPemilik     string    `gorm:"not null"`
    NomorRekening   string    `gorm:"unique;not null"`
    Saldo           float64   `gorm:"not null"`
    TanggalPembuatan time.Time `gorm:"not null"`
}
