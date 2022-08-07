package http

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type healthCheckResponse struct {
	Message string `json:"message"`
}

// ヘルスチェック用
//
//	{
//	    "message": "Hello, C Team. you've requested: /health_check"
//	}
//
// が返ってくる
func healthCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		message := fmt.Sprintf("Hello, C Team. you've requested: %s", c.FullPath())

		c.JSON(
			http.StatusOK,
			healthCheckResponse{
				Message: message,
			},
		)
	}
}
