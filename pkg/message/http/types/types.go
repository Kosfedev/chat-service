package types

import "time"

// NewChatData is ...
type NewChatData struct {
	Usernames []string `json:"usernames"`
}

// NewMessageData is ...
type NewMessageData struct {
	From      string    `json:"from"`
	Text      string    `json:"text"`
	Timestamp time.Time `json:"timestamp"`
}
