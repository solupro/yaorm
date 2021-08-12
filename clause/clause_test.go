package clause

import (
	"reflect"
	"testing"
)

func testSelect(t *testing.T) {
	var clause Clause
	clause.Set(LIMIT, 3)
	clause.Set(SELECT, "User", []string{"*"})
	clause.Set(WHERE, "Name = ?", "solu")
	clause.Set(ORDER_BY, "age asc")
	sql, vars := clause.Build(SELECT, WHERE, ORDER_BY, LIMIT)
	t.Log(sql, vars)

	if sql != "select * from User where Name = ? order by age asc limit ?" {
		t.Fatal("build sql error")
	}
	if !reflect.DeepEqual(vars, []interface{}{"solu", 3}) {
		t.Fatal("buil vars error")
	}
}

func TestClause_Build(t *testing.T) {
	t.Run("select", func(t *testing.T) {
		testSelect(t)
	})
}
