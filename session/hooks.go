package session

import (
	"reflect"
	"yaorm/log"
)

const (
	BeforeInsert = "BeforeInsert"
	AfterInsert  = "AfterInsert"
	BeforeDelete = "BeforeDelete"
	AfterDelete  = "AfterDelete"
	BeforeUpdate = "BeforeUpdate"
	AfterUpdate  = "AfterUpdate"
	BeforeQuery  = "BeforeQuery"
	AfterQuery   = "AfterQuery"
)

func (s *Session) CallMethod(method string, value interface{}) {
	fm := reflect.ValueOf(s.RefTable().Model).MethodByName(method)
	if nil != value {
		fm = reflect.ValueOf(value).MethodByName(method)
	}

	param := []reflect.Value{reflect.ValueOf(s)}
	if fm.IsValid() {
		if v := fm.Call(param); len(v) > 0 {
			if err, ok := v[0].Interface().(error); ok {
				log.Error(err)
			}
		}
	}

	return
}
