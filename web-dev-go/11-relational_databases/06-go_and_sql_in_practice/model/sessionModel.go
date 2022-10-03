package model

import (
	"reflect"

	"github.com/udemy-go/web-dev-go/11-relational_databases/06-go_and_sql_in_practice/database"
)

type SessionData struct {
	Id           string
	User_id      int
	LastActivity string
}

func GetSessionById(id string) (SessionData, error) {
	query := `SELECT * FROM sessions WHERE Id=?`

	rows, err := database.Db.Query(query, id)
	if err != nil {
		return SessionData{}, err
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
		if err != nil {
			return SessionData{}, err
		}
	}

	return sd, nil
}

func CreateSession(d SessionData) error {
	query := `INSERT INTO sessions VALUES (?, ?, ?)`

	res, err := database.Db.Exec(query, d.Id, d.User_id, d.LastActivity)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}

func DeleteSession(id string) error {
	query := `DELETE FROM sessions WHERE Id=?`

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

func UpdateSession(d SessionData) error {
	query := `UPDATE sessions SET Id = ?, User_id = ?, LastActivity = ? WHERE Id=?`

	res, err := database.Db.Exec(query, d.Id, d.User_id, d.LastActivity, d.Id)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}

	return nil
}
