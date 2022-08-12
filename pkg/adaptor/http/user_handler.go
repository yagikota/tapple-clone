package http

import (
	"net/http"
	"strconv"

	"github.com/CyberAgentHack/2208-ace-go-server/pkg/usecase"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
	uUsecase usecase.IUserUsecase
}

func NewUserHandler(uu usecase.IUserUsecase) *userHandler {
	return &userHandler{
		uUsecase: uu,
	}
}

func (uh *userHandler) findUserByUserID() gin.HandlerFunc {
	return func(c *gin.Context) {
		// パスパラメータ取得
		// TODO: usecaseに渡す前にvalidation
		userID, err := strconv.Atoi(c.Param("user_id"))
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		user, err := uh.uUsecase.FindByUserID(c, userID)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, user)
	}
}

func (uh *userHandler) findUsers() gin.HandlerFunc {
	return func(c *gin.Context) {

		users, err := uh.uUsecase.FindAll(c)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, users)
	}
}

func (uh *userHandler) findRooms() gin.HandlerFunc {
	return func(c *gin.Context) {

		userID, err := strconv.Atoi(c.Param("user_id"))
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		rooms, err := uh.uUsecase.FindAllRooms(c, userID)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, rooms)
	}
}

func (uh *userHandler) findMessages() gin.HandlerFunc {
	return func(c *gin.Context) {

		userID, err := strconv.Atoi(c.Param("user_id"))
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		roomID, err := strconv.Atoi(c.Param("room_id"))
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		messages, err := uh.uUsecase.FindAllRoomMessages(c, userID, roomID)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, messages)
	}
}
