package session

import "yaorm/log"

func (s *Session) Begin() (err error) {
	log.Info("transaction begin")

	if s.tx, err = s.db.Begin(); nil != err {
		log.Error(err)
		return
	}
	return
}

func (s *Session) Commit() (err error) {
	log.Info("transaction commit")

	if err := s.tx.Commit(); nil != err {
		log.Error(err)
	}
	return
}

func (s *Session) Rollback() (err error) {
	log.Info("transaction rollback")

	if err := s.tx.Rollback(); nil != err {
		log.Error(err)
	}
	return
}
