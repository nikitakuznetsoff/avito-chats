package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"log"
	"net/http"
	"os"

	"github.com/nikitakuznetsoff/avito-chats/internal/handlers"
	"github.com/nikitakuznetsoff/avito-chats/internal/repository/postgres"
)

const (
	port = ":9000"
	dsn = "postgres://nick:pass@localhost:5432/linksdb"
)

func main() {
	conn, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	repo := postgres.CreatePostgresRepo(conn)
	handler := handlers.Handler{Repo: repo}

	http.HandleFunc("/users/add", handler.CreateUser)
	http.HandleFunc("/chats/add", handler.CreateChat)
	http.HandleFunc("/chats/get", handler.GetChats)
	http.HandleFunc("/messages/get", handler.GetMessages)
	http.HandleFunc("/messages/add", handler.SendMessage)

	fmt.Printf("Starting server on port %v\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}