// Package database menyediakan fungsi-fungsi untuk koneksi dan manajemen database PostgreSQL
package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// DB adalah variabel global yang menyimpan koneksi database GORM
var DB *gorm.DB

// createDatabaseIfNotExist membuat database baru jika belum ada
// Fungsi ini akan mencoba membuat database bank_db
// Jika database sudah ada, fungsi akan menampilkan pesan bahwa database sudah ada
// Jika terjadi error lain, fungsi akan menghentikan program
func createDatabaseIfNotExist() {
	connStr := "host=postgres user=postgres password=root dbname=bank_db port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("failed to connect to PostgreSQL:", err)
	}
	defer db.Close()

	_, err = db.Exec("CREATE DATABASE bank_db")
	if err != nil {
		if err.Error() == `pq: database "bank_db" already exists` {
			fmt.Println("Database already exists")
		} else {
			log.Fatal("failed to create database:", err)
		}
	} else {
		fmt.Println("Database created successfully")
	}
}

// Connect menginisialisasi koneksi ke database PostgreSQL menggunakan GORM
// Fungsi ini akan:
// 1. Memanggil createDatabaseIfNotExist untuk memastikan database tersedia
// 2. Membuat koneksi ke database menggunakan GORM
// 3. Menyimpan koneksi ke variabel global DB
// Jika terjadi error saat koneksi, fungsi akan menghentikan program
func Connect() {
	createDatabaseIfNotExist()

	dsn := "host=postgres user=postgres password=root dbname=bank_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	DB = db
}
