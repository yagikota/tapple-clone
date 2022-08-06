package http

import (
	"fmt"

	"github.com/CyberAgentHack/2208-ace-go-server/adaptor/mysql"
	"github.com/CyberAgentHack/2208-ace-go-server/infra"
	"github.com/CyberAgentHack/2208-ace-go-server/usecase"
	"github.com/gin-gonic/gin"
)

const (
	apiVersion  = "/v1"
	userAPIRoot = apiVersion + "/users"
	userIDParam = "user_id"
)

func Router() *gin.Engine {
	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	// TODO:
	mySQLConn := infra.NewMySQLConnector()

	// 別ファイルにした方がいいかも
	userRepository := mysql.NewUserRepository(mySQLConn.Conn)
	userUsecase := usecase.NewUserUsecase(userRepository)

	usersGroup := router.Group(userAPIRoot)
	{
		userHandler := NewUserHandler(userUsecase)

		// /users/{user_id}
		relativePath := fmt.Sprintf("/:%s", userIDParam)
		usersGroup.GET(relativePath, userHandler.getUser())
	}

	return router
}
