package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq" // importa o driver para PostgreSQL
	"github.com/marmota-alpina/orders-service/internal/config"
)

// GetDatabase retorna uma conexão com o banco de dados PostgreSQL
func GetDatabase() (*sql.DB, error) {
	cfg := config.LoadConfig()

	connStr := cfg.GetURL()

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("error opening database connection: %v", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("error pinging database: %v", err)
	}

	return db, nil
}
