package model

import (
	"time"

	"github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/entity"
)

// domain entityとは別で定義する。

type UserID int

type UserSlice []*User

type User struct {
	ID       UserID    `json:"id"`
	Name     string    `json:"name"`
	Icon     string    `json:"icon"`
	Gender   int       `json:"gender"`
	BirthDay time.Time `json:"birthday"`
	Location int       `json:"location"`
}

func UserFromEntity(entity *entity.User) *User {
	u := &User{
		ID:       UserID(entity.ID),
		Name:     entity.Name,
		Icon:     entity.Icon,
		Gender:   entity.Gender,
		BirthDay: entity.Birthday,
		Location: entity.Location,
	}

	return u
}

type RoomID int

type RoomSlice []*Room

type Room struct {
	ID            RoomID   `json:"id"`
	Unread        int      `json:"unread"`
	IsPinned      bool     `json:"is_pinned"`
	LatestMessage *Message `json:"latest_message"`
	User          *User    `json:"user"`
}

func RoomFromEntity(entity *entity.Room) *Room {
	r := &Room{
		ID: RoomID(entity.ID),
	}
	r.IsPinned = entity.R.RoomUsers[0].IsPinned
	r.LatestMessage = MessageFromEntity(entity.R.Messages[0])
	r.User = UserFromEntity(entity.R.RoomUsers[0].R.User)

	return r
}

type MessageID int

type MessageSlice []*Message

type Message struct {
	ID        MessageID `json:"id,omitempty"`
	Content   string    `json:"content"`
	IsRead    bool      `json:"is_read"` //TODO:　一応返す　使わなかったら削除
	CreatedAt time.Time `json:"created_at"`
}

func MessageFromEntity(entity *entity.Message) *Message {
	m := &Message{
		Content:   entity.Content,
		CreatedAt: entity.CreatedAt,
	}

	return m
}
