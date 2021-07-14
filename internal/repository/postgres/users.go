package postgres

import (
	"context"
	"github.com/nikitakuznetsoff/avito-chats/internal/models"
)

func (repo *PostgresRepo) CreateUser(username string) (int, error) {
	var id int
	err := repo.conn.
		QueryRow(context.Background(),
		"INSERT INTO users (username) VALUES $1 RETURNING id", username).
		Scan(&id)
	if err != nil {
		return -1, err
	}
	return id, nil
}

func (repo *PostgresRepo) GetUserByID(id int) (*models.User, error) {
	user := &models.User{}
	err := repo.conn.
		QueryRow(context.Background(),
			"SELECT id, username, created_at FROM users WHERE id = ?", id).
		Scan(&user.ID, &user.Username, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return user, nil
}
