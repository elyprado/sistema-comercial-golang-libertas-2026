package controllers

import (
	"database/sql"
	"encoding/json"
	"go-crud-api/config"
	"go-crud-api/models"
	"go-crud-api/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

func getClienteID(w http.ResponseWriter, r *http.Request) (int, bool) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id <= 0 {
		http.Error(w, "ID do cliente inválido", http.StatusBadRequest)
		return 0, false
	}
	return id, true
}

func validarCliente(cliente models.Cliente) string {
	if cliente.Nome == nil || strings.TrimSpace(*cliente.Nome) == "" {
		return "Nome é obrigatório"
	}
	*cliente.Nome = strings.TrimSpace(*cliente.Nome)

	if cliente.Cpf == nil {
		return "CPF inválido"
	}
	*cliente.Cpf = somenteNumeros(*cliente.Cpf)
	if !cpfValido(*cliente.Cpf) {
		return "CPF inválido"
	}
	return ""
}

func somenteNumeros(valor string) string {
	var numeros strings.Builder
	for _, char := range valor {
		if char >= '0' && char <= '9' {
			numeros.WriteRune(char)
		}
	}
	return numeros.String()
}

func cpfValido(cpf string) bool {
	if len(cpf) != 11 {
		return false
	}

	todosIguais := true
	for i := 1; i < len(cpf); i++ {
		if cpf[i] != cpf[0] {
			todosIguais = false
			break
		}
	}
	if todosIguais {
		return false
	}

	return digitoCPF(cpf, 9) == int(cpf[9]-'0') && digitoCPF(cpf, 10) == int(cpf[10]-'0')
}

func digitoCPF(cpf string, tamanho int) int {
	soma := 0
	for i := 0; i < tamanho; i++ {
		soma += int(cpf[i]-'0') * (tamanho + 1 - i)
	}
	resto := (soma * 10) % 11
	if resto == 10 {
		return 0
	}
	return resto
}

func GetClientes(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidarTokenRequest(w, r) {
		return
	}

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
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clientes)
}

func GetClienteById(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidarTokenRequest(w, r) {
		return
	}

	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	id, ok := getClienteID(w, r)
	if !ok {
		return
	}

	var cliente models.Cliente
	err := db.QueryRow("SELECT idcliente, nome, telefone, cpf, email, logradouro, numero, bairro, cidade, uf FROM cliente WHERE idcliente = ?", id).
		Scan(&cliente.Idcliente, &cliente.Nome, &cliente.Telefone, &cliente.Cpf, &cliente.Email, &cliente.Logradouro, &cliente.Numero, &cliente.Bairro, &cliente.Cidade, &cliente.Uf)
	if err == sql.ErrNoRows {
		http.Error(w, "Cliente não encontrado", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(cliente)
}

func CreateCliente(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidarTokenRequest(w, r) {
		return
	}

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
	if msg := validarCliente(cliente); msg != "" {
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	query := "INSERT INTO cliente (nome, telefone, cpf, email, logradouro, numero, bairro, cidade, uf) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, erro = db.Exec(query, cliente.Nome, cliente.Telefone, cliente.Cpf, cliente.Email, cliente.Logradouro, cliente.Numero, cliente.Bairro, cliente.Cidade, cliente.Uf)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Sucesso"})
}

func UpdateCliente(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidarTokenRequest(w, r) {
		return
	}

	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	id, ok := getClienteID(w, r)
	if !ok {
		return
	}

	var cliente models.Cliente
	err := json.NewDecoder(r.Body).Decode(&cliente)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if msg := validarCliente(cliente); msg != "" {
		http.Error(w, msg, http.StatusBadRequest)
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Cliente atualizado com sucesso"})
}

func DeleteCliente(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidarTokenRequest(w, r) {
		return
	}

	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	id, ok := getClienteID(w, r)
	if !ok {
		return
	}

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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Cliente removido com sucesso"})
}
