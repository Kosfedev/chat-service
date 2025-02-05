package main

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"time"

	"github.com/go-chi/chi"
)

const (
	baseURL      = "localhost:8081"
	chatsPostfix = "/chats"
	chatPostfix  = chatsPostfix + "/{id}"
)

// NewChatData is ...
type NewChatData struct {
	chatnames []string `json:"chatnames"`
}

// NewMessageData is ...
type NewMessageData struct {
	From      string    `json:"from"`
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
}

func createChatHandler(w http.ResponseWriter, r *http.Request) {
	newchat := &NewChatData{}
	if err := json.NewDecoder(r.Body).Decode(newchat); err != nil {
		http.Error(w, "Failed to decode new chat data", http.StatusBadRequest)
		return
	}

	id := createChat(newchat)

	fmt.Println("new chat id:", id)
}

func createChat(chat *NewChatData) int64 {
	id := generatechatID()

	fmt.Printf("new chat data: %+v\n", *chat)

	return id
}

func generatechatID() int64 {
	nBig, err := rand.Int(rand.Reader, big.NewInt(27))
	if err != nil {
		panic(err)
	}

	return nBig.Int64()
}

func deleteChatHandler(w http.ResponseWriter, r *http.Request) {
	chatIDStr := chi.URLParam(r, "id")
	chatID, err := parseID(chatIDStr)
	if err != nil {
		http.Error(w, "Invalid chat ID", http.StatusBadRequest)
		return
	}

	deleteChat(chatID)
}

func deleteChat(id int64) {
	fmt.Printf("delete chat id: %v\n", id)
}

func parseID(idStr string) (int64, error) {
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func sendMessageHandler(w http.ResponseWriter, r *http.Request) {
	newMessageData := &NewMessageData{}
	chatIDStr := chi.URLParam(r, "id")
	chatID, err := parseID(chatIDStr)
	if err != nil {
		http.Error(w, "Invalid chat ID", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(newMessageData); err != nil {
		http.Error(w, "Failed to decode new message data", http.StatusBadRequest)
		return
	}

	sendMessage(chatID, newMessageData)
}

func sendMessage(id int64, message *NewMessageData) {
	fmt.Printf("get chat id: %v\n", id)
	fmt.Printf("get message: %+v\n", *message)
}

func main() {
	r := chi.NewRouter()
	r.Post(chatsPostfix, createChatHandler)
	r.Delete(chatPostfix, deleteChatHandler)
	r.Post(chatPostfix, sendMessageHandler)

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
