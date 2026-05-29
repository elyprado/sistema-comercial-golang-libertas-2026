package controllers

import (
	"encoding/json"
	"go-crud-api/config"
	"go-crud-api/models"
	"net/http"
)

// Controller para a entidade CompraItem, responsável por lidar com as requisições relacionadas a itens de compra, como criar, ler, atualizar e deletar registros de itens de compra no banco de dados. Cada função se conecta ao banco de dados, executa a operação necessária e retorna a resposta apropriada ao cliente.

// GetCompraItem é responsável por buscar todos os itens de compra no banco de dados e retornar os resultados em formato JSON. Ele se conecta ao banco de dados, executa uma consulta SQL para selecionar os campos relevantes da tabela de itens de compra, e depois codifica os resultados em JSON para enviar de volta ao cliente.
func GetCompraItem(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT idcompraitem, idcompra, idproduto, quantidade, custo_unitario, custo_total FROM compraItem")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var compraItems []models.CompraItem
	for rows.Next() {
		var compraItem models.CompraItem
		if err := rows.Scan(&compraItem.IdCompraItem, &compraItem.IdCompra, &compraItem.IdProduto, &compraItem.Quantidade, &compraItem.CustoUnitario, &compraItem.CustoTotal); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		compraItems = append(compraItems, compraItem)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(compraItems)
}

func GetCompraItemById(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
}

func CreateCompraItem(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
}

func UpdateCompraItem(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
}

func DeleteCompraItem(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
}
