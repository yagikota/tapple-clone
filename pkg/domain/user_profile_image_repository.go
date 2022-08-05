package domain

import (
	"context"

	"github.com/CyberAgentHack/2208-ace-go-server/domain/entity"
)

// IHogeHogeRepository: Iはinterfaceを意味している
type IUserProfileImageRepository interface {
	UserProfileImage(ctx context.Context) (*entity.UserProfileImage, error)      // 1プロフィールイメージ取得
	UserProfileImages(ctx context.Context) (entity.UserProfileImageSlice, error) // 全プロフィールイメージ取得
}
