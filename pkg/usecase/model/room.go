package model

import "github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/entity"

type RoomID int

type RoomSlice []*Room

type Room struct {
	ID            RoomID       `json:"id"`
	Unread        int          `json:"unread,omitempty"`
	IsPinned      bool         `json:"is_pinned,omitempty"`
	Name          string       `json:"name,omitempty"`
	Icon          string       `json:"icon,omitempty"`
	LatestMessage *Message     `json:"latest_message,omitempty"`
	Users         UserSlice    `json:"users,omitempty"`
	Messages      MessageSlice `json:"messages,omitempty"`
}

// ルーム一覧で使用
func RoomFromEntity(entity *entity.Room) *Room {
	r := &Room{
		ID: RoomID(entity.ID),
	}

	// TODO: これでもいいのか？
	// r.IsPinned = entity.R.RoomUsers[0].IsPinned
	// r.Name = UserFromEntity(entity.R.RoomUsers[0].R.User).Name
	// r.Icon = UserFromEntity(entity.R.RoomUsers[0].R.User).Icon
	// r.LatestMessage = MessageFromEntity(entity.R.Messages[0])

	if entity.R != nil {
		if entity.R.RoomUsers != nil {
			roomUser := entity.R.RoomUsers[0]
			r.IsPinned = roomUser.IsPinned
			if roomUser.R != nil && roomUser.R.User != nil {
				r.Name = UserFromEntity(roomUser.R.User).Name
				r.Icon = UserFromEntity(roomUser.R.User).Icon
			}
		}
		if entity.R.Messages != nil {
			r.LatestMessage = MessageFromEntity(entity.R.Messages[0])
		}
	}

	return r
}

// ルーム詳細で使用
func RoomDetailFromEntity(entity *entity.Room) *Room {
	rm := &Room{
		ID: RoomID(entity.ID),
	}

	if entity.R != nil {
		if entity.R.RoomUsers != nil {
			uSlice := make(UserSlice, 0, len(entity.R.RoomUsers))
			for _, roomUser := range entity.R.RoomUsers {
				uSlice = append(uSlice, UserFromEntity(roomUser.R.User))
			}
			rm.Users = uSlice
			// ルームの人数が2人の場合、ルームネーム＝相手の名前、ルームアイコン＝相手のアイコン
			const defaultRoomUsersNumber int = 2
			if len(entity.R.RoomUsers) == defaultRoomUsersNumber {
				rm.Name = entity.R.RoomUsers[0].R.User.Name
				rm.Icon = entity.R.RoomUsers[0].R.User.Icon
			}
		}
		if entity.R.Messages != nil {
			mSlice := make(MessageSlice, 0, len(entity.R.Messages))
			for _, message := range entity.R.Messages {
				mSlice = append(mSlice, MessageFromEntity(message))
			}
			rm.Messages = mSlice
		}
	}

	return rm
}
