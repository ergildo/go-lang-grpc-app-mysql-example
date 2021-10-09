package setup

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

func SetUpDB() {

	cfg := mysql.Config{
		User:   "root",
		Passwd: "root@123",
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
	}

	db, err := sql.Open("mysql", cfg.FormatDSN())
	defer db.Close()
	if err != nil {
		log.Fatal("could not connect to database", err)
	}
	log.Println("Creating database userdb...")
	db.Exec("create database if not exists userdb")
	db.Exec("use userdb")
	log.Println("Creating table users...")
	db.Exec(`create table if not exists users (
		id integer auto_increment,
		name varchar(80),
		PRIMARY KEY(id)
	)`)

}
