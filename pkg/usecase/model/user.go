package model

import (
	"time"

	"github.com/CyberAgentHack/2208-ace-go-server/domain/entity"
)

// domain entityとは別で定義する。

type UserSlice []*User

type User struct {
	UserID   int       `json:"user_id"`
	Name     string    `json:"name"`
	Icon     string    `json:"icon"`
	Gender   int       `json:"gender"`
	BirthDay time.Time `json:"birthday"`
	Location int       `json:"location"`
}

func UserFromEntity(entity *entity.User) *User {
	u := &User{
		UserID:   int(entity.ID),
		Name:     entity.Name,
		Icon:     entity.Icon,
		Gender:   entity.Gender,
		BirthDay: entity.Birthday,
		Location: entity.Location,
	}

	return u
}
