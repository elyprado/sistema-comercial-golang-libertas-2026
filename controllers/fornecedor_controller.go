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

func getFornecedorID(w http.ResponseWriter, r *http.Request) (int, bool) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil || id <= 0 {
		http.Error(w, "ID do fornecedor inválido", http.StatusBadRequest)
		return 0, false
	}
	return id, true
}

func validarFornecedor(forn models.Fornecedor) string {
	if forn.Nome == nil || strings.TrimSpace(*forn.Nome) == "" {
		return "Nome é obrigatório"
	}
	*forn.Nome = strings.TrimSpace(*forn.Nome)

	if forn.Cnpj == nil {
		return "CNPJ inválido"
	}
	*forn.Cnpj = somenteNumerosF(*forn.Cnpj)
	if !cnpjValido(*forn.Cnpj) {
		return "CNPJ inválido"
	}
	return ""
}

func somenteNumerosF(valor string) string {
	var numeros strings.Builder
	for _, char := range valor {
		if char >= '0' && char <= '9' {
			numeros.WriteRune(char)
		}
	}
	return numeros.String()
}

func cnpjValido(cnpj string) bool {
	if len(cnpj) != 14 {
		return false
	}

	todosIguais := true
	for i := 1; i < len(cnpj); i++ {
		if cnpj[i] != cnpj[0] {
			todosIguais = false
			break
		}
	}
	if todosIguais {
		return false
	}

	return digitoCNPJ(cnpj, 12) == int(cnpj[12]-'0') &&
		digitoCNPJ(cnpj, 13) == int(cnpj[13]-'0')
}

func digitoCNPJ(cnpj string, tamanho int) int {
	pesos := []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}
	offset := len(pesos) - tamanho
	soma := 0
	for i := 0; i < tamanho; i++ {
		soma += int(cnpj[i]-'0') * pesos[offset+i]
	}
	resto := soma % 11
	if resto < 2 {
		return 0
	}
	return 11 - resto
}

func GetFornecedores(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidarTokenRequest(w, r) {
		return
	}

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
	if err := rows.Err(); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(fornecedores)
}

func GetFornecedorById(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidarTokenRequest(w, r) {
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	id, ok := getFornecedorID(w, r)
	if !ok {
		return
	}

	var forn models.Fornecedor
	err = db.QueryRow("SELECT idfornecedor, nome, telefone, cnpj, email, logradouro, numero, bairro, cidade, uf FROM fornecedor WHERE idfornecedor = ?", id).
		Scan(&forn.Idfornecedor, &forn.Nome, &forn.Telefone, &forn.Cnpj, &forn.Email, &forn.Logradouro, &forn.Numero, &forn.Bairro, &forn.Cidade, &forn.Uf)
	if err == sql.ErrNoRows {
		http.Error(w, "Fornecedor não encontrado", http.StatusNotFound)
		return
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(forn)
}

func CreateFornecedor(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidarTokenRequest(w, r) {
		return
	}

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

	if msg := validarFornecedor(forn); msg != "" {
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	query := "INSERT INTO fornecedor (nome, telefone, cnpj, email, logradouro, numero, bairro, cidade, uf) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)"
	_, err = db.Exec(query, forn.Nome, forn.Telefone, forn.Cnpj, forn.Email, forn.Logradouro, forn.Numero, forn.Bairro, forn.Cidade, forn.Uf)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Fornecedor criado com sucesso"})
}

func UpdateFornecedor(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidarTokenRequest(w, r) {
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	id, ok := getFornecedorID(w, r)
	if !ok {
		return
	}

	var forn models.Fornecedor
	if err := json.NewDecoder(r.Body).Decode(&forn); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if msg := validarFornecedor(forn); msg != "" {
		http.Error(w, msg, http.StatusBadRequest)
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Fornecedor atualizado com sucesso"})
}

func DeleteFornecedor(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidarTokenRequest(w, r) {
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	id, ok := getFornecedorID(w, r)
	if !ok {
		return
	}

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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "Fornecedor removido com sucesso"})
}