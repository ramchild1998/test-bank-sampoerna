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

func TestCreateRekening(t *testing.T) {
	router := mux.NewRouter()
	router.HandleFunc("/rekening", handlers.CreateRekening).Methods("POST")

	rekening := models.Rekening{NamaPemilik: "Test", NomorRekening: "12345", Saldo: 1000}
	jsonRekening, _ := json.Marshal(rekening)
	req, _ := http.NewRequest("POST", "/rekening", bytes.NewBuffer(jsonRekening))
	response := httptest.NewRecorder()
	router.ServeHTTP(response, req)

	if response.Code != http.StatusOK {
		t.Errorf("Expected %v, got %v", http.StatusOK, response.Code)
	}
}
