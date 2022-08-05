package model

type UserProfileImage struct {
	ID        int    `json:"id"`
	UserID    int    `json:"user_id"`
	ImagePath string `json:"path"`
}
