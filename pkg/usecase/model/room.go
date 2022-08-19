package model

import (
	"sort"
	"strconv"
	"time"

<<<<<<< HEAD
	"github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/model"
=======
	"github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/entity"
	pref "github.com/diverse-inc/jp_prefecture"
>>>>>>> e9e0fce (refactor: リファクタリング)
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
	return &Room{
		ID:            RoomID(m.ID),
		IsPinned:      m.R.RoomUsers[0].IsPinned,
		Name:          UserFromDomainModel(m.R.RoomUsers[0].R.User).Name,
		Icon:          UserFromDomainModel(m.R.RoomUsers[0].R.User).Icon,
		LatestMessage: MessageFromDomainModel(m.R.Messages[0]),
	}

	age, err := calcAge(UserFromDomainModel(m.R.RoomUsers[0].R.User).BirthDay)
	if err != nil {
		return nil
	}

	// 都道府県コードをいい感じに県名に変えてくれるpakage
	prefInfo, ok = pref.FindByCode(UserFromDomainModel(m.R.RoomUsers[0].R.User).Location)
	if !ok {
		location = "その他"
	} else {
		location = prefInfo.KanjiShort()
	}

	r.SubName = strconv.Itoa(age) + "歳・" + location

	return r
}

func prefCodeToPrefKanji(prefCode int) (pref.Prefecture, bool) {
	prefInfo, ok := pref.FindByCode(prefCode)

	return prefInfo, ok
}

// ルーム一覧で使用
func calcAge(birthday time.Time) (int, error) {
	// 現在日時を数字だけで表現 (YYYYMMDD)
	dateFormatOnlyNumber := "20060102" // YYYYMMDD

	nowOnlyNumber := time.Now().Format(dateFormatOnlyNumber)
	birthdayOnlyNumber := birthday.Format(dateFormatOnlyNumber)

	// 日付文字列をそのまま数値化
	nowInt, err := strconv.Atoi(nowOnlyNumber)
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
