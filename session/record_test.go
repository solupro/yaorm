package session

import "testing"

var (
	user1 = &User{Id: 1, Name: "solu", Age: 32}
	user2 = &User{Id: 2, Name: "Tom", Age: 28}
	user3 = &User{Id: 3, Name: "zhangsan", Age: 45}
)

func testRecordInit(t *testing.T) *Session {
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
	s := testRecordInit(t)
	aff, err := s.Insert(user3)
	if nil != err || 1 != aff {
		t.Fatal("insert user error")
	}
}

func TestSession_Find(t *testing.T) {
	s := testRecordInit(t)
	var users []User
	if err := s.Find(&users); nil != err || 2 != len(users) {
		t.Fatal("find error")
	}
}

func TestSession_Limit(t *testing.T) {
	s := testRecordInit(t)
	var users []User
	err := s.Limit(1).Find(&users)
	if err != nil || len(users) != 1 {
		t.Fatal("failed to query with limit condition")
	}
}

func TestSession_Update(t *testing.T) {
	s := testRecordInit(t)
	affected, _ := s.Where("Name = ?", "Tom").Update("Age", 34)
	u := &User{}
	_ = s.OrderBy("Age DESC").First(u)

	if affected != 1 || u.Age != 34 {
		t.Fatal("failed to update")
	}
}

func TestSession_DeleteAndCount(t *testing.T) {
	s := testRecordInit(t)
	affected, _ := s.Where("Name = ?", "Tom").Delete()
	count, _ := s.Count()

	if affected != 1 || count != 1 {
		t.Fatal("failed to delete or count")
	}
}
