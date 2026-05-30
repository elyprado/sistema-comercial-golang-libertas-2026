package controllers

import (
	"encoding/json"
	"go-crud-api/config"
	"go-crud-api/models"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shopspring/decimal"
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

// GetCompraItemById é responsável por buscar um item de compra específico no banco de dados com base no ID fornecido na URL. Ele se conecta ao banco de dados, executa uma consulta SQL para selecionar os campos relevantes da tabela de itens de compra onde o ID corresponde ao valor fornecido, e depois codifica o resultado em JSON para enviar de volta ao cliente.
func GetCompraItemById(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	rows, err := db.Query("SELECT idcompraitem, idcompra, idproduto, quantidade, custo_unitario, custo_total FROM compraItem WHERE idcompraitem = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var compraItem models.CompraItem
	for rows.Next() {
		if err := rows.Scan(&compraItem.IdCompraItem, &compraItem.IdCompra, &compraItem.IdProduto, &compraItem.Quantidade, &compraItem.CustoUnitario, &compraItem.CustoTotal); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(compraItem)
}

// CreateCompraItem é responsável por criar um novo item de compra no banco de dados. Ele se conecta ao banco de dados, decodifica os dados do item de compra enviados no corpo da requisição em formato JSON, valida os campos obrigatórios e depois executa uma consulta SQL para inserir os dados do novo item de compra na tabela de itens de compra. Se a operação for bem-sucedida, ele retorna uma mensagem de sucesso em formato JSON.
func CreateCompraItem(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var compraItem models.CompraItem
	if err := json.NewDecoder(r.Body).Decode(&compraItem); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if compraItem.IdCompra == nil || compraItem.IdProduto == nil || compraItem.Quantidade == nil || *compraItem.Quantidade <= 0 || (*compraItem.CustoUnitario).LessThanOrEqual(decimal.Zero) || (*compraItem.CustoTotal).LessThanOrEqual(decimal.Zero) {
		http.Error(w, "Campos obrigatórios faltando ou inválidos", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO compraItem (idcompra, idproduto, quantidade, custo_unitario, custo_total) VALUES (?, ?, ?, ?, ?)"
	_, err = db.Exec(query, compraItem.IdCompra, compraItem.IdProduto, compraItem.Quantidade, compraItem.CustoUnitario, compraItem.CustoTotal)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Sucesso"})
}

// UpdateCompraItem é responsável por atualizar um item de compra existente no banco de dados com base no ID fornecido na URL. Ele se conecta ao banco de dados, decodifica os dados do item de compra enviados no corpo da requisição em formato JSON, valida os campos obrigatórios e depois executa uma consulta SQL para atualizar os dados do item de compra na tabela de itens de compra onde o ID corresponde ao valor fornecido. Se a operação for bem-sucedida, ele retorna uma mensagem de sucesso em formato JSON.
func UpdateCompraItem(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]
	var compraItem models.CompraItem
	err := json.NewDecoder(r.Body).Decode(&compraItem)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	query := "UPDATE compraItem SET idcompra=?, idproduto=?, quantidade=?, custo_unitario=?, custo_total=? WHERE idcompraitem=?"
	result, err := db.Exec(query, compraItem.IdCompra, compraItem.IdProduto, compraItem.Quantidade, compraItem.CustoUnitario, compraItem.CustoTotal, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Item não encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Item atualizado com sucesso"})
}

// DeleteCompraItem é responsável por deletar um item de compra existente no banco de dados com base no ID fornecido na URL. Ele se conecta ao banco de dados, executa uma consulta SQL para deletar o item de compra da tabela de itens de compra onde o ID corresponde ao valor fornecido. Se a operação for bem-sucedida, ele retorna uma mensagem de sucesso em formato JSON. Se o item de compra não for encontrado, ele retorna um erro 404.
func DeleteCompraItem(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	query := "DELETE FROM compraItem WHERE idcompraitem=?"
	result, err := db.Exec(query, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Item não encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Item removido com sucesso"})
}
