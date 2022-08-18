package mysql

import (
	"context"
	"database/sql"
	"fmt"
	"testing"
	"time"

	"github.com/CyberAgentHack/2208-ace-go-server/pkg/domain/entity"
	"github.com/CyberAgentHack/2208-ace-go-server/pkg/infra"
	"github.com/google/go-cmp/cmp"

	_ "github.com/go-sql-driver/mysql"
)

var (
	userID = 1
)

// テスト用に DB接続(テスト用のデータベースつく)
func NewMySQLConnectorForTest(t *testing.T) *infra.MySQLConnector {
	t.Helper()

	driverName := "mysql"
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=Local",
		"root",
		"",
		"localhost:3306",
		"tapple_c")
	conn, err := sql.Open(driverName, dsn)
	if err != nil {
		t.Fatal(err)
	}

	t.Cleanup(
		func() { _ = conn.Close() },
	)

	return &infra.MySQLConnector{Conn: conn}
}

func prepareUser() *entity.User {
	return &entity.User{
		ID:       1,
		Name:     "カイ",
		Icon:     " male/n000029/main_0001_01.jpg",
		Gender:   0,
		Birthday: time.Date(2000, 9, 7, 0, 0, 0, 0, time.Local),
		Location: 34,
	}
}

func Test_userRepository_FindUserByUserID(t *testing.T) {
	ctx := context.TODO()

	// トランザクション開始
	conn := NewMySQLConnectorForTest(t).Conn
	tx, err := conn.BeginTx(ctx, nil)
	t.Cleanup(
		func() { _ = tx.Rollback() },
	)
	if err != nil {
		t.Fatal(err)
	}

	wantUser := prepareUser()
	ur := NewUserRepository(conn)
	gotUser, err := ur.FindUserByUserID(ctx, userID)
	if err != nil {
		t.Fatal(err)
	}
	if d := cmp.Diff(wantUser, gotUser); len(d) != 0 {
		t.Fatal(err)
	}
}
