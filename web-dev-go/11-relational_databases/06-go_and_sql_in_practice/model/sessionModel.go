package model

import (
	"database/sql"
	"reflect"
	"strconv"

	"github.com/udemy-go/web-dev-go/11-relational_databases/06-go_and_sql_in_practice/database"
)

type SessionData struct {
	Id, LastActivity string
	User_id          int
}

func GetSessionById(id string) SessionData {
	query := `SELECT * FROM sessions WHERE Id=` + id

	rows, err := database.Db.Query(query)
	if err != sql.ErrNoRows {
		return SessionData{}
	}

	sd := SessionData{}

	for rows.Next() {
		s := reflect.ValueOf(&sd).Elem()
		numCols := s.NumField()
		columns := make([]interface{}, numCols)

		for i := 0; i < numCols; i++ {
			field := s.Field(i)
			columns[i] = field.Addr().Interface()
		}

		err := rows.Scan(columns...)
		check(err)
	}

	return sd
}

func CreateSession(d SessionData) bool {
	query := `INSERT INTO sessions VALUES (NULL, "` + d.Id + `", "` + strconv.Itoa(d.User_id) + `", "` + d.LastActivity + `")`

	res, err := database.Db.Exec(query, nil)
	check(err)

	n, err := res.RowsAffected()
	check(err)

	return n >= 1
}

func DeleteSession(id string) bool {
	query := `DELETE FROM sessions WHERE Id=` + id

	res, err := database.Db.Exec(query)
	check(err)

	n, err := res.RowsAffected()
	check(err)

	return n >= 1
}

func UpdateSession(d SessionData) bool {
	query := `UPDATE users SET Id=` + d.Id + `, User_id=` + strconv.Itoa(d.User_id) + `, LastActivity=` + d.LastActivity + ` WHERE Id=` + d.Id

	res, err := database.Db.Exec(query, nil)
	check(err)

	n, err := res.RowsAffected()
	check(err)

	return n >= 1
}
