package utils

import (
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var SecretKey = []byte("jwt-aula-seguranca-2026")

func GerarToken(id int, login string, role string) (string, error) {
	claims := jwt.MapClaims{
		"id":    id,
		"login": login,
		"role":  role,
		"exp":   time.Now().Add(2 * time.Hour).Unix(),
		"iat":   time.Now().Unix(),
	}
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		claims,
	)
	return token.SignedString(SecretKey)
}
func ValidarToken(tokenString string) (*jwt.Token, error) {
	return jwt.Parse(
		tokenString,
		func(token *jwt.Token) (interface{}, error) {
			return SecretKey, nil
		},
	)
}
func ValidarTokenRequest(w http.ResponseWriter, r *http.Request) bool {

	// valida token JWT
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(w, "Token não fornecido", http.StatusUnauthorized)
		return false
	}

	if !strings.HasPrefix(tokenString, "Bearer ") {
		http.Error(w, "Authorization header inválido", http.StatusUnauthorized)
		return false
	}

	// remove a palavra Bearer do token
	tokenString = tokenString[len("Bearer "):]
	token, err := ValidarToken(tokenString)
	if err != nil || !token.Valid {
		http.Error(w, "Token inválido", http.StatusUnauthorized)
		return false
	}

	return true
}
