package postgres

import (
	"context"
	"github.com/jackc/pgx/v4"
	"github.com/nikitakuznetsoff/avito-chats/internal/models"
	"log"
)

func (repo *PostgresRepo) CreateChat(chat *models.Chat) (int, error){
	// Проверка наличия в БД всех пользователей
	userIDs := make([]int, len(chat.Users))
	for i, userID := range chat.Users {
		_, err := repo.GetUserByID(userID)
		if err != nil {
			log.Println(err)
			return -1, err
		}
		userIDs[i] = userID
	}
	var id int
	err := repo.conn.
		QueryRow(context.Background(),
		"INSERT INTO chats (name) VALUES ($1) RETURNING id", chat.Name).
		Scan(&id)
	if err != nil {
		log.Println(err)
		return -1, err
	}
	// Добавление пользователей в отношение c чатом
	for i := range chat.Users {
		_, err := repo.conn.Exec(
			context.Background(),
			"INSERT INTO user_chat_relation (chat_id, user_id) VALUES ($1, $2)",
			id, userIDs[i])
		if err != nil {
			log.Println(err)
			return -1, err
		}
	}
	return id, nil
}

func (repo *PostgresRepo) GetChatByID(id int) (*models.Chat, error) {
	chat := &models.Chat{}
	err := repo.conn.
		QueryRow(
			context.Background(),
			"SELECT id, name, created_at FROM chats WHERE id = $1", id).
		Scan(&chat.ID, &chat.Name, &chat.CreatedAt)
	if err != nil {
		return nil, err
	}
	// Находим список пользователей в чате
	users, err := repo.GetUsersInChat(chat)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	chat.Users = users
	return chat, nil
}

// Получение списка чатов пользователя
func (repo *PostgresRepo) GetUserChats(id int) ([]*models.Chat, error) {
	// Подзапрос для получения времени последнего сообщения в чате
	// для возможности сортировки по убыванию
	rows, err := repo.conn.Query(context.Background(),
		"SELECT chats.id, chats.name, chats.created_at FROM " +
			"(chats JOIN " +
				"(SELECT chat, max(created_at) as last_time FROM messages GROUP BY chat) as t1 " +
			"ON chats.id=t1.chat) " +
			"JOIN user_chat_relation ON chats.id=user_chat_relation.chat_id " +
			"WHERE user_chat_relation.user_id = $1 " +
			"ORDER BY t1.last_time DESC", id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var chats []*models.Chat
	for rows.Next() {
		chat := &models.Chat{}
		err = rows.Scan(&chat.ID, &chat.Name, &chat.CreatedAt)
		if err != nil {
			log.Println(err)
			return nil, err
		}

		users, err := repo.GetUsersInChat(chat)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		chat.Users = users
		chats = append(chats, chat)
	}
	if err = rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}
	return chats, err
}
// Получение списка идентификаторов пользователей чата
func (repo *PostgresRepo) GetUsersInChat(chat *models.Chat) ([]int, error) {
	rows, err := repo.conn.Query(context.Background(),
		"SELECT user_id FROM user_chat_relation WHERE chat_id = $1",
		chat.ID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer rows.Close()

	var users []int
	for rows.Next() {
		user := 0
		err = rows.Scan(&user)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		users = append(users, user)
	}
	if err = rows.Err(); err != nil {
		log.Println(err)
		if err == pgx.ErrNoRows {
			return users, nil
		}
		return nil, err
	}
	return users, nil
}