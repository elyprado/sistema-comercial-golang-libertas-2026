package controllers

import (
	"encoding/json"
	"go-crud-api/config"
	"go-crud-api/models"
	"go-crud-api/utils"
	"net/http"
)

// Retorna a lista de clientes disponíveis para seleção em vendas
func GetClientes(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidarTokenRequest(w, r) {
		return
	}
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT idcliente, nome FROM cliente ORDER BY nome")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	clientes := make([]models.Client, 0)
	for rows.Next() {
		var cliente models.Client
		if err := rows.Scan(&cliente.Idcliente, &cliente.Nome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		clientes = append(clientes, cliente)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clientes)
}

// Retorna a lista de vendedores para seleção na criação de venda
func GetVendedores(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidarTokenRequest(w, r) {
		return
	}
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT idvendedor, nome FROM vendedor ORDER BY nome")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	vendedores := make([]models.Vendor, 0)
	for rows.Next() {
		var vendedor models.Vendor
		if err := rows.Scan(&vendedor.Idvendedor, &vendedor.Nome); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		vendedores = append(vendedores, vendedor)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vendedores)
}

// Retorna produtos com preço de venda para composição de itens de venda
func GetProdutos(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidarTokenRequest(w, r) {
		return
	}
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT idproduto, descricao, preco_de_venda FROM produto ORDER BY descricao")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	produtos := make([]models.Product, 0)
	for rows.Next() {
		var produto models.Product
		if err := rows.Scan(&produto.Idproduto, &produto.Descricao, &produto.PrecoVenda); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		produtos = append(produtos, produto)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produtos)
}
