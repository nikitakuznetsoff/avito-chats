package postgres

import (
	"github.com/jackc/pgx/v4/pgxpool"
)

type PostgresRepo struct {
	conn *pgxpool.Pool
}

func CreatePostgresRepo(db *pgxpool.Pool) *PostgresRepo {
	return &PostgresRepo{conn: db}
}