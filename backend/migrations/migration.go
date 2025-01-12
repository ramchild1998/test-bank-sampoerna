package migrations

import (
	"backend/database"
	"backend/models"
)

func Migrate() {
	database.DB.AutoMigrate(&models.Rekening{}, &models.Transaksi{})
}
