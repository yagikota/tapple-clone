package model

import (
	"os"
	"testing"
	"time"

	"github.com/go-playground/assert/v2"
)

var (
	birthday1 time.Time
	birthday2 time.Time
	birthday3 time.Time
	birthday4 time.Time
	prefCode1 int
	prefCode2 int
	prefCode3 int
)

func TestMain(m *testing.M) {
	println("before all...")
	birthday1 = time.Date(2000, 4, 1, 0, 0, 0, 0, time.Local)
	birthday2 = time.Date(1996, 10, 22, 0, 0, 0, 0, time.Local)
	birthday3 = time.Date(2000, 1, 12, 0, 0, 0, 0, time.Local)
	prefCode1 = 0
	prefCode2 = -1
	prefCode3 = 47

	code := m.Run()

	println("after all...")

	os.Exit(code)
}

func TestCalcAge_1(t *testing.T) {
	age, err := calcAge(birthday1)
	assert.Equal(t, age, 22)
	assert.Equal(t, err, nil)
}

func TestCalcAge_2(t *testing.T) {
	age, err := calcAge(birthday2)
	assert.Equal(t, age, 25)
	assert.Equal(t, err, nil)
}
func TestCalcAge_3(t *testing.T) {
	age, err := calcAge(birthday3)
	assert.Equal(t, age, 22)
	assert.Equal(t, err, nil)
}

func TestPrefCodeToPrefKanji_1(t *testing.T) {
	location1 := prefCodeToPrefKanji(prefCode1)
	location2 := prefCodeToPrefKanji(prefCode2)
	location3 := prefCodeToPrefKanji(prefCode3)
	assert.Equal(t, location1, "その他")
	assert.Equal(t, location2, "その他")
	assert.Equal(t, location3, "沖縄")
}
