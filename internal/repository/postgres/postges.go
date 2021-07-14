package postgres

import (
	"github.com/jackc/pgx/v4"
)

type PostgresRepo struct {
	conn *pgx.Conn 
}

func CreatePostgresRepo(db *pgx.Conn) *PostgresRepo {
	return &PostgresRepo{conn: db}
}