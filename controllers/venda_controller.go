package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"go-crud-api/config"
	"go-crud-api/models"
	"go-crud-api/utils"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

type vendaPayload struct {
	Data           string             `json:"data"`
	Idcliente      int                `json:"idcliente"`
	Idvendedor     int                `json:"idvendedor"`
	DataVencimento string             `json:"data_vencimento"`
	Items          []models.VendaItem `json:"items"`
}

func parseDateTime(value string) (time.Time, error) {
	layouts := []string{
		"2006-01-02",
		"2006-01-02T15:04",
		"2006-01-02T15:04:05",
		"2006-01-02 15:04:05",
		"2006-01-02T15:04:05Z",
	}
	for _, layout := range layouts {
		if t, err := time.Parse(layout, value); err == nil {
			return t, nil
		}
	}
	return time.Time{}, fmt.Errorf("formato de data inválido: %s", value)
}

func CreateVenda(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidarTokenRequest(w, r) {
		return
	}
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var payload vendaPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if payload.Idcliente == 0 || payload.Idvendedor == 0 || payload.Data == "" || payload.DataVencimento == "" {
		http.Error(w, "Dados incompletos da venda", http.StatusBadRequest)
		return
	}

	if len(payload.Items) == 0 {
		http.Error(w, "A venda deve conter pelo menos um item", http.StatusBadRequest)
		return
	}

	dataVenda, err := parseDateTime(payload.Data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dataVencimento, err := parseDateTime(payload.DataVencimento)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	tx, err := db.Begin()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer func() {
		if p := recover(); p != nil {
			tx.Rollback()
			http.Error(w, fmt.Sprintf("Erro interno"), http.StatusInternalServerError)
		}
	}()

	insertVenda := "INSERT INTO venda (data, idcliente, idvendedor, data_vencimento) VALUES (?, ?, ?, ?)"
	result, err := tx.Exec(insertVenda, dataVenda, payload.Idcliente, payload.Idvendedor, dataVencimento)
	if err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	vendaID, err := result.LastInsertId()
	if err != nil {
		tx.Rollback()
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	insertItem := "INSERT INTO vendaitem (idvenda, idproduto, quantidade, valor_unitario, valor_total) VALUES (?, ?, ?, ?, ?)"
	for index, item := range payload.Items {
		if item.Idproduto == 0 || item.Quantidade == 0 || item.ValorUnitario == 0 {
			tx.Rollback()
			http.Error(w, fmt.Sprintf("Item %d inválido: produto, quantidade e valor unitário são obrigatórios", index+1), http.StatusBadRequest)
			return
		}
		valorTotal := float64(item.Quantidade) * item.ValorUnitario
		if _, err := tx.Exec(insertItem, vendaID, item.Idproduto, item.Quantidade, item.ValorUnitario, valorTotal); err != nil {
			tx.Rollback()
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	if err := tx.Commit(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]any{"message": "Venda registrada com sucesso", "idvenda": vendaID})
}

func GetVendas(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidarTokenRequest(w, r) {
		return
	}
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	filters := []string{}
	values := []interface{}{}

	queryParams := r.URL.Query()
	if dataInicial := strings.TrimSpace(queryParams.Get("data_inicial")); dataInicial != "" {
		if data, err := parseDateTime(dataInicial); err == nil {
			filters = append(filters, "v.data >= ?")
			values = append(values, data)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	if dataFinal := strings.TrimSpace(queryParams.Get("data_final")); dataFinal != "" {
		if data, err := parseDateTime(dataFinal); err == nil {
			filters = append(filters, "v.data <= ?")
			values = append(values, data)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}

	query := "SELECT v.idvenda, v.data, v.idcliente, v.idvendedor, v.data_vencimento, IFNULL(SUM(vi.valor_total), 0) AS total, COUNT(vi.idvendaitem) AS item_count FROM venda v LEFT JOIN vendaitem vi ON vi.idvenda = v.idvenda"
	if len(filters) > 0 {
		query += " WHERE " + strings.Join(filters, " AND ")
	}
	query += " GROUP BY v.idvenda ORDER BY v.data DESC"

	rows, err := db.Query(query, values...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	vendas := make([]models.Venda, 0)
	for rows.Next() {
		var venda models.Venda
		var dataVenda time.Time
		var dataVencimento time.Time
		if err := rows.Scan(&venda.Idvenda, &dataVenda, &venda.Idcliente, &venda.Idvendedor, &dataVencimento, &venda.Total, &venda.ItemCount); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		venda.Data = dataVenda
		venda.DataVencimento = dataVencimento
		vendas = append(vendas, venda)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vendas)
}

func GetVendaById(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidarTokenRequest(w, r) {
		return
	}
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	if err != nil {
		http.Error(w, "ID inválido", http.StatusBadRequest)
		return
	}

	row := db.QueryRow("SELECT idvenda, data, idcliente, idvendedor, data_vencimento FROM venda WHERE idvenda = ?", id)
	var venda models.Venda
	var dataVenda time.Time
	var dataVencimento time.Time
	if err := row.Scan(&venda.Idvenda, &dataVenda, &venda.Idcliente, &venda.Idvendedor, &dataVencimento); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Venda não encontrada", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	venda.Data = dataVenda
	venda.DataVencimento = dataVencimento

	itemRows, err := db.Query("SELECT idvendaitem, idvenda, idproduto, quantidade, valor_unitario, valor_total FROM vendaitem WHERE idvenda = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer itemRows.Close()

	for itemRows.Next() {
		var item models.VendaItem
		if err := itemRows.Scan(&item.Idvendaitem, &item.Idvenda, &item.Idproduto, &item.Quantidade, &item.ValorUnitario, &item.ValorTotal); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		venda.Items = append(venda.Items, item)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(venda)
}
