// Package router berisi konfigurasi routing untuk API
package router

import (
	"backend/handlers"

	"github.com/gorilla/mux"
)

// InitRouter menginisialisasi dan mengkonfigurasi router HTTP
// Fungsi ini mendaftarkan semua endpoint API yang tersedia:
//
// Endpoint untuk Rekening:
// - POST /rekening - Membuat rekening baru
// - GET /rekening/{id} - Mengambil data rekening berdasarkan ID
// - PUT /rekening/{id} - Memperbarui data rekening berdasarkan ID
// - DELETE /rekening/{id} - Menghapus rekening berdasarkan ID
//
// Endpoint untuk Transaksi:
// - POST /transaksi - Membuat transaksi baru
// - GET /transaksi/{id} - Mengambil data transaksi berdasarkan ID
// - PUT /transaksi/{id} - Memperbarui data transaksi berdasarkan ID
// - DELETE /transaksi/{id} - Menghapus transaksi berdasarkan ID
//
// Mengembalikan pointer ke mux.Router yang telah dikonfigurasi
func InitRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/rekening", handlers.CreateRekening).Methods("POST")
	router.HandleFunc("/rekening/{id}", handlers.ReadRekening).Methods("GET")
	router.HandleFunc("/rekening/{id}", handlers.UpdateRekening).Methods("PUT")
	router.HandleFunc("/rekening/{id}", handlers.DeleteRekening).Methods("DELETE")

	router.HandleFunc("/transaksi", handlers.CreateTransaksi).Methods("POST")
	router.HandleFunc("/transaksi/{id}", handlers.ReadTransaksi).Methods("GET")
	router.HandleFunc("/transaksi/{id}", handlers.UpdateTransaksi).Methods("PUT")
	router.HandleFunc("/transaksi/{id}", handlers.DeleteTransaksi).Methods("DELETE")

	return router
}
