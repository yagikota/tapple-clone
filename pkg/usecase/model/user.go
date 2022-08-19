package model

import (
	"time"

	"github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/model"
)

// domain modelを抽象化して使い回す

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

func UserFromDomainModel(m *model.User) *User {
	return &User{
		ID:       UserID(m.ID),
		Name:     m.Name,
		Icon:     m.Icon,
		Gender:   m.Gender,
		BirthDay: m.Birthday,
		Location: m.Location,
	}
}
