package http

import (
	"net/http"
	"strconv"

	"github.com/CyberAgentHack/2208-ace-go-server/usecase"
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

func (uh *userHandler) getUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		// パスパラメータ取得
		userID, err := strconv.Atoi(c.Param("user_id"))
		if err != nil {
			c.Status(http.StatusInternalServerError)
		}

		user, err := uh.uUsecase.User(c, userID)
		if err != nil {
			c.Status(http.StatusInternalServerError)
		}

		c.JSON(http.StatusOK, user)
	}
}
