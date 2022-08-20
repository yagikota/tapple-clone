package http

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/CyberAgentHack/2208-ace-go-server/pkg/adaptor/mysql"
	"github.com/gin-gonic/gin"
)

// TODO:
// トランザクションのベストプラクティス聞く。
// commit rollbackのタイミングあっているか
// panicはいつするのか

// トランザクション開始
func beginTxAndSetToContext(c *gin.Context, conn *sql.DB) (*sql.Tx, error) {
	tx, err := conn.Begin()
	if err != nil {
		return nil, err
	}
	log.Println("begin database transaction")
	return tx, nil
}

// トランザクション用のmiddleware
func transactMiddleware(conn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tx, err := beginTxAndSetToContext(c, conn)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		// TODO: 処理が汚いので整理する
		// Rollbackのあとエラーハンドリングするのか？
		// panic起こしていいのか
		defer func() {
			// panicした場合(優先度がerror handlingよりも高い)
			if r := recover(); r != nil {
				if rbErr := tx.Rollback(); rbErr != nil {
					log.Println("rollback error: ", rbErr) // DB接続が切れた時など
					c.Error(rbErr)
					return
				}
				log.Println("successful rollback")
				return
			}

			// エラーが発生した場合(こちらを先にした場合、panicがハンドリンングされない)
			if err := c.Errors.Last(); err != nil {
				log.Println(err)
				if rbErr := tx.Rollback(); rbErr != nil {
					log.Println("rollback error: ", rbErr) // DB接続が切れた時など
					c.Error(rbErr)
					return
				}
				log.Println("successful rollback")
				return
			}

			// コミット
			if cErr := tx.Commit(); cErr != nil {
				log.Println("commit error: ", cErr)
				c.Error(cErr)
				return
			}
			log.Println("successful committed")
		}()

		// TODO; 依存関係が怪しい
		if err := mysql.SetTxToContext(c, tx); err != nil {
			log.Println("can not set transaction to gin context")
		}
		c.Next() // ハンドラーが実行される
	}
}

// Status Code チェック用のmiddleware
func checkStatusMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		// エラーが発生した場合
		if err := c.Errors.Last(); err != nil {
			code := c.Writer.Status()
			c.JSON(code, createErrorResponse(code))
			return
		}
	}
}
