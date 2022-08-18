package enum

type Prefecture int

const (
	Sonota Prefecture = iota
	Hokkaido
	Aomori
	Iwate
	Miyagi
	Akita
	Yamagata
	Fukushima
	Ibaraki
	Tochigi
	Gunma
	Saitama
	Chiba
	Tokyo
	Kanagawa
	Niigata
	Yamanashi
	Nagano
	Toyama
	Ishikawa
	Fukui
	Gifu
	Shizuoka
	Aichi
	Mie
	Shiga
	Kyoto
	Osaka
	Hyogo
	Nara
	Wakayama
	Tottori
	Shimane
	Oakayama
	Hiroshima
	Yamaguchi
	Tokushima
	Kagawa
	Ehime
	Kochi
	Fukuoka
	Saga
	Nagasaki
	Kumamoto
	Oita
	Miyazaki
	Kagoshima
	Okinawa
)

func (p Prefecture) String() string {
	switch p {
	case Sonota:
		return "その他"
	case Hokkaido:
		return "北海道"
	case Aomori:
		return "青森"
	case Iwate:
		return "岩手"
	case Miyagi:
		return "宮城"
	case Akita:
		return "秋田"
	case Yamagata:
		return "山形"
	case Fukushima:
		return "福島"
	case Ibaraki:
		return "茨城"
	case Tochigi:
		return "栃木"
	case Gunma:
		return "群馬"
	case Saitama:
		return "埼玉"
	case Chiba:
		return "千葉"
	case Tokyo:
		return "東京"
	case Kanagawa:
		return "神奈川"
	case Niigata:
		return "新潟"
	case Yamanashi:
		return "山梨"
	case Nagano:
		return "長野"
	case Toyama:
		return "富山"
	case Ishikawa:
		return "石川"
	case Fukui:
		return "福井"
	case Gifu:
		return "岐阜"
	case Shizuoka:
		return "静岡"
	case Aichi:
		return "愛知"
	case Mie:
		return "三重"
	case Shiga:
		return "滋賀"
	case Kyoto:
		return "京都"
	case Osaka:
		return "大阪"
	case Hyogo:
		return "兵庫"
	case Nara:
		return "奈良"
	case Wakayama:
		return "和歌山"
	case Tottori:
		return "鳥取"
	case Shimane:
		return "島根"
	case Oakayama:
		return "岡山"
	case Hiroshima:
		return "広島"
	case Yamaguchi:
		return "山口"
	case Tokushima:
		return "徳島"
	case Kagawa:
		return "香川"
	case Ehime:
		return "愛媛"
	case Kochi:
		return "高知"
	case Fukuoka:
		return "福岡"
	case Saga:
		return "佐賀"
	case Nagasaki:
		return "長崎"
	case Kumamoto:
		return "熊本"
	case Oita:
		return "大分"
	case Miyazaki:
		return "宮崎"
	case Kagoshima:
		return "鹿児島"
	case Okinawa:
		return "沖縄"
	}
	return p.String()
}
