package main

import (
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi"

	"github.com/Kosfedev/chat-service/pkg/message/http/handlers"
)

const (
	baseURL      = "localhost:8091"
	chatsPostfix = "/chats"
	chatPostfix  = chatsPostfix + "/{id}"
)

func main() {
	r := chi.NewRouter()
	r.Post(chatsPostfix, handlers.CreateChatHandler)
	r.Delete(chatPostfix, handlers.DeleteChatHandler)
	r.Post(chatPostfix, handlers.CreateMessageHandler)

	server := http.Server{
		Addr:         baseURL,
		Handler:      r,
		ReadTimeout:  time.Second * 5,
		WriteTimeout: time.Second * 5,
	}

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
