package handlers

import (
	"crypto/rand"
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"

	"github.com/Kosfedev/chat-service/pkg/message/http/types"
)

// CreateChatHandler is...
func CreateChatHandler(w http.ResponseWriter, r *http.Request) {
	newChat := &types.NewChatData{}
	if err := json.NewDecoder(r.Body).Decode(newChat); err != nil {
		http.Error(w, "Failed to decode new chat data", http.StatusBadRequest)
		return
	}

	nBig, err := rand.Int(rand.Reader, big.NewInt(27))
	if err != nil {
		panic(err)
	}
	id := nBig.Int64()
	fmt.Println("new chat id:", id)
	fmt.Printf("new chat data: %#v\n", newChat.Usernames)
}

// DeleteChatHandler is...
func DeleteChatHandler(w http.ResponseWriter, r *http.Request) {
	chatIDStr := chi.URLParam(r, "id")
	chatID, err := strconv.ParseInt(chatIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid chat ID", http.StatusBadRequest)
		return
	}

	fmt.Printf("delete chat id: %v\n", chatID)
}

// CreateMessageHandler is...
func CreateMessageHandler(w http.ResponseWriter, r *http.Request) {
	newMessageData := &types.NewMessageData{}
	chatIDStr := chi.URLParam(r, "id")
	chatID, err := strconv.ParseInt(chatIDStr, 10, 64)
	if err != nil {
		http.Error(w, "Invalid chat ID", http.StatusBadRequest)
		return
	}
	if err := json.NewDecoder(r.Body).Decode(newMessageData); err != nil {
		http.Error(w, "Failed to decode new message data", http.StatusBadRequest)
		return
	}

	fmt.Printf("get chat id: %v\n", chatID)
	fmt.Printf("get message: %#v\n", *newMessageData)
}
