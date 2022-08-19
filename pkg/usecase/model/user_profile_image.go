package model

import "github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/model"

type UserProfileImageID int

type UserProfileImageSlice []*UserProfileImage

type UserProfileImage struct {
	ID        UserProfileImageID `json:"id"`
	UserID    int                `json:"user_id"`
	ImagePath string             `json:"image_path"`
}

func UserProfileImageFromDomainModel(m *model.UserProfileImage) *UserProfileImage {
	return &UserProfileImage{
		ID:        UserProfileImageID(m.ID),
		UserID:    m.UserID,
		ImagePath: m.ImagePath,
	}
}
