package controllers

import (
	"encoding/json"
	"go-crud-api/config"
	"go-crud-api/models"
	"go-crud-api/utils" 
	"net/http"

	"github.com/gorilla/mux"
)

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

	rows, err := db.Query("SELECT idvendedor, nome, telefone, cpf, email, comissao_percentual, salario FROM vendedor ORDER BY nome")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var vendedores []models.Vendedor
	for rows.Next() {
		var vendedor models.Vendedor
		if err := rows.Scan(&vendedor.Idvendedor, &vendedor.Nome, &vendedor.Telefone, &vendedor.Cpf, &vendedor.Email, &vendedor.ComissaoPercentual, &vendedor.Salario); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		vendedores = append(vendedores, vendedor)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vendedores)
}

func GetVendedorById(w http.ResponseWriter, r *http.Request) {
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
	id := params["id"]

	rows, err := db.Query("SELECT idvendedor, nome, telefone, cpf, email, comissao_percentual, salario FROM vendedor WHERE idvendedor = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var vendedor models.Vendedor
	for rows.Next() {
		if err := rows.Scan(&vendedor.Idvendedor, &vendedor.Nome, &vendedor.Telefone, &vendedor.Cpf, &vendedor.Email, &vendedor.ComissaoPercentual, &vendedor.Salario); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(vendedor)
}

func CreateVendedor(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidarTokenRequest(w, r) {
		return
	}

	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var vendedor models.Vendedor
	err := json.NewDecoder(r.Body).Decode(&vendedor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "INSERT INTO vendedor (nome, telefone, cpf, email, comissao_percentual, salario) VALUES (?, ?, ?, ?, ?, ?)"
	_, erro = db.Exec(query, vendedor.Nome, vendedor.Telefone, vendedor.Cpf, vendedor.Email, vendedor.ComissaoPercentual, vendedor.Salario)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Sucesso"})
}

func UpdateVendedor(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidarTokenRequest(w, r) {
		return
	}

	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	var vendedor models.Vendedor
	err := json.NewDecoder(r.Body).Decode(&vendedor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "UPDATE vendedor SET nome=?, telefone=?, cpf=?, email=?, comissao_percentual=?, salario=? WHERE idvendedor=?"
	result, erro := db.Exec(query, vendedor.Nome, vendedor.Telefone, vendedor.Cpf, vendedor.Email, vendedor.ComissaoPercentual, vendedor.Salario, id)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Vendedor não encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Vendedor atualizado com sucesso"})
}

func DeleteVendedor(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidarTokenRequest(w, r) {
		return
	}

	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	query := "DELETE FROM vendedor WHERE idvendedor=?"
	result, erro := db.Exec(query, id)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Vendedor não encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Vendedor removido com sucesso"})
}