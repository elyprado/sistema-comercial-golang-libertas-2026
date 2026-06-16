package controllers

import (
	"encoding/json"
	"net/http"

	"go-crud-api/config"
	"go-crud-api/models"

	"github.com/gorilla/mux"
)

func CreateContaPagar(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var contasPagar models.ContasPagarRequest

	err = json.NewDecoder(r.Body).Decode(&contasPagar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "INSERT INTO contas_pagar (id_fornecedor, data, data_vencimento, valor)  VALUES (?, ?,?,?)"

	_, err = db.Exec(query, contasPagar.IDFornecedor, contasPagar.Data, contasPagar.DataVencimento, contasPagar.Valor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Sucesso"})
}

func GetContasPagar(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT id_contas_pagar,id_fornecedor ,data, data_vencimento ,valor ,data_pagamento  FROM contas_pagar")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var contasPagar []models.ContasPagarResponse
	for rows.Next() {
		var contaPagar models.ContasPagarResponse

		if err := rows.Scan(&contaPagar.IDContasPagar, &contaPagar.IDFornecedor, &contaPagar.Data, &contaPagar.DataVencimento, &contaPagar.Valor, &contaPagar.DataPagamento); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		contasPagar = append(contasPagar, contaPagar)

	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contasPagar)
}

func GetContasPagarById(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	rows, err := db.Query("SELECT id_contas_pagar,id_fornecedor ,data, data_vencimento ,valor ,data_pagamento  FROM contas_pagar WHERE id_contas_pagar = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var contasPagar []models.ContasPagarResponse
	for rows.Next() {
		var contaPagar models.ContasPagarResponse

		if err := rows.Scan(&contaPagar.IDContasPagar, &contaPagar.IDFornecedor, &contaPagar.Data, &contaPagar.DataVencimento, &contaPagar.Valor, &contaPagar.DataPagamento); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		contasPagar = append(contasPagar, contaPagar)

	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contasPagar)
}

func UpdateContaPagar(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	var contasPagar models.ContasPagarRequest

	err = json.NewDecoder(r.Body).Decode(&contasPagar)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "UPDATE contas_pagar SET id_fornecedor=?, data=?, data_vencimento=?, valor=? WHERE id_contas_pagar=?"

	result, err := db.Exec(query, contasPagar.IDFornecedor, contasPagar.Data, contasPagar.DataVencimento, contasPagar.Valor, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Conta a pagar não encontrada", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Conta a pagar atualizada com sucesso"})
}

func DeleteContaPagar(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	query := "DELETE FROM contas_pagar WHERE id_contas_pagar=?"

	result, err := db.Exec(query, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Conta a pagar não encontrada", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Conta a pagar removida com sucesso"})
}