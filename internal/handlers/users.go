package handlers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func (handler *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
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

	req := struct{Username string `json:"username"`}{}
	err = json.Unmarshal(body, &req)
	if err != nil || req.Username == "" {
		http.Error(w, "incorrect response body", http.StatusBadRequest)
		return
	}

	id, err := handler.Repo.CreateUser(req.Username)
	if err != nil {
		http.Error(w, "error in user creation", http.StatusInternalServerError)
		return
	}

	resp, err := json.Marshal(map[string]int{"user_id": id})
	if err != nil {
		http.Error(w, "error in creation response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(resp)
}
