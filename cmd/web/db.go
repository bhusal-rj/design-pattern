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
	dsn := os.Getenv("mysql_dsn")
	if len(dsn) < 5 {
		dsn = "root:hello@tcp(localhost:3306)/breeders?parseTime=true&tls=false"
	}
	db, err := sql.Open("mysql", dsn)
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
