package schema

import (
	"testing"
	"yaorm/dialect"
)

type User struct {
	Id   int64 `yaorm:"PRIMARY KEY"`
	Name string
	Age  int
}

var TestDialect, _ = dialect.GetDialect("sqlite3")

func TestParse(t *testing.T) {
	scheme := Parse(&User{}, TestDialect)

	if scheme.Name != "User" || len(scheme.Fields) != 3 {
		t.Fatal("failed to parse User")
	}

	if scheme.GetField("Id").Tag != "PRIMARY KEY" {
		t.Fatal("fialed to parse Tag")
	}
}
