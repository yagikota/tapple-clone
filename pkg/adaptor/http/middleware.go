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
func TransactMiddleware(conn *sql.DB) gin.HandlerFunc {
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

			httpStatus := c.Writer.Status()
			if httpStatus == http.StatusOK || httpStatus == http.StatusCreated {
				log.Println("committing transactions")
				if commitErr := tx.Commit(); commitErr != nil {
					log.Println("transactions commit error: ", commitErr)
					panic(commitErr)
				}
				log.Println("transactions successful committed")
			} else {
				var responseCode int
				if httpStatus == http.StatusBadRequest {
					responseCode = http.StatusBadRequest
				} else {
					responseCode = http.StatusInternalServerError
				}
				c.JSON(responseCode, CreateErrorResponse(responseCode))
				if rbErr := tx.Rollback(); rbErr != nil {
					log.Println("rollback error: ", rbErr)
				}
				log.Println("successful rollback")
			}

			// panicした場合
			if r := recover(); r != nil {
				if rbErr := tx.Rollback(); rbErr != nil {
					log.Println("rollback error: ", rbErr)
				}
				log.Println("successful rollback")
			}
		}()

		// TODO; 依存関係が怪しい
		if err := mysql.SetTxToContext(c, tx); err != nil {
			log.Println("can not set transaction to gin context")
			panic(err)
		}

		c.Next() // ハンドラーが実行される

	}
}
