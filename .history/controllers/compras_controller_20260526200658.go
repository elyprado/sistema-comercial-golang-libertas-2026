package controllers

import (
	"encoding/json"
	"go-crud-api/config"
	"go-crud-api/models"
	"net/http"

	"github.com/gorilla/mux"
)

//Controller para a entidade Compras, responsável por lidar com as requisições relacionadas a compras, como criar, ler, atualizar e deletar registros de compras no banco de dados. Cada função se conecta ao banco de dados, executa a operação necessária e retorna a resposta apropriada ao cliente.

// GetCompras é responsável por buscar todas as compras no banco de dados e retornar os resultados em formato JSON. Ele se conecta ao banco de dados, executa uma consulta SQL para selecionar os campos relevantes da tabela de compras, e depois codifica os resultados em JSON para enviar de volta ao cliente.
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

	params := mux.Vars(r)
	id := params["id"]

	rows, err := db.Query("SELECT idcompra, date, idfornecedor, data_vencimento FROM compra WHERE idcompra = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var compra models.Compras
	for rows.Next() {
		if err := rows.Scan(&compra.IdCompra, &compra.Data, &compra.IdFornecedor, &compra.DataVencimento); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(compra)
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
