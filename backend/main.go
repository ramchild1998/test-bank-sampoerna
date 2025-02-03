// Package main adalah package utama yang menjalankan aplikasi backend bank
package main

import (
	"backend/database"
	"backend/migrations"
	"backend/router"
	"fmt"
	"net/http"
)

// main adalah fungsi utama yang menginisialisasi dan menjalankan server HTTP
// Fungsi ini melakukan beberapa hal:
// 1. Menghubungkan ke database PostgreSQL menggunakan database.Connect()
// 2. Menjalankan migrasi database menggunakan migrations.Migrate()
// 3. Menginisialisasi router HTTP menggunakan router.InitRouter()
// 4. Menjalankan server HTTP pada port 8000
func main() {
	// Menghubungkan ke database
	database.Connect()

	// Menjalankan migrasi database
	migrations.Migrate()

	// Inisialisasi router
	r := router.InitRouter()

	// Menampilkan pesan server berjalan
	fmt.Println("Server is running on port 8000")

	// Menjalankan server HTTP
	http.ListenAndServe(":8000", r)
}
