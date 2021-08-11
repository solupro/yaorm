package session

import (
	"database/sql"
	"os"
	"testing"
	"yaorm/dialect"

	_ "github.com/mattn/go-sqlite3"
)

type User struct {
	Id   int64 `yaorm:"PRIMARY KEY"`
	Name string
	Age  int
}

var (
	TestDB      *sql.DB
	TestDial, _ = dialect.GetDialect("sqlite3")
)

func TestMain(m *testing.M) {
	TestDB, _ = sql.Open("sqlite3", "../data.db")
	code := m.Run()
	_ = TestDB.Close()
	os.Exit(code)
}

func NewSession() *Session {
	return New(TestDB, TestDial)
}

func TestSession_CreateTable(t *testing.T) {
	session := NewSession().Model(&User{})
	session.DropTable()
	session.CreateTable()
	if !session.HasTable() {
		t.Fatal("create table error")
	}
}
