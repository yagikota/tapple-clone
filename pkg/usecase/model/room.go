package model

import (
	"sort"

	constant "github.com/CyberAgentHack/2208-ace-go-server/pkg"
	"github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/entity"
)

type RoomID int

type RoomSlice []*Room

// ルーム一覧で使用
type Room struct {
	ID            RoomID   `json:"id"`
	Unread        int      `json:"unread"`
	IsPinned      bool     `json:"is_pinned"`
	Name          string   `json:"name"`
	Icon          string   `json:"icon"`
	LatestMessage *Message `json:"latest_message"`
}

// ルーム一覧で使用
type Rooms struct {
	Rooms RoomSlice `json:"rooms"`
}

// ルーム詳細で使用
type RoomDetail struct {
	ID       RoomID       `json:"id"`
	Name     string       `json:"name"`
	Icon     string       `json:"icon"`
	Users    UserSlice    `json:"users"`
	Messages MessageSlice `json:"messages"`
	IsLast   bool         `json:"is_last"`
}

// ルーム一覧で使用
func RoomFromEntity(entity *entity.Room) *Room {
	return &Room{
		ID:            RoomID(entity.ID),
		IsPinned:      entity.R.RoomUsers[0].IsPinned,
		Name:          UserFromEntity(entity.R.RoomUsers[0].R.User).Name,
		Icon:          UserFromEntity(entity.R.RoomUsers[0].R.User).Icon,
		LatestMessage: MessageFromEntity(entity.R.Messages[0]),
	}
}

// ルーム詳細で使用
func RoomDetailFromEntity(entity *entity.Room) *RoomDetail {
	rm := &RoomDetail{
		ID:   RoomID(entity.ID),
		Name: entity.R.RoomUsers[0].R.User.Name,
		Icon: entity.R.RoomUsers[0].R.User.Icon,
	}

	uSlice := make(UserSlice, 0, len(entity.R.RoomUsers))
	for _, roomUser := range entity.R.RoomUsers {
		uSlice = append(uSlice, UserFromEntity(roomUser.R.User))
	}
	rm.Users = uSlice

	mSlice := make(MessageSlice, 0, len(entity.R.Messages))
	for _, message := range entity.R.Messages {
		mSlice = append(mSlice, MessageFromEntity(message))
	}

	// 古い順番にソート
	sort.Slice(mSlice, func(i, j int) bool {
		return mSlice[i].CreatedAt.Before(mSlice[j].CreatedAt)
	})
	rm.Messages = mSlice

	// 取得したメッセージの数がLIMIT_RECORDより少なくなったらフラグをtureに変更
	if len(mSlice) < constant.LIMIT_RECORD {
		rm.IsLast = true
	}

	return rm
}
