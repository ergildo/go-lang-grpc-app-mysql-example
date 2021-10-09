package database

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/go-sql-driver/mysql"
)

var (
	cfg = mysql.Config{
		User:   "root",
		Passwd: "root@123",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "userdb",
	}
	db *sql.DB
)

func GetDB() (*sql.DB, error) {
	newDb, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return nil, errors.New(fmt.Sprintf("could not open a database connection: %v", err))
	}
	db = newDb
	return db, nil
}

func CloseDB() {
	db.Close()
}
