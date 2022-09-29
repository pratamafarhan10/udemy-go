package database

import (
	"net/http"
	"time"
)

type User struct {
	Username, Password, FirstName, LastName, Role string
	Age                                           int
}

type SessionData struct {
	Username     string
	LastActivity time.Time
}

var DbUser = map[string]User{}
var DbSession = map[string]SessionData{}

// var DbSessionCleaned time.Time
var SessionLength int = 30

func CleanSession() {
	for k, v := range DbSession {
		if time.Since(v.LastActivity) > time.Second*30 {
			delete(DbSession, k)
		}
	}
	// DbSessionCleaned = time.Now()
}

func DeleteSession(c *http.Cookie) *http.Cookie {
	c.MaxAge = -1

	delete(DbSession, c.Value)

	c.Value = ""

	return c
}
