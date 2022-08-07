package model

import "github.com/CyberAgentHack/2208-ace-go-server/domain/entity"

// domain entityとは別で定義する。

type UserSlice []*User

type User struct {
	UserID        int                   `json:"user_id"`
	Name          string                `json:"name"`
	Icon          string                `json:"icon"`
	ProfileImages UserProfileImageSlice `json:"profile_image"`
}

func UserFromEntity(entity *entity.User) *User {
	u := &User{
		UserID: int(entity.ID),
		Name:   entity.Name,
		Icon:   entity.Icon,
	}

	if entity.R != nil {
		if entity.R.UserProfileImages != nil {
			imgSlice := make(UserProfileImageSlice, 0, len(entity.R.UserProfileImages))
			for _, img := range entity.R.UserProfileImages {
				imgSlice = append(imgSlice, UserProfileImageFromEntity(img))
			}
			u.ProfileImages = imgSlice
		}
	}

	return u
}
