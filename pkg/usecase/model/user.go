package model

import (
	"time"

	"github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/model"
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

func UserFromEntity(entity *model.User) *User {
	return &User{
		ID:       UserID(entity.ID),
		Name:     entity.Name,
		Icon:     entity.Icon,
		Gender:   entity.Gender,
		BirthDay: entity.Birthday,
		Location: entity.Location,
	}
}
