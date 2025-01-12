package tests

import (
	"backend/handlers"
	"backend/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
)

func TestCreateTransaksi(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/transaksi", handlers.CreateTransaksi).Methods("POST")

	transaksi := models.Transaksi{NomorRekening: "12345", JenisTransaksi: "debit", JumlahTransaksi: 500}
	jsonTransaksi, _ := json.Marshal(transaksi)
	req, _ := http.NewRequest("POST", "/transaksi", bytes.NewBuffer(jsonTransaksi))
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	if response.Code != http.StatusOK {
		t.Errorf("Expected %v, got %v", http.StatusOK, response.Code)
	}
}
