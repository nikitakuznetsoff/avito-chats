package handlers

import (
	"github.com/nikitakuznetsoff/avito-chats/internal/repository"
)

type Handler struct {
	Repo	repository.ChatsRepository
}