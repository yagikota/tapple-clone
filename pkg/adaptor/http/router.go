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
	usersAPIRoot    = apiVersion + "/users"
	userIDParam     = "user_id"
	roomIDParam     = "room_id"
)

func InitRouter() *gin.Engine {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	// TODO:
	// 別ファイルにした方がいいかも
	mySQLConn := infra.NewMySQLConnector()

	userRepository := mysql.NewUserRepository(mySQLConn.Conn)
	userUsecase := usecase.NewUserUsecase(userRepository)

	healthCheckGroup := router.Group(healthCheckRoot)
	{
		relativePath := ""
		healthCheckGroup.GET(relativePath, healthCheck())
	}

	usersGroup := router.Group(usersAPIRoot)
	// TODO: 引数が正しく無い気がする
	usersGroup.Use(TransactMiddleware(mySQLConn.Conn))
	{
		relativePath := ""
		userHandler := NewUserHandler(userUsecase)

		// 確認用API
		// v1/users
		usersGroup.GET(relativePath, userHandler.findUsers())
		// v1/users/{user_id}
		relativePath = fmt.Sprintf("/:%s", userIDParam)
		usersGroup.GET(relativePath, userHandler.findUserByUserID())
		// v1/users/{user_id}/rooms
		relativePath = fmt.Sprintf("/:%s/rooms", userIDParam)
		usersGroup.GET(relativePath, userHandler.findRooms())
		// v1/users/{user_id}/rooms/{room_id}
		relativePath = fmt.Sprintf("/:%s/rooms/:%s", userIDParam, roomIDParam)
		usersGroup.GET(relativePath, userHandler.findRoomDetailByRoomID())
		// v1/users/{user_id}/rooms/{room_id}/messages
		relativePath = fmt.Sprintf("/:%s/rooms/:%s/messages", userIDParam, roomIDParam)
		usersGroup.POST(relativePath, userHandler.sendMessage())
		// v1/users/{user_id}/profile
		relativePath = fmt.Sprintf("/:%s/profile", userIDParam)
		usersGroup.GET(relativePath, userHandler.findUserDetailByUserID())
	}

	return router
}
