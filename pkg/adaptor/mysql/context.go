package mysql

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/gin-gonic/gin"
)

// TODO: このファイルはここに置いていいのか？

type key string

const (
	txKey key = "transaction"
)

// ginのコンテキストにトランザクションを格納
func SetTxToContext(c context.Context, tx *sql.Tx) error {
	gc, err := ginContextFromContext(c)
	if err != nil {
		return err
	}
	gc.Set(string(txKey), tx)
	return nil
}

// ginのコンテキストからトランザクションを抽出
func txFromContext(c context.Context) (*sql.Tx, error) {
	gc, err := ginContextFromContext(c)
	if err != nil {
		return nil, err
	}

	v, exists := gc.Get(string(txKey))
	if !exists {
		return nil, fmt.Errorf("can not get gin context value for transaction")
	}

	tx, ok := v.(*sql.Tx)
	if !ok {
		return nil, fmt.Errorf("can not convert value type to *sql.Tx")
	}
	return tx, nil
}

// https://gqlgen.com/master/recipes/gin/
// context.Contextから*gin.Contextに変換
func ginContextFromContext(ctx context.Context) (*gin.Context, error) {
	gc, ok := ctx.(*gin.Context)
	if !ok {
		err := fmt.Errorf("gin.Context has wrong type")
		return nil, err
	}
	return gc, nil
}
