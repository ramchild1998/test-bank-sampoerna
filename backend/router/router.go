package router

import (
	"backend/handlers"

	"github.com/gorilla/mux"
)

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
