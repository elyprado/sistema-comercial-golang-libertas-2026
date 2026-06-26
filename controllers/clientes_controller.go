package controllers

import (
	"encoding/json"
	"go-crud-api/config"
	"go-crud-api/models"
	"net/http"

	"github.com/gorilla/mux"
)

func GetClientes(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT idcliente, nome, telefone, cpf, email, logradouro, numero, bairro, cidade, uf FROM cliente ORDER BY nome")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var clientes []models.Cliente
	for rows.Next() {
		var cliente models.Cliente
		if err := rows.Scan(&cliente.Idcliente, &cliente.Nome, &cliente.Telefone, &cliente.Cpf, &cliente.Email, &cliente.Logradouro, &cliente.Numero, &cliente.Bairro, &cliente.Cidade, &cliente.Uf); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		clientes = append(clientes, cliente)
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clientes)
}

func GetClienteById(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	rows, err := db.Query("SELECT idcliente, nome, telefone, cpf, email, logradouro, numero, bairro, cidade, uf FROM cliente WHERE idcliente = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var cliente models.Cliente
	for rows.Next() {
		if err := rows.Scan(&cliente.Idcliente, &cliente.Nome, &cliente.Telefone, &cliente.Cpf, &cliente.Email, &cliente.Logradouro, &cliente.Numero, &cliente.Bairro, &cliente.Cidade, &cliente.Uf); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cliente)
}

func CreateCliente(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var cliente models.Cliente
	err := json.NewDecoder(r.Body).Decode(&cliente)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "INSERT INTO cliente (nome, telefone, cpf, email, logradouro, numero, bairro, cidade, uf) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, erro = db.Exec(query, cliente.Nome, cliente.Telefone, cliente.Cpf, cliente.Email, cliente.Logradouro, cliente.Numero, cliente.Bairro, cliente.Cidade, cliente.Uf)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Sucesso"})
}

func UpdateCliente(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	params := mux.Vars(r)
	id := params["id"]
	var cliente models.Cliente
	err := json.NewDecoder(r.Body).Decode(&cliente)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	query := "UPDATE cliente SET nome=?, telefone=?, cpf=?, email=?, logradouro=?, numero=?, bairro=?, cidade=?, uf=? WHERE idcliente=?"
	result, erro := db.Exec(query, cliente.Nome, cliente.Telefone, cliente.Cpf, cliente.Email, cliente.Logradouro, cliente.Numero, cliente.Bairro, cliente.Cidade, cliente.Uf, id)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Cliente não encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Cliente atualizado com sucesso"})
}

func DeleteCliente(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	query := "DELETE FROM cliente WHERE idcliente=?"
	result, erro := db.Exec(query, id)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Cliente não encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Cliente removido com sucesso"})
}
