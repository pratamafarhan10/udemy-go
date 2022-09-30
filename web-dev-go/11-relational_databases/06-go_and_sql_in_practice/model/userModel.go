package model

import (
	"database/sql"
	"reflect"
	"strconv"

	"github.com/udemy-go/web-dev-go/11-relational_databases/06-go_and_sql_in_practice/database"
)

type User struct {
	Username, Password, FirstName, LastName, Role string
	Id, Age                                       int
}

func CreateUser(user User) bool {
	query := `INSERT INTO users VALUES (NULL, ?, ?, ?, ?, ?, ?)`
	res, err := database.Db.Exec(query, user.Username, user.Password, user.FirstName, user.LastName, user.Role, user.Age)
	check(err)

	n, err := res.RowsAffected()
	check(err)

	return n >= 1
}

func GetUserById(id int) User {
	query := `SELECT * FROM users WHERE Id=?`

	rows, err := database.Db.Query(query, id)
	if err != sql.ErrNoRows {
		return User{}
	}

	user := User{}
	for rows.Next() {

		s := reflect.ValueOf(&user).Elem()
		numCols := s.NumField()
		columns := make([]interface{}, numCols)

		for i := 0; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}

		err = rows.Scan(columns...)
		check(err)
	}

	return user
}

func GetUserByUsername(username string) User {
	query := `SELECT * FROM users WHERE Username=?`

	rows, err := database.Db.Query(query, username)

	if err != sql.ErrNoRows {
		return User{}
	}

	user := User{}
	for rows.Next() {

		s := reflect.ValueOf(&user).Elem()
		numCols := s.NumField()
		columns := make([]interface{}, numCols)

		for i := 0; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}

		err = rows.Scan(columns...)
		check(err)
	}

	return user
}

func UpdateUser(user User) bool {
	query := `UPDATE users 
	SET Username=` + user.Username + `, Password=` + user.Password + `, FirstName=` + user.FirstName + `, LastName=` + user.LastName + `, Role=` + user.Role + `, Age=` + strconv.Itoa(user.Age) + ` 
	WHERE Id=` + strconv.Itoa(user.Id)

	res, err := database.Db.Exec(query, nil)
	check(err)

	n, err := res.RowsAffected()
	check(err)

	return n >= 1
}

func DeleteUser(id int) bool {
	query := `DELETE FROM users WHERE Id=` + strconv.Itoa(id)

	res, err := database.Db.Exec(query, nil)
	check(err)

	n, err := res.RowsAffected()
	check(err)

	return n >= 1
}
