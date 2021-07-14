package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/nikitakuznetsoff/avito-chats/internal/models"
)

func (handler *Handler) SendMessage(w http.ResponseWriter, r *http.Request) {
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
		Chat 	int 	`json:"chat"`
		Author 	int		`json:"author"`
		Text 	string	`json:"text"`
	}{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "incorrect response body", http.StatusInternalServerError)
		return
	}

	_, err = handler.Repo.GetChatByID(req.Chat)
	if err != nil {
		http.Error(w, "unknown chat id", http.StatusBadRequest)
		return
	}

	_, err = handler.Repo.GetUserByID(req.Author)
	if err != nil {
		http.Error(w, "unknown user id", http.StatusBadRequest)
		return
	}

	message := models.Message{
		Chat: req.Chat,
		Author: req.Author,
		Text: req.Text,
	}
	id, err := handler.Repo.SendMessage(message)
	if err != nil {
		fmt.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(map[string]int{"message_id": id})
	if err != nil {
		http.Error(w, "error in response creation", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}

func (handler *Handler) GetMessages(w http.ResponseWriter, r *http.Request) {
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

	req := struct{Chat int `json:"chat"`}{}
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, "incorrect request body", http.StatusBadRequest)
		return
	}

	_, err = handler.Repo.GetChatByID(req.Chat)
	if err != nil {
		http.Error(w, "unknown chat id", http.StatusBadRequest)
		return
	}

	messages, err := handler.Repo.GetChatMessages(req.Chat)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(messages)
	if err != nil {
		http.Error(w, "error in response creation", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}