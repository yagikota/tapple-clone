package http

import (
	"fmt"

	"github.com/CyberAgentHack/2208-ace-go-server/pkg/adaptor/mysql"
	"github.com/CyberAgentHack/2208-ace-go-server/pkg/infra"
	"github.com/CyberAgentHack/2208-ace-go-server/pkg/usecase"
	"github.com/gin-gonic/gin"
)

const (
	apiVersion      = "/v1"
	healthCheckRoot = "/health_check"
	// user系
	usersAPIRoot = apiVersion + "/users"
	userIDParam  = "user_id"
	// room系
	roomsAPIRoot = apiVersion + "/users/:" + userIDParam + "/rooms"
	roomIDParam  = "room_id"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	// ヘルスチェック
	healthCheckGroup := router.Group(healthCheckRoot)
	{
		relativePath := ""
		healthCheckGroup.GET(relativePath, healthCheck())
	}

	mySQLConn := infra.NewMySQLConnector()

	// user
	userRepository := mysql.NewUserRepository(mySQLConn.Conn)
	userUsecase := usecase.NewUserUsecase(userRepository)
	usersGroup := router.Group(usersAPIRoot)
	usersGroup.Use(transactMiddleware(mySQLConn.Conn))
	usersGroup.Use(checkStatusMiddleware())
	{
		relativePath := ""
		userHandler := NewUserHandler(userUsecase)
		// 確認用API
		// v1/users
		usersGroup.GET(relativePath, userHandler.findUsers())
		// v1/users/{user_id}
		relativePath = fmt.Sprintf("/:%s", userIDParam)
		usersGroup.GET(relativePath, userHandler.findUserByUserID())
		// v1/users/{user_id}/profile
		relativePath = fmt.Sprintf("/:%s/profile", userIDParam)
		usersGroup.GET(relativePath, userHandler.findUserDetailByUserID())
	}

	// room
	roomRepository := mysql.NewRoomRepository(mySQLConn.Conn)
	roomUsecase := usecase.NewRoomUsecase(roomRepository)
	roomsGroup := router.Group(roomsAPIRoot)
	roomsGroup.Use(transactMiddleware(mySQLConn.Conn))
	roomsGroup.Use(checkStatusMiddleware())
	{
		relativePath := ""
		roomHandler := NewRoomHandler(roomUsecase)
		// v1/users/{user_id}/rooms
		relativePath = fmt.Sprintf("/:%s/rooms", userIDParam)
		usersGroup.GET(relativePath, roomHandler.findRooms())
		// v1/users/{user_id}/rooms/{room_id}?message_id=xx (xx:ページング用のクエリパラメータ)
		relativePath = fmt.Sprintf("/:%s/rooms/:%s", userIDParam, roomIDParam)
		usersGroup.GET(relativePath, roomHandler.findRoomDetailByRoomID())
		// v1/users/{user_id}/rooms/{room_id}/messages
		relativePath = fmt.Sprintf("/:%s/rooms/:%s/messages", userIDParam, roomIDParam)
		usersGroup.POST(relativePath, roomHandler.sendMessage())
		// v1/users/{user_id}/rooms/{room_id}/images
		relativePath = fmt.Sprintf("/:%s/rooms/:%s/images", userIDParam, roomIDParam)
		usersGroup.POST(relativePath, roomHandler.sendImage())
	}

	return router
}
