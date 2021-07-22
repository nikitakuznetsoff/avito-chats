package repository

import "github.com/nikitakuznetsoff/avito-chats/internal/models"

type ChatsRepository interface {
	CreateChat(chat *models.Chat) (int, error)
	GetChatByID(id int) (*models.Chat, error)

	GetUserChats(id int) ([]*models.Chat, error)
	GetUsersInChat(chat *models.Chat) ([]int, error)

	SendMessage(message *models.Message) (int, error)
	GetChatMessages(chatID int) ([]*models.Message, error)

	CreateUser(username string) (int, error)
	GetUserByID(id int) (*models.User, error)
}