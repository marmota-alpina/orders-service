package graph

import (
	"database/sql"
)

// Resolver struct para injetar a conexão do banco
type Resolver struct {
	DB *sql.DB
}
