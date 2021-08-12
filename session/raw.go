package session

import (
	"database/sql"
	"strings"
	"yaorm/clause"
	"yaorm/dialect"
	"yaorm/log"
	"yaorm/schema"
)

type Session struct {
	db      *sql.DB
	sql     strings.Builder
	sqlVars []interface{}

	refTable *schema.Schema
	dialect  dialect.Dialect

	clause clause.Clause
}

func New(db *sql.DB, d dialect.Dialect) *Session {
	return &Session{db: db, dialect: d}
}

func (s *Session) Clear() {
	s.sql.Reset()
	s.sqlVars = nil
	s.clause = clause.Clause{}
}

func (s *Session) DB() *sql.DB {
	return s.db
}

func (s *Session) Raw(sql string, values ...interface{}) *Session {
	s.sql.WriteString(sql)
	s.sqlVars = append(s.sqlVars, values...)

	return s
}

func (s *Session) Exec() (result sql.Result, err error) {
	defer s.Clear()
	s.logSql()

	if result, err = s.db.Exec(s.sql.String(), s.sqlVars...); nil != err {
		log.Error(err)
	}

	return
}

func (s *Session) QueryRow() *sql.Row {
	defer s.Clear()
	s.logSql()

	return s.db.QueryRow(s.sql.String(), s.sqlVars...)
}

func (s *Session) QueryRows() (rows *sql.Rows, err error) {
	defer s.Clear()
	s.logSql()

	if rows, err = s.DB().Query(s.sql.String(), s.sqlVars...); nil != err {
		log.Error(err)
	}
	return
}

func (s *Session) logSql() {
	log.Info(s.sql.String(), s.sqlVars)
}
