package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB

func init() {
	Dbb, err := sql.Open("mysql", "root:password@tcp(127.0.0.1:3307)/coba_golang_1")
	check(err)

	Db = Dbb

	// defer Db.Close()
	createUserTable()
	createSessionTable()
}

func createUserTable() {
	stmt, err := Db.Prepare("CREATE TABLE users (Id int NOT NULL AUTO_INCREMENT, Username varchar(255), Password varchar(255), FirstName varchar(255), LastName varchar(255), Role varchar(255), Age int, PRIMARY KEY (Id))")
	check(err)

	res, err := stmt.Exec()
	if err != nil {
		log.Println("Table already created")
		return
	}

	n, err := res.RowsAffected()
	check(err)

	fmt.Println("Create user table success:", n, "rows affected")
}

func createSessionTable() {
	stmt, err := Db.Prepare("CREATE TABLE sessions (Id varchar(255) NOT NULL, User_id int NOT NULL, LastActivity varchar(255), PRIMARY KEY (Id), FOREIGN KEY (User_id) REFERENCES users(Id))")
	check(err)

	res, err := stmt.Exec()
	if err != nil {
		log.Println("Table already created")
		return
	}

	n, err := res.RowsAffected()
	check(err)

	fmt.Println("Create session table success:", n, "rows affected")
}

func check(err error) {
	if err != nil {
		log.Println(err)
	}
}
