package model

import (
	"strconv"
	"time"

	"github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/entity"
	pref "github.com/diverse-inc/jp_prefecture"
)

var (
	prefInfo pref.Prefecture
	ok       bool
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
func RoomFromEntity(entity *entity.Room) *Room {
	r := &Room{
		ID:            RoomID(entity.ID),
		IsPinned:      entity.R.RoomUsers[0].IsPinned,
		Name:          UserFromEntity(entity.R.RoomUsers[0].R.User).Name,
		Icon:          UserFromEntity(entity.R.RoomUsers[0].R.User).Icon,
		LatestMessage: MessageFromEntity(entity.R.Messages[0]),
	}

	age, err := calcAge(UserFromEntity(entity.R.RoomUsers[0].R.User).BirthDay)
	if err != nil {
		return nil
	}

	// 都道府県コードをいい感じに県名に変えてくれるpakage
	prefInfo, ok = pref.FindByCode(UserFromEntity(entity.R.RoomUsers[0].R.User).Location)
	if !ok {
		location := "その他"
		r.SubName = strconv.Itoa(age) + "歳・" + location
	} else {
		r.SubName = strconv.Itoa(age) + "歳・" + prefInfo.KanjiShort()
	}

	return r
}

// ルーム一覧で使用
func calcAge(birthday time.Time) (int, error) {
	// 現在日時を数値のみでフォーマット (YYYYMMDD)
	dateFormatOnlyNumber := "20060102" // YYYYMMDD

	nowOnlyNnmber := time.Now().Format(dateFormatOnlyNumber)
	birthdayOnlyNumber := birthday.Format(dateFormatOnlyNumber)

	// 日付文字列をそのまま数値化
	nowInt, err := strconv.Atoi(nowOnlyNnmber)
	if err != nil {
		return 0, err
	}
	birthdayInt, err := strconv.Atoi(birthdayOnlyNumber)
	if err != nil {
		return 0, err
	}

	// (今日の日付 - 誕生日) / 10000 = 年齢
	age := (nowInt - birthdayInt) / 10000
	return age, nil
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
	rm.Messages = mSlice

	return rm
}
