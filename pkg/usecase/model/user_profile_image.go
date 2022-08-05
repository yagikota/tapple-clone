package model

import "github.com/CyberAgentHack/2208-ace-go-server/domain/entity"

type UserProfileImageSlice []*UserProfileImage

type UserProfileImage struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	ImagePath string `json:"path"`
}

func UserProfileImageFromEntity(entity *entity.UserProfileImage) *UserProfileImage {
	return &UserProfileImage{
		ID:        int(entity.ID),
		UserID:    int(entity.UserID),
		ImagePath: entity.ImagePath,
	}
}
