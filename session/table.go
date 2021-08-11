package session

import (
	"fmt"
	"reflect"
	"strings"
	"yaorm/log"
	"yaorm/schema"
)

func (s *Session) Model(v interface{}) *Session {
	if nil == s.refTable || reflect.TypeOf(v) != reflect.TypeOf(s.refTable.Model) {
		s.refTable = schema.Parse(v, s.dialect)
	}

	return s
}

func (s *Session) RefTable() *schema.Schema {
	if nil == s.refTable {
		log.Error("Model is not set")
	}

	return s.refTable
}

func (s *Session) CreateTable() error {
	table := s.refTable
	var columns []string
	for _, field := range table.Fields {
		columns = append(columns, fmt.Sprintf("%s %s %s", field.Name, field.Type, field.Tag))
	}
	desc := strings.Join(columns, ",")
	_, err := s.Raw(fmt.Sprintf("create table %s (%s)", table.Name, desc)).Exec()
	return err
}

func (s *Session) DropTable() error {
	_, err := s.Raw(fmt.Sprintf("drop table if exists %s", s.RefTable().Name)).Exec()
	return err
}

func (s *Session) HasTable() bool {
	sql, values := s.dialect.TableExistsSQL(s.RefTable().Name)
	row := s.Raw(sql, values...).QueryRow()
	var tmp string
	_ = row.Scan(&tmp)
	return tmp == s.RefTable().Name
}
