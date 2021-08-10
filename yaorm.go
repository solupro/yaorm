package yaorm

import (
	"database/sql"
	"yaorm/log"
	"yaorm/session"

	_ "github.com/mattn/go-sqlite3"
)

type Engine struct {
	db *sql.DB
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

	e = &Engine{db: db}
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
	return session.New(e.db)
}
