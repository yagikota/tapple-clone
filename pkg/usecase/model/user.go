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
	LatestMessage string    `json:"latest_message"`
	CreatedAt     time.Time `json:"created_at"`
	Unread        int       `json:"unread"`
	IsPinned      bool      `json:"is_pinned"`
	User          *User     `json:"user"`
}

func RoomFromEntity(entity *entity.Room) *Room {
	u := &Room{
		ID:        RoomID(entity.ID),
		CreatedAt: entity.CreatedAt,
	}

	// u.Unread = 2 //entity.unread

	return u
}
