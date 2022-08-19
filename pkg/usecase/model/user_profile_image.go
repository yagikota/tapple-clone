package model

import "github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/model"

type UserProfileImageID int

type UserProfileImageSlice []*UserProfileImageSlice

type UserProfileImage struct {
	ID        UserProfileImageID `json:"id"`
	ImagePath string             `json:"url"`
}

func UserProfileImageFromDomainEntity(m *model.UserProfileImage) *UserProfileImage {
	return &UserProfileImage{
		ID:        UserProfileImageID(m.UserID),
		ImagePath: m.ImagePath,
	}
}
