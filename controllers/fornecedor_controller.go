package controllers

import (
	"database/sql"
	"encoding/json"
	"go-crud-api/config"
	"go-crud-api/models"
	"net/http"

	"github.com/gorilla/mux"
)

func GetFornecedores(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT idfornecedor, nome, telefone, cnpj, email, logradouro, numero, bairro, cidade, uf FROM fornecedor ORDER BY nome")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var fornecedores []models.Fornecedor
	for rows.Next() {
		var forn models.Fornecedor
		if err := rows.Scan(&forn.Idfornecedor, &forn.Nome, &forn.Telefone, &forn.Cnpj, &forn.Email, &forn.Logradouro, &forn.Numero, &forn.Bairro, &forn.Cidade, &forn.Uf); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		fornecedores = append(fornecedores, forn)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fornecedores)
}

func GetFornecedorById(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	row := db.QueryRow("SELECT idfornecedor, nome, telefone, cnpj, email, logradouro, numero, bairro, cidade, uf FROM fornecedor WHERE idfornecedor = ?", id)
	var forn models.Fornecedor
	if err := row.Scan(&forn.Idfornecedor, &forn.Nome, &forn.Telefone, &forn.Cnpj, &forn.Email, &forn.Logradouro, &forn.Numero, &forn.Bairro, &forn.Cidade, &forn.Uf); err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Fornecedor não encontrado", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(forn)
}

func CreateFornecedor(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var forn models.Fornecedor
	if err := json.NewDecoder(r.Body).Decode(&forn); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "INSERT INTO fornecedor (nome, telefone, cnpj, email, logradouro, numero, bairro, cidade, uf) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err = db.Exec(query, forn.Nome, forn.Telefone, forn.Cnpj, forn.Email, forn.Logradouro, forn.Numero, forn.Bairro, forn.Cidade, forn.Uf)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Fornecedor criado com sucesso"})
}

func UpdateFornecedor(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	var forn models.Fornecedor
	if err := json.NewDecoder(r.Body).Decode(&forn); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "UPDATE fornecedor SET nome=?, telefone=?, cnpj=?, email=?, logradouro=?, numero=?, bairro=?, cidade=?, uf=? WHERE idfornecedor=?"
	result, err := db.Exec(query, forn.Nome, forn.Telefone, forn.Cnpj, forn.Email, forn.Logradouro, forn.Numero, forn.Bairro, forn.Cidade, forn.Uf, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Fornecedor não encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Fornecedor atualizado com sucesso"})
}

func DeleteFornecedor(w http.ResponseWriter, r *http.Request) {
	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	query := "DELETE FROM fornecedor WHERE idfornecedor=?"
	result, err := db.Exec(query, id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Fornecedor não encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Fornecedor removido com sucesso"})
}
