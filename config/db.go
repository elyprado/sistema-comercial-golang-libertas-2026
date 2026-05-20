package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	// Carregar variáveis de ambiente do arquivo .env
	if err := godotenv.Load(); err != nil {
		log.Println("Aviso: Não foi possível carregar o arquivo .env")
	}
}

func Connect() (*sql.DB, error) {
	// Obter credenciais do banco de dados das variáveis de ambiente
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbname := os.Getenv("DB_NAME")

	if host == "" || port == "" || user == "" || password == "" || dbname == "" {
		return nil, fmt.Errorf("Erro: Variáveis de ambiente do banco de dados não configuradas corretamente")
	}

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("Erro ao conectar ao banco de dados: %v", err)
	}

	return db, err
}
