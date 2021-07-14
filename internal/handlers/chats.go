package handlers

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/nikitakuznetsoff/avito-chats/internal/models"
)

func (handler *Handler) CreateChat(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "incorrect method", http.StatusBadRequest)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req := struct{
		Name string `json:"name"`
		Users []int	`json:"users"`
	}{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "incorrect request body", http.StatusInternalServerError)
		return
	}

	// Проверка на уникальность идентификаторов пользователей
	uniqueMap := make(map[int]bool)
	for _, userID := range req.Users {
		if !uniqueMap[userID] {
			uniqueMap[userID] = true
		} else {
			http.Error(w, "user id's aren't unique", http.StatusBadRequest)
			return
		}
	}

	chat := models.Chat{
		Name: req.Name,
		Users: req.Users,
	}
	id, err := handler.Repo.CreateChat(chat)
	if err != nil {
		http.Error(w, "error in char creation", http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(map[string]int{"chat_id": id})
	if err != nil {
		http.Error(w, "error in response creation", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}

func (handler *Handler) GetChats(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "incorrect method", http.StatusBadRequest)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	defer r.Body.Close()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	req := struct{User int `json:"user"`}{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "incorrect request body", http.StatusBadRequest)
		return
	}

	chats, err := handler.Repo.GetUserChats(req.User)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "user not found", http.StatusNotFound)
		} else {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		return
	}

	resp, err := json.Marshal(chats)
	if err != nil {
		http.Error(w, "error in response creation", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}