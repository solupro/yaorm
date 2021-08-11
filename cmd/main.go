package main

import (
	"yaorm"
	"yaorm/log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	engine, _ := yaorm.NewEngine("sqlite3", "data.db")
	defer engine.Close()

	s := engine.NewSession()
	_, _ = s.Raw("drop table if exists User;").Exec()
	_, _ = s.Raw("create table User(name text);").Exec()
	_, _ = s.Raw("create table User(name text);").Exec()
	result, _ := s.Raw("insert into User(`name`) values (?), (?)", "solu", "Sam").Exec()
	count, _ := result.RowsAffected()
	log.Infof("Exec success, %d affected\n", count)

}
