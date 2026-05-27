package controllers

import (
	"encoding/json"
	"go-crud-api/config"
	"go-crud-api/models"
	"net/http"

	"github.com/gorilla/mux"
)

func GetProdutos(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT idproduto, descricao, marca, preco_de_custo, preco_de_venda, codigo_referencia, codigo_barras FROM produto ORDER BY descricao")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var produtos []models.Produto

	for rows.Next() {
		var p models.Produto
		if err := rows.Scan(
			&p.IdProduto,
			&p.Descricao,
			&p.Marca,
			&p.PrecoDeCusto,
			&p.PrecoDeVenda,
			&p.CodigoReferencia,
			&p.CodigoBarras,
		); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		produtos = append(produtos, p)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(produtos)
}

func GetProdutoById(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	rows, err := db.Query("SELECT idproduto, descricao, marca, preco_de_custo, preco_de_venda, codigo_referencia, codigo_barras FROM produto WHERE idproduto = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var p models.Produto

	for rows.Next() {
		if err := rows.Scan(
			&p.IdProduto,
			&p.Descricao,
			&p.Marca,
			&p.PrecoDeCusto,
			&p.PrecoDeVenda,
			&p.CodigoReferencia,
			&p.CodigoBarras,
		); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(p)
}

func CreateProduto(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var produto models.Produto
	err := json.NewDecoder(r.Body).Decode(&produto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "INSERT INTO produto (descricao, marca, preco_de_custo, preco_de_venda, codigo_referencia, codigo_barras) VALUES (?, ?, ?, ?, ?, ?)"
	_, erro = db.Exec(query,
		produto.Descricao,
		produto.Marca,
		produto.PrecoDeCusto,
		produto.PrecoDeVenda,
		produto.CodigoReferencia,
		produto.CodigoBarras,
	)

	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Produto criado com sucesso"})
}

func UpdateProduto(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	var produto models.Produto
	err := json.NewDecoder(r.Body).Decode(&produto)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "UPDATE produto SET descricao=?, marca=?, preco_de_custo=?, preco_de_venda=?, codigo_referencia=?, codigo_barras=? WHERE idproduto=?"
	result, erro := db.Exec(query,
		produto.Descricao,
		produto.Marca,
		produto.PrecoDeCusto,
		produto.PrecoDeVenda,
		produto.CodigoReferencia,
		produto.CodigoBarras,
		id,
	)

	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Produto não encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Produto atualizado"})
}

func DeleteProduto(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	query := "DELETE FROM produto WHERE idproduto=?"
	result, erro := db.Exec(query, id)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Produto não encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Produto removido"})
}