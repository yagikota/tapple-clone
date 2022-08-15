package model

import (
	"time"

	"github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/entity"
)

// https://github.com/gin-gonic/gin#model-binding-and-validation

type MessageID int

type MessageSlice []*Message

type Message struct {
	ID        MessageID `json:"id,omitempty"`
	UserID    int       `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func MessageFromEntity(entity *entity.Message) *Message {
	m := &Message{
		ID:        MessageID(entity.ID),
		UserID:    entity.UserID,
		Content:   entity.Content,
		CreatedAt: entity.CreatedAt,
	}

	return m
}

type NewMessage struct {
	Content string `json:"content" binding:"required"`
}

// TODO: メソッドと関数の使い分け
func (m *NewMessage) ToEntity(userID int, roomID int) *entity.Message {
	return &entity.Message{
		UserID:  userID,
		RoomID:  roomID,
		Content: m.Content,
	}
}
