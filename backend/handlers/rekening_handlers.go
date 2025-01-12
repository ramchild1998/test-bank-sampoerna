package handlers

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
    "backend/models"
    "backend/database"
    "time"
)

func CreateRekening(w http.ResponseWriter, r *http.Request) {
    var rekening models.Rekening
    json.NewDecoder(r.Body).Decode(&rekening)
    rekening.TanggalPembuatan = time.Now()
    database.DB.Create(&rekening)
    json.NewEncoder(w).Encode(rekening)
}

func ReadRekening(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var rekening models.Rekening
    database.DB.First(&rekening, params["id"])
    json.NewEncoder(w).Encode(rekening)
}

func UpdateRekening(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var rekening models.Rekening
    database.DB.First(&rekening, params["id"])
    json.NewDecoder(r.Body).Decode(&rekening)
    database.DB.Save(&rekening)
    json.NewEncoder(w).Encode(rekening)
}

func DeleteRekening(w http.ResponseWriter, r *http.Request) {
    params := mux.Vars(r)
    var rekening models.Rekening
    database.DB.Delete(&rekening, params["id"])
    w.WriteHeader(http.StatusNoContent)
}
