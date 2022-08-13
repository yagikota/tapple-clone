package middleware

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	txKey = "transaction"
)

// トランザクション開始
func beginTxAndSetToContext(c *gin.Context, conn *sql.DB) (*sql.Tx, error) {
	tx, err := conn.Begin()
	if err != nil {
		return nil, err
	}
	log.Println("begin database transaction")
	return tx, nil
}

func statusInList(status int, statusCodeList []int) bool {
	for _, code := range statusCodeList {
		if code == status {
			return true
		}
	}
	return false
}

// トランザクション用のmiddleware
func TransactMiddleware(conn *sql.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		tx, err := beginTxAndSetToContext(c, conn)
		if err != nil {
			c.AbortWithError(http.StatusInternalServerError, err)
			return
		}

		defer func() {
			if r := recover(); r != nil {
				if rbErr := tx.Rollback(); rbErr != nil {
					log.Println("rollback error: ", rbErr)
				}
				panic(r)
			}
		}()

		c.Set(txKey, tx)
		c.Next() // ハンドラーが実行される

		wantStatusCodes := []int{http.StatusOK, http.StatusCreated}
		if statusInList(c.Writer.Status(), wantStatusCodes) {
			log.Println("committing transactions")
			if commitErr := tx.Commit(); commitErr != nil {
				log.Println("trx commit error: ", commitErr)
				panic(commitErr)
			}
		} else {
			log.Println("invalid status code: ", c.Writer.Status())
			panic(fmt.Sprintf("invalid status code: %d", c.Writer.Status()))
		}
	}
}

// func GinContextFromContext(ctx context.Context) (*gin.Context, error) {
// 	ginContext := ctx.Value("GinContextKey")
// 	if ginContext == nil {
// 		err := fmt.Errorf("could not retrieve gin.Context")
// 		return nil, err
// 	}

// 	gc, ok := ginContext.(*gin.Context)
// 	if !ok {
// 		err := fmt.Errorf("gin.Context has wrong type")
// 		return nil, err
// 	}
// 	return gc, nil
// }
