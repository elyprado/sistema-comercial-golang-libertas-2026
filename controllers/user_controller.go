package controllers

import (
	"encoding/json"
	"go-crud-api/config"
	"go-crud-api/models"
	"go-crud-api/utils"
	"net/http"

	"github.com/gorilla/mux"
)

// Busca todos os usuários e os retorna em JSON
func GetUsuarios(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidarTokenRequest(w, r) {
		return
	}

	db, err := config.Connect()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	rows, err := db.Query("SELECT idusuario, nome, email, senha, telefone FROM usuario ORDER BY nome")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Idusuario, &user.Nome, &user.Email, &user.Senha, &user.Telefone); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(users)
}

// Busca um usuário pelo ID informado na rota
func GetUsuarioById(w http.ResponseWriter, r *http.Request) {
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

	rows, err := db.Query("SELECT idusuario, nome, email, senha, telefone FROM usuario WHERE idusuario = ?", id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var user models.User
	for rows.Next() {
		if err := rows.Scan(&user.Idusuario, &user.Nome, &user.Email, &user.Senha, &user.Telefone); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}

// Insere um novo usuário a partir do JSON recebido
func CreateUsuario(w http.ResponseWriter, r *http.Request) {
	if !utils.ValidarTokenRequest(w, r) {
		return
	}

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

// Atualiza os dados de um usuário existente pelo ID
func UpdateUsuario(w http.ResponseWriter, r *http.Request) {
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

// Remove o usuário identificado pelo ID na rota
func DeleteUsuario(w http.ResponseWriter, r *http.Request) {
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

func Login(w http.ResponseWriter, r *http.Request) {
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

	query := "SELECT idusuario, nome FROM usuario WHERE email = ? AND senha = ?"
	rows, erro := db.Query(query, usuario.Email, usuario.Senha)
	if erro != nil {
		http.Error(w, erro.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	if rows.Next() {
		var id int
		var nome string
		err := rows.Scan(&id, &nome)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		token, err := utils.GerarToken(id, nome, "user")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(map[string]string{"token": token})
		return
	}

	w.WriteHeader(http.StatusUnauthorized)
	json.NewEncoder(w).Encode(map[string]string{"message": "falha ao autenticar"})
}
