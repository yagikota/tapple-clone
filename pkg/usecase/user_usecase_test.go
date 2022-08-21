package usecase

import (
	"context"
	"strconv"
	"testing"
	"time"

	dmodel "github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/model"
	mock "github.com/CyberAgentHack/2208-ace-go-server/pkg/mock/service"
	"github.com/CyberAgentHack/2208-ace-go-server/pkg/usecase/model"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/suite"
)

// ----- BEGIN デフォルトのテストデータ -----
var (
	defaultTime        time.Time
	userID             int
	roomID             int
	messageID          int
	UserProfileImageID int
)

func prepareUserDomainModel(id, gender, location int) *dmodel.User {
	user := new(dmodel.User)
	user.ID = id
	user.Name = "name" + strconv.Itoa(id)
	user.Icon = "icon" + strconv.Itoa(id)
	user.Gender = gender
	user.Birthday = defaultTime
	user.Location = location
	user.IsPrincipal = true
	user.R = user.R.NewStruct()
	user.R.Hobbies = dmodel.HobbySlice{
		{
			ID:     1,
			UserID: 1,
			Tag:    "tag1",
		},
	}
	user.R.UserProfileImages = dmodel.UserProfileImageSlice{
		{
			ID:        1,
			UserID:    1,
			ImagePath: "image1",
		},
	}

	return user
}

func prepareUser(id, gender int, location string) *model.User {
	return &model.User{
		ID:          model.UserID(id),
		Name:        "name" + strconv.Itoa(id),
		Icon:        "icon" + strconv.Itoa(id),
		Gender:      gender,
		BirthDay:    defaultTime,
		Location:    location,
		IsPrincipal: true,
	}
}

func prepareUserDetail(id, gender int, location string) *model.UserDetail {
	return &model.UserDetail{
		ID:          model.UserID(id),
		Name:        "name" + strconv.Itoa(id),
		Age:         0,
		Location:    location,
		IsPrincipal: true,
		TagCount:    1,
		ProfileImages: []*model.UserProfileImage{
			{
				ID:        1,
				UserID:    1,
				ImagePath: "image1",
			},
		},
		Hobbies: []*model.Hobby{
			{
				ID:  1,
				Tag: "tag1",
			},
		},
	}
}

// ----- END デフォルトのテストデータ -----

type UserUsecaseTestSuite struct {
	suite.Suite
	mock    *mock.MockIUserService
	usecase IUserUsecase
}

func (suite *UserUsecaseTestSuite) SetupSuite() {
	mockCtl := gomock.NewController(suite.T())
	defer mockCtl.Finish()
	suite.mock = mock.NewMockIUserService(mockCtl)
	suite.usecase = NewUserUsecase(suite.mock)

	defaultTime = time.Date(2022, 4, 1, 0, 0, 0, 0, time.UTC)
	userID = 1
	roomID = 1
	messageID = 1
}

func TestUserHandlerSuite(t *testing.T) {
	suite.Run(t, new(UserUsecaseTestSuite))
}

func (suite *UserUsecaseTestSuite) Test_userUsecase_FindUserByUserID() {
	suite.mock.EXPECT().FindUserByUserID(context.Background(), 1).Return(prepareUserDomainModel(1, 0, 0), nil)
	res, err := suite.usecase.FindUserByUserID(context.Background(), 1)
	suite.Equal(err, nil)
	suite.Equal(res, prepareUser(1, 0, "その他"))
}

func (suite *UserUsecaseTestSuite) Test_userUsecase_FindAllUsers() {
	suite.mock.EXPECT().FindAllUsers(context.Background()).Return(
		dmodel.UserSlice{prepareUserDomainModel(1, 0, 0), prepareUserDomainModel(2, 1, 1)},
		nil,
	)
	res, err := suite.usecase.FindAllUsers(context.Background())
	suite.Equal(err, nil)
	suite.Equal(
		res,
		model.UserSlice{prepareUser(1, 0, "その他"), prepareUser(2, 1, "北海道")},
	)
}

func (suite *UserUsecaseTestSuite) Test_userUsecase_FindUserDetailByUserID() {
	suite.mock.EXPECT().FindUserDetailByUserID(context.Background(), 1).Return(prepareUserDomainModel(1, 0, 0), nil)
	res, err := suite.usecase.FindUserDetailByUserID(context.Background(), 1)
	suite.Equal(err, nil)
	suite.Equal(res, prepareUserDetail(1, 0, "その他"))
}
