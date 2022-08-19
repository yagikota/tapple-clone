package model

import (
	"sort"
	"strconv"
	"time"

	"github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/model"

	pref "github.com/diverse-inc/jp_prefecture"
)

type RoomID int

type RoomSlice []*Room

// ルーム一覧で使用
type Room struct {
	ID            RoomID   `json:"id"`
	Unread        int      `json:"unread"`
	IsPinned      bool     `json:"is_pinned"`
	Name          string   `json:"name"`
	SubName       string   `json:"sub_name"`
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
}

// ルーム一覧で使用
func RoomFromDomainModel(m *model.Room) *Room {
	u := UserFromDomainModel(m.R.RoomUsers[0].R.User)
	r := &Room{
		ID:            RoomID(m.ID),
		IsPinned:      m.R.RoomUsers[0].IsPinned,
		Name:          u.Name,
		Icon:          u.Icon,
		LatestMessage: MessageFromDomainModel(m.R.Messages[0]),
	}

	age, err := calcAge(u.BirthDay)
	if err != nil {
		return nil
	}

	// 都道府県コードを県名に変換
	location := prefCodeToPrefKanji(u.Location)
	r.SubName = strconv.Itoa(age) + "歳・" + location

	// 未読数計算
	unread := 0
	partnerID := UserFromEntity(entity.R.RoomUsers[0].R.User).ID
	for _, message := range entity.R.Messages {
		if message.UserID == int(partnerID) && message.IsRead {
			break
		}
		if message.UserID == int(partnerID) && !message.IsRead {
			unread += 1
		}
	}
	r.Unread = unread

	return r
}

func prefCodeToPrefKanji(prefCode int) string {
	location := "その他"
	prefInfo, ok := pref.FindByCode(prefCode)
	if ok {
		location = prefInfo.KanjiShort()
	}

	return location
}

// ルーム一覧で使用
func calcAge(birthday time.Time) (int, error) {
	// タイムゾーンをJSTに設定
	now := time.Now()
	nowUTC := now.UTC()
	jst := time.FixedZone("Asia/Tokyo", 9*60*60)

	nowJST := nowUTC.In(jst)

	thisYear, thisMonth, thisDay := nowJST.Date()
	age := thisYear - birthday.Year()

	// 誕生日を迎えていない時の処理
	if thisMonth < birthday.Month() && thisDay < birthday.Day() {
		age -= 1
	}

	return age, nil
}

// ルーム詳細で使用
func RoomDetailFromDomainModel(m *model.Room) *RoomDetail {
	rm := &RoomDetail{
		ID:   RoomID(m.ID),
		Name: m.R.RoomUsers[0].R.User.Name,
		Icon: m.R.RoomUsers[0].R.User.Icon,
	}

	uSlice := make(UserSlice, 0, len(m.R.RoomUsers))
	for _, roomUser := range m.R.RoomUsers {
		uSlice = append(uSlice, UserFromDomainModel(roomUser.R.User))
	}
	rm.Users = uSlice

	mSlice := make(MessageSlice, 0, len(m.R.Messages))
	for _, message := range m.R.Messages {
		mSlice = append(mSlice, MessageFromDomainModel(message))
	}

	// 古い順番にソート
	sort.Slice(mSlice, func(i, j int) bool {
		return mSlice[i].CreatedAt.Before(mSlice[j].CreatedAt)
	})
	rm.Messages = mSlice

	return rm
}
