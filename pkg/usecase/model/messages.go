package model

import (
	"time"

	"github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/entity"
)

// https://github.com/gin-gonic/gin#model-binding-and-validation

type Message struct {
	ID        int64     `json:"id"`
	UserID    int       `json:"user_id"`
	RoomID    int       `json:"room_id"`
	Content   string    `json:"content"`
	IsRead    bool      `json:"is_read"`
	CreatedAt time.Time `json:"created_at"`
}

type NewMessage struct {
	Content string `json:"content" binding:"required"`
}

// TODO: メソッドにする
func MessageToEntity(m *NewMessage, userID int, roomID int) *entity.Message {
	return &entity.Message{
		UserID:  userID,
		RoomID:  roomID,
		Content: m.Content,
	}
}
