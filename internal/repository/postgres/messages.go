package postgres

import (
	"context"
	"log"

	"github.com/nikitakuznetsoff/avito-chats/internal/models"
)

func (repo *PostgresRepo) SendMessage(message *models.Message) (int, error) {
	var id int
	err := repo.conn.
		QueryRow(
			context.Background(),
			"INSERT INTO messages (chat, author, text) VALUES ($1, $2, $3) RETURNING id",
			message.Chat, message.Author, message.Text).
		Scan(&id)

	if err != nil {
		log.Println(err)
		return -1, err
	}
	return id, nil
}

func (repo *PostgresRepo) GetChatMessages(chatID int) ([]*models.Message, error) {
	rows, err := repo.conn.Query(
		context.Background(),
		"SELECT id, chat, author, text, created_at " +
		"FROM messages WHERE chat = $1 " +
		"ORDER BY messages.created_at DESC", chatID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var messages []*models.Message
	for rows.Next() {
		message := &models.Message{}
		err = rows.Scan(
			&message.ID,
			&message.Chat,
			&message.Author,
			&message.Text,
			&message.CreatedAt)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		messages = append(messages, message)
	}
	if err = rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}
	return messages, nil
}