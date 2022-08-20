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

type UserDetail struct {
	ID            UserID                `json:"id"`
	Name          string                `json:"name"`
	Age           int                   `json:"age"`
	Location      string                `json:"location"`
	IsPrincipal   bool                  `json:"is_principal"`
	TagCount      int                   `json:"tag_count"`
	ProfileImages UserProfileImageSlice `json:"profile_images"`
	Hobbies       HobbySlice            `json:"hobbies"`
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

func UserDetailFromDomainModel(m *model.User) *UserDetail {
	ud := &UserDetail{
		ID:   UserID(m.ID),
		Name: m.Name,
	}

	age, err := calcAge(m.Birthday)
	if err != nil {
		return nil
	}
	ud.Age = age

	// 都道府県コードを県名に変換
	ud.Location = prefCodeToPrefKanji(m.Location)

	uSlice := make(UserProfileImageSlice, 0, len(m.R.UserProfileImages))
	for _, profileImage := range m.R.UserProfileImages {
		uSlice = append(uSlice, UserProfileImageFromDomainModel(profileImage))
	}
	ud.ProfileImages = uSlice

	numberOfHobbyTag := len(m.R.Hobbies)
	ud.TagCount = numberOfHobbyTag
	hSlice := make(HobbySlice, 0, numberOfHobbyTag)
	for _, hobby := range m.R.Hobbies {
		hSlice = append(hSlice, HobbyFromDomainModel(hobby))
	}
	ud.Hobbies = hSlice

	return ud
}
