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
	ID            RoomID    `json:"id"`
	Unread        int       `json:"unread"`
	IsPinned      bool      `json:"is_pinned"`
	LatestMessage *Message  `json:"latest_message"`
	Users         UserSlice `json:"users"`
}

func RoomFromEntity(entity *entity.Room) *Room {
	r := &Room{
		ID: RoomID(entity.ID),
	}
	if entity.R != nil {
		if len(entity.R.RoomUsers) != 0 {
			r.IsPinned = entity.R.RoomUsers[0].IsPinned
		}
		if len(entity.R.Messages) != 0 {
			r.LatestMessage = MessageFromEntity(entity.R.Messages[0])
		}
		if len(entity.R.RoomUsers) != 0 {
			if entity.R.RoomUsers != nil {
				uSlice := make(UserSlice, 0, len(entity.R.RoomUsers))
				for i := range entity.R.RoomUsers {
					uSlice = append(uSlice, UserFromEntity(entity.R.RoomUsers[i].R.User))
				}
				r.Users = uSlice
			}
		}
	}

	return r
}

type MessageID int

type MessageSlice []*Message

type Message struct {
	ID        MessageID `json:"id,omitempty"`
	UserID    int       `json:"user_id"`
	Content   string    `json:"content"`
	IsRead    bool      `json:"is_read"` //TODO:　一応返す　使わなかったら削除
	CreatedAt time.Time `json:"created_at"`
}

func MessageFromEntity(entity *entity.Message) *Message {
	m := &Message{
		ID:        MessageID(entity.ID),
		UserID:    entity.UserID,
		Content:   entity.Content,
		IsRead:    entity.IsRead,
		CreatedAt: entity.CreatedAt,
	}

	return m
}

type RoomMessageSlice []*RoomMessage

type RoomMessage struct {
	ID       RoomID   `json:"id"`
	Users    *User    `json:"users"`
	Messages *Message `json:"messages"`
}

func RoomMessageFromEntity(entity *entity.Message) *RoomMessage {
	rm := &RoomMessage{
		ID: RoomID(entity.RoomID),
	}

	return rm
}
