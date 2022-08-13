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

func RoomFromEntity(entity *entity.Room) *Room {
	r := &Room{
		ID: RoomID(entity.ID),
	}
	r.IsPinned = entity.R.RoomUsers[0].IsPinned
	r.Name = UserFromEntity(entity.R.RoomUsers[0].R.User).Name
	r.Icon = UserFromEntity(entity.R.RoomUsers[0].R.User).Icon
	r.LatestMessage = MessageFromEntity(entity.R.Messages[0])

	// if entity.R != nil {
	// 	if len(entity.R.RoomUsers) != 0 {
	// 		r.IsPinned = entity.R.RoomUsers[0].IsPinned
	// 	}
	// 	if len(entity.R.Messages) != 0 {
	// 		r.LatestMessage = MessageFromEntity(entity.R.Messages[0])
	// 	}
	// 	if len(entity.R.RoomUsers) != 0 {
	// 		if entity.R.RoomUsers != nil {
	// 			uSlice := make(UserSlice, 0, len(entity.R.RoomUsers))
	// 			for i := range entity.R.RoomUsers {
	// 				uSlice = append(uSlice, UserFromEntity(entity.R.RoomUsers[i].R.User))
	// 			}
	// 			r.Users = uSlice
	// 		}
	// 	}
	// }

	return r
}

func RoomMessageFromEntity(entity *entity.Room) *Room {
	rm := &Room{
		ID: RoomID(entity.ID),
	}

	uSlice := make(UserSlice, 0, len(entity.R.RoomUsers))
	for i := range entity.R.RoomUsers {
		uSlice = append(uSlice, UserFromEntity(entity.R.RoomUsers[i].R.User))
	}
	rm.Users = uSlice

	mSlice := make(MessageSlice, 0, len(entity.R.Messages))
	for i := range entity.R.Messages {
		mSlice = append(mSlice, MessageFromEntity(entity.R.Messages[i]))
	}
	rm.Messages = mSlice

	return rm
}
