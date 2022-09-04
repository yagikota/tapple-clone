package http

import (
	"net/http"
	"strconv"

	"github.com/CyberAgentHack/2208-ace-go-server/pkg/usecase"
	"github.com/CyberAgentHack/2208-ace-go-server/pkg/usecase/model"
	"github.com/gin-gonic/gin"
)

type roomHandler struct {
	rUsecase usecase.IRoomUsecase
}

func NewRoomHandler(ru usecase.IRoomUsecase) *roomHandler {
	return &roomHandler{
		rUsecase: ru,
	}
}

func (rh *roomHandler) findRooms() gin.HandlerFunc {
	return func(c *gin.Context) {

		userID, err := strconv.Atoi(c.Param("user_id"))
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}
		rooms, err := rh.rUsecase.FindAllRooms(c, userID)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, rooms)
	}
}

func (rh *roomHandler) findRoomDetailByRoomID() gin.HandlerFunc {
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

		queryMessageID := c.DefaultQuery("message_id", "0")
		messageID, err := strconv.Atoi(queryMessageID)
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		roomDetail, err := rh.rUsecase.FindRoomDetailByRoomID(c, userID, roomID, messageID)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, roomDetail)
	}
}

func (rh *roomHandler) sendMessage() gin.HandlerFunc {
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

		// リクエストボディーを取り出す
		var newMessage model.NewMessage
		if err := c.ShouldBindJSON(&newMessage); err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		message, err := rh.rUsecase.SendMessage(c, userID, roomID, &newMessage)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, message)
	}
}

func (rh *roomHandler) sendImage() gin.HandlerFunc {
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

		// 画像ファイルを取り出す
		fileHeader, err := c.FormFile("file")
		if err != nil {
			c.AbortWithError(http.StatusBadRequest, err)
			return
		}

		file, err := fileHeader.Open()
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}
		defer file.Close()



		message, err := rh.rUsecase.SendMessage(c, userID, roomID, &newMessage)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		c.JSON(http.StatusOK, message)
	}
}
