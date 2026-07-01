package controllers

import (
	"database/sql"
	"encoding/json"
	"go-crud-api/config"
	"go-crud-api/models"
	"go-crud-api/utils"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func GetContasAReceber(w http.ResponseWriter, r *http.Request) {

	if !utils.ValidarTokenRequest(w, r) {
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	dataVencimento := r.URL.Query().Get("data_vencimento")

	query := `
		SELECT
			car.idcontaareceber,
			car.idcliente,
			COALESCE(c.nome, '') AS nome_cliente,
			DATE_FORMAT(car.data, '%Y-%m-%d') AS data,
			DATE_FORMAT(car.data_vencimento, '%Y-%m-%d') AS data_vencimento,
			car.valor,
			DATE_FORMAT(car.data_pagamento, '%Y-%m-%d') AS data_pagamento
		FROM contaareceber car
		LEFT JOIN cliente c ON c.idcliente = car.idcliente
	`

	var rows *sql.Rows
	if dataVencimento != "" {
		query += " WHERE car.data_vencimento = ?"
		query += " ORDER BY car.data_vencimento, car.idcontaareceber DESC"
		rows, err = db.Query(query, dataVencimento)
	} else {
		query += " ORDER BY car.data_vencimento, car.idcontaareceber DESC"
		rows, err = db.Query(query)
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var contas []models.ContaAReceber
	for rows.Next() {
		var conta models.ContaAReceber
		var dataPagamento sql.NullString

		if err := rows.Scan(
			&conta.Idcontaareceber,
			&conta.Idcliente,
			&conta.NomeCliente,
			&conta.Data,
			&conta.DataVencimento,
			&conta.Valor,
			&dataPagamento,
		); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		if dataPagamento.Valid {
			conta.DataPagamento = &dataPagamento.String
		}

		contas = append(contas, conta)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(contas)
}

func CreateContaAReceber(w http.ResponseWriter, r *http.Request) {

	if !utils.ValidarTokenRequest(w, r) {
		return
	}

	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var conta models.ContaAReceber
	err := json.NewDecoder(r.Body).Decode(&conta)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if conta.Idcliente == nil || conta.Data == nil || conta.DataVencimento == nil || conta.Valor == nil {
		http.Error(w, "Informe idcliente, data, data_vencimento e valor", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO contaareceber (idcliente, data, data_vencimento, valor) VALUES (?, ?, ?, ?)"
	result, erro := db.Exec(query, conta.Idcliente, conta.Data, conta.DataVencimento, conta.Valor)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}

	id, _ := result.LastInsertId()
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]interface{}{
		"message": "Conta a receber cadastrada com sucesso",
		"id":      id,
	})
}

func BaixarContaAReceber(w http.ResponseWriter, r *http.Request) {

	if !utils.ValidarTokenRequest(w, r) {
		return
	}

	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]
	dataPagamento := time.Now().Format("2006-01-02")

	query := "UPDATE contaareceber SET data_pagamento = ? WHERE idcontaareceber = ?"
	result, erro := db.Exec(query, dataPagamento, id)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Conta a receber nao encontrada", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"message":        "Conta a receber baixada com sucesso",
		"data_pagamento": dataPagamento,
	})
}