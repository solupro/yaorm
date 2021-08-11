package yaorm

import (
	"database/sql"
	"yaorm/dialect"
	"yaorm/log"
	"yaorm/session"

	_ "github.com/mattn/go-sqlite3"
)

type Engine struct {
	db      *sql.DB
	dialect dialect.Dialect
}

func NewEngine(driver, source string) (e *Engine, err error) {
	db, err := sql.Open(driver, source)
	if nil != err {
		log.Error(err)
		return
	}

	if err = db.Ping(); nil != err {
		log.Error(err)
		return
	}

	dial, ok := dialect.GetDialect(driver)
	if !ok {
		log.Error("dialect %s not found", driver)
		return
	}

	e = &Engine{db: db, dialect: dial}
	log.Info("Connect database success")
	return
}

func (e *Engine) Close() {
	if err := e.db.Close(); nil != err {
		log.Error("Failed to close database:" + err.Error())
	}
	log.Info("Close database success")
}

func (e *Engine) NewSession() *session.Session {
	return session.New(e.db, e.dialect)
}
