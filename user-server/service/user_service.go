package service

import (
	"github.com/ergildo/go-lang-grcp-app-mysql-example/user-server/database"
	"github.com/ergildo/go-lang-grcp-app-mysql-example/user-server/model"
	_ "github.com/go-sql-driver/mysql"
	"log"
)



func ListAll() []model.User {

	db, err := database.GetDB()
	defer database.CloseDB()
	if err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("select Id, Name from users")
	var users []model.User
	for rows.Next() {
		var u model.User
		rows.Scan(&u.Id, &u.Name)
		users = append(users, u)
	}
	return users
}

func FindById(id int64) model.User {

	db, err := database.GetDB()
	defer database.CloseDB()
	if err != nil {
		log.Fatal(err)
	}
	var u model.User
	db.QueryRow("select Id, Name from users where Id=?", id).Scan(&u.Id, &u.Name)

	return u
}

func Save(user model.User) model.User {

	db, err := database.GetDB()
	defer database.CloseDB()
	if err != nil {
		log.Fatal(err)
	}
	stm, err := db.Prepare("insert into users(Name) values(?)")
	rs, err := stm.Exec(user.Name)
	if err != nil {
		log.Fatal(err)
	}
	lastId, err := rs.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}

	u := FindById(lastId)
	return u
}

func Update(user model.User) model.User {

	db, err := database.GetDB()
	defer database.CloseDB()
	if err != nil {
		log.Fatal(err)
	}
	stm, err := db.Prepare("update users set name =? where id =?")
	_, err = stm.Exec(user.Name, user.Id)
	if err != nil {
		log.Fatal(err)
	}
	u := FindById(user.Id)
	return u
}

func Delete(id int64) {

	db, err := database.GetDB()
	defer database.CloseDB()

	stm, err := db.Prepare("delete from users where id =?")
	_, err = stm.Exec(id)
	if err != nil {
		log.Fatal(err)
	}

}
