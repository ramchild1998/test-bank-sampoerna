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

var DB *gorm.DB

func createDatabaseIfNotExist() {
	connStr := "host=localhost user=postgres password=root dbname=bank_db port=5432 sslmode=disable"
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

func Connect() {
	createDatabaseIfNotExist()

	dsn := "host=localhost user=postgres password=root dbname=bank_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		log.Fatal("failed to connect to database:", err)
	}
	DB = db
}
