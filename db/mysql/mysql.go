package mysql

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	user string
	pass string
	name string
}

var db *sql.DB

func DBInit(c Config) error {
	var err error
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", c.user, c.pass, c.name))
	if err != nil {
		return err
	}
	err = db.Ping()
	if err != nil {
		return err
	}
	return nil
}

func Query(qs string, args ...interface{}) (*sql.Rows, error) {
	tx, err := db.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(qs)
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(args...)
	if err != nil {
		return nil, err
	}

	err = stmt.Close()
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		return nil, err
	}

	return rows, nil
}
