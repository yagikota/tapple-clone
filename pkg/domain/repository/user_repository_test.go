package domain

import (
	"context"
	"testing"
	"time"

	"github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/entity"
	mock "github.com/CyberAgentHack/2208-ace-go-server/pkg/mock/repository"
	"github.com/golang/mock/gomock"
)

func TestFindByUserID(t *testing.T) {
	// 期待値
	testID := 1
	testName := "カイ"
	testIcon := "male/n000029/main_0001_01.jpg"
	testGender := 0
	testBirhday := time.Date(2000, 9, 7, 0, 0, 0, 0, time.Local)
	testLocation := 34

	expectedUser := &entity.User{
		ID:       testID,
		Name:     testName,
		Icon:     testIcon,
		Gender:   testGender,
		Birthday: testBirhday,
		Location: testLocation,
	}

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	MockIUserRepository := mock.NewMockIUserRepository(ctrl)
	MockIUserRepository.EXPECT().FindByUserID(context.Background(), testID).Return(expectedUser, nil)
}
