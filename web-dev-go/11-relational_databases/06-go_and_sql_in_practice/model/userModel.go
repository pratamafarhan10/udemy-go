package model

import (
	"log"
	"reflect"

	"github.com/udemy-go/web-dev-go/11-relational_databases/06-go_and_sql_in_practice/database"
)

type User struct {
	Id                                            int
	Username, Password, FirstName, LastName, Role string
	Age                                           int
}

func CreateUser(user User) error {
	query := `INSERT INTO users VALUES (NULL, ?, ?, ?, ?, ?, ?)`
	res, err := database.Db.Exec(query, user.Username, user.Password, user.FirstName, user.LastName, user.Role, user.Age)
	if err != nil {
		log.Panicln(err)
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		log.Panicln(err)
		return err
	}

	return err
}

func GetUserById(id int) (User, error) {
	query := `SELECT * FROM users WHERE Id=?`

	rows, err := database.Db.Query(query, id)
	if err != nil {
		return User{}, err
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
		if err != nil {
			return User{}, err
		}
	}

	return user, nil
}

func GetUserByUsername(username string) (User, error) {
	query := `SELECT * FROM users WHERE Username=?`

	rows, err := database.Db.Query(query, username)

	if err != nil {
		return User{}, err
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
		if err != nil {
			return User{}, err
		}
	}

	return user, nil
}

func UpdateUser(user User) error {
	query := `UPDATE users 
	SET Username=?, Password=?, FirstName=?, LastName=?, Role=?, Age=? 
	WHERE Id=?`

	res, err := database.Db.Exec(query, user.Username, user.Password, user.FirstName, user.LastName, user.Role, user.Age, user.Id)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func DeleteUser(id int) error {
	query := `DELETE FROM users WHERE Id=?`

	res, err := database.Db.Exec(query, id)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
