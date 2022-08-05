package model

// domain entityとは別で定義する。

type UserSlice []*User

type User struct {
	UserID int    `json:"user_id"`
	Name   string `json:"name"`
	Icon   string `json:"icon"`
}

