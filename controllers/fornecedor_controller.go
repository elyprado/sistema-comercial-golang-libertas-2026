package controllers

import (
	"encoding/json"
	"go-crud-api/config"
	"go-crud-api/models"
	"net/http"
)

func GetFornecedores(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT idfornecedor, nome FROM fornecedor ORDER BY nome")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var fornecedores []models.Fornecedor
	for rows.Next() {
		var fornecedor models.Fornecedor
		if err := rows.Scan(&fornecedor.IdFornecedor, &fornecedor.Nome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fornecedores = append(fornecedores, fornecedor)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fornecedores)
}
