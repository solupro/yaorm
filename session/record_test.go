package session

import "testing"

var (
	user1 = &User{Id: 1, Name: "solu", Age: 32}
	user2 = &User{Id: 2, Name: "Tom", Age: 28}
	user3 = &User{Id: 3, Name: "zhangsan", Age: 45}
)

func testRecordInt(t *testing.T) *Session {
	t.Helper()
	s := NewSession().Model(&User{})
	err1 := s.DropTable()
	err2 := s.CreateTable()
	_, err3 := s.Insert(user1, user2)
	if nil != err1 || nil != err2 || nil != err3 {
		t.Fatal("init test record error")
	}

	return s
}

func TestSession_Insert(t *testing.T) {
	s := testRecordInt(t)
	aff, err := s.Insert(user3)
	if nil != err || 1 != aff {
		t.Fatal("insert user error")
	}
}

func TestSession_Find(t *testing.T) {
	s := testRecordInt(t)
	var users []User
	if err := s.Find(&users); nil != err || 2 != len(users) {
		t.Fatal("find error")
	}
}
