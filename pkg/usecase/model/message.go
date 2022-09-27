package model

import (
	"time"

	"github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/model"
)

// https://github.com/gin-gonic/gin#model-binding-and-validation

type MessageID int

type MessageSlice []*Message

type Message struct {
	ID        MessageID `json:"id,omitempty"`
	User      *User     `json:"user"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func MessageFromDomainModel(m *model.Message) *Message {
	return &Message{
		ID:        MessageID(m.ID),
		User:      UserFromDomainModel(m.R.User),
		Content:   m.Content,
		CreatedAt: m.CreatedAt,
	}
}

type NewMessage struct {
	Content string `json:"content" binding:"required"`
}

// TODO: メソッドと関数の使い分け
func (m *NewMessage) ToDomainModel(userID, roomID int) *model.Message {
	return &model.Message{
		UserID:  userID,
		RoomID:  roomID,
		Content: m.Content,
	}
}
