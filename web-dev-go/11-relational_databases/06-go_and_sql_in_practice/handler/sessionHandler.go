package handler

import (
	"log"
	"net/http"
	"time"

	"github.com/udemy-go/web-dev-go/11-relational_databases/06-go_and_sql_in_practice/model"
)

func RevokeSession(c *http.Cookie) *http.Cookie {
	c.MaxAge = -1

	go func(v string) {
		err := model.DeleteSession(c.Value)

		if err != nil {
			log.Println(err)
		}
	}(c.Value)

	c.Value = ""

	return c
}

func StillAuthenticated(s string) bool {
	session, err := model.GetSessionById(s)
	if err != nil {
		log.Println(err)
	}

	st, err := time.Parse(time.Layout, session.LastActivity)
	if err != nil {
		log.Println(err)
	}

	return time.Since(st) < time.Minute
}
