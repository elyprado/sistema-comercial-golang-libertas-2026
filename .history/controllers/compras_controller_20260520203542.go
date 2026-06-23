package controllers

import (
	"encoding/json"
	"go-crud-api/config"
	"go-crud-api/models"
	"net/http"
)

func GetCompras(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT idcompra, date, idfornecedor, data_vencimento FROM compra")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var compras []models.Compras
	for rows.Next() {
		var compra models.Compras
		if err := rows.Scan(&compra.IdCompra, &compra.Data, &compra.IdFornecedor, &compra.DataVencimento); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		compras = append(compras, compra)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(compras)
}

func GetComprasById(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
}

func CreateCompras(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
}

func UpdateCompras(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
}

func DeleteCompras(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
}
