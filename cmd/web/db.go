package main

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	maxOpenDbConn = 25
	maxIdleConn   = 25
	maxDBLifetime = 5 * time.Minute
)

func initMySQLDB() (*sql.DB, error) {
	db, err := sql.Open("mysql", os.Getenv("mysql_dsn"))
	if err != nil {
		return nil, err
	}

	// test our database
	if err = db.Ping(); err != nil {
		return nil, err
	}
	db.SetMaxOpenConns(maxOpenDbConn)
	db.SetMaxIdleConns(maxIdleConn)
	db.SetConnMaxLifetime(maxDBLifetime)
	return db, nil
}
