package controllers

import (
	"encoding/json"
	"go-crud-api/config"
	"go-crud-api/models"
	"net/http"

	"time"

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

// GetComprasById é responsável por buscar uma compra específica no banco de dados com base no ID fornecido na URL. Ele se conecta ao banco de dados, executa uma consulta SQL para selecionar os campos relevantes da tabela de compras onde o ID corresponde ao valor fornecido, e depois codifica o resultado em JSON para enviar de volta ao cliente.
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

// CreateCompras é responsável por criar uma nova compra no banco de dados. Ele se conecta ao banco de dados, decodifica os dados da compra enviados no corpo da requisição em formato JSON, valida os campos obrigatórios e as datas, e depois executa uma consulta SQL para inserir os dados da nova compra na tabela de compras. Se a operação for bem-sucedida, ele retorna uma mensagem de sucesso em formato JSON.
func CreateCompras(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var compra models.Compras
	if err := json.NewDecoder(r.Body).Decode(&compra); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// verifica se variaveis não são nulas
	if compra.Data == nil || compra.IdFornecedor == nil || compra.DataVencimento == nil {
		http.Error(w, "Campos 'data', 'idfornecedor' e 'data_vencimento' são obrigatórios", http.StatusBadRequest)
		return
	}

	// Converter as strings de data para time
	dataConvertida, err := time.Parse("2006-01-02", *compra.Data)
	if err != nil {
		http.Error(w, "Formato de 'date' inválido. Use o formato YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	vencimentoConvertido, err := time.Parse("2006-01-02", *compra.DataVencimento)
	if err != nil {
		http.Error(w, "Formato de 'data_vencimento' inválido. Use o formato YYYY-MM-DD", http.StatusBadRequest)
		return
	}

	query := "INSERT INTO compra (date, idfornecedor, data_vencimento) VALUES (?, ?, ?)"
	_, err = db.Exec(query, dataConvertida, compra.IdFornecedor, vencimentoConvertido)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Sucesso"})
}

// UpdateCompras é responsável por atualizar uma compra existente no banco de dados com base no ID fornecido na URL. Ele se conecta ao banco de dados, decodifica os dados da compra enviados no corpo da requisição em formato JSON, valida os campos obrigatórios e as datas, e depois executa uma consulta SQL para atualizar os dados da compra na tabela de compras onde o ID corresponde ao valor fornecido. Se a operação for bem-sucedida, ele retorna uma mensagem de sucesso em formato JSON.
func UpdateCompras(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	params := mux.Vars(r)
	id := params["id"]
	var compra models.Compras
	err := json.NewDecoder(r.Body).Decode(&compra)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	query := "UPDATE compra SET date=?, idfornecedor=?, data_vencimento=? WHERE idcompra=?"
	result, err := db.Exec(query, compra.Data, compra.IdFornecedor, compra.DataVencimento, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Compra não encontrada", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Compra atualizada com sucesso"})
}

// DeleteCompras é responsável por deletar uma compra existente no banco de dados com base no ID fornecido na URL. Ele se conecta ao banco de dados, executa uma consulta SQL para deletar a compra da tabela de compras onde o ID corresponde ao valor fornecido. Se a operação for bem-sucedida, ele retorna uma mensagem de sucesso em formato JSON. Se a compra não for encontrada, ele retorna um erro 404.
func DeleteCompras(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	query := "DELETE FROM compra WHERE idcompra=?"
	result, err := db.Exec(query, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Compra não encontrada", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Compra removida com sucesso"})
}

// ----------------- Consulta de fornecedores para compras -----------------

// GetFornecedoresDisponiveis retorna fornecedores disponíveis para serem associados a uma compra. Ele se conecta ao banco de dados, executa uma consulta SQL para selecionar os campos relevantes da tabela de fornecedores, e depois codifica os resultados em JSON para enviar de volta ao cliente. Essa função é útil para preencher dropdowns ou listas de seleção no frontend quando o usuário estiver criando ou editando uma compra, permitindo que ele escolha um fornecedor existente.

func GetFornecedoresDisponiveis(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT idfornecedor, nome FROM fornecedor")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	fornecedores := []models.Fornecedor{}
	for rows.Next() {
		var fornecedor models.Fornecedor
		err := rows.Scan(&fornecedor.IdFornecedor, &fornecedor.Nome)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fornecedores = append(fornecedores, fornecedor)
	}

	json.NewEncoder(w).Encode(fornecedores)
}

// GetProdutosDisponiveis retorna produtos disponíveis para serem associados a uma compra. Ele se conecta ao banco de dados, executa uma consulta SQL para selecionar os campos relevantes da tabela de produtos, e depois codifica os resultados em JSON para enviar de volta ao cliente. Essa função é útil para preencher dropdowns ou listas de seleção no frontend quando o usuário estiver criando ou editando uma compra, permitindo que ele escolha um produto existente.

func GetProdutosDisponiveis(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT idproduto, descricao FROM produto")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	produtos := []models.Produto{}
	for rows.Next() {
		var produto models.Produto
		err := rows.Scan(&produto.IdProduto, &produto.Descricao)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		produtos = append(produtos, produto)
	}
	json.NewEncoder(w).Encode(produtos)
}
