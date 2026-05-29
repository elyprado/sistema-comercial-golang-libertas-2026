package controllers

import (
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
	rows, err := db.Query("SELECT idfornecedor, nome, telefone, cnpj, e-mail, logradouro, numero, bairro, cidade, uf FROM fornecedor ORDER BY nome")
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

	rows, err := db.Query("SELECT idfornecedor, nome, telefone, cnpj, e-mail, logradouro, numero, bairro, cidade, uf FROM fornecedor WHERE idfornecedor = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()
	var forn models.Fornecedor
	for rows.Next() {
		if err := rows.Scan(&forn.Idfornecedor, &forn.Nome, &forn.Telefone, &forn.Cnpj, &forn.Email, &forn.Logradouro, &forn.Numero, &forn.Bairro, &forn.Cidade, &forn.Uf); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(forn)
}

//----------------------------------------------------------------------------------------------------------------

func CreateUsuario(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	var usuario models.User
	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	query := "INSERT INTO usuario (nome, email, senha, telefone) VALUES (?, ?, ?, ?)"
	_, erro = db.Exec(query, usuario.Nome, usuario.Email, usuario.Senha,
		usuario.Telefone)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(map[string]string{"message": "Sucesso"})
}

func UpdateUsuario(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()
	params := mux.Vars(r)
	id := params["id"]
	var usuario models.User
	err := json.NewDecoder(r.Body).Decode(&usuario)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	query := "UPDATE usuario SET nome=?, email=?, senha=?, telefone=? WHERE idusuario=?"
	result, erro := db.Exec(query, usuario.Nome, usuario.Email, usuario.Senha,
		usuario.Telefone, id)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Usuário atualizado com sucesso"})
}

func DeleteUsuario(w http.ResponseWriter, r *http.Request) {
	db, erro := config.Connect()
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	params := mux.Vars(r)
	id := params["id"]

	query := "DELETE FROM usuario WHERE idusuario=?"
	result, erro := db.Exec(query, id)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		http.Error(w, "Usuário não encontrado", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(map[string]string{"message": "Usuário removido com sucesso"})
}
