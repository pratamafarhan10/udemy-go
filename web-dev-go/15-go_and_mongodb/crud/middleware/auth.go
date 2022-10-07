package auth

import (
	"fmt"
	"net/http"
	"time"

	"github.com/udemy-go/web-dev-go/15-go_and_mongodb/crud/models"
	"go.mongodb.org/mongo-driver/bson"
)

func Authenticate(r *http.Request) (models.User, *http.Cookie, error) {
	// Check cookie
	c, err := r.Cookie("session")
	if err != nil {
		return models.User{}, nil, http.ErrNoCookie
	}

	// Get user
	user, err := GetUserBySessionId(c)
	if err != nil {
		return models.User{}, c, err
	}

	t, err := time.Parse(time.Layout, user.Session.LastActivity)
	if err != nil {
		return models.User{}, c, err
	}

	if time.Since(t) > 30*time.Second {
		return models.User{}, c, fmt.Errorf("%v", "session timeout")
	}

	// Update last activity
	err = UpdateLastActivity(c.Value)
	if err != nil {
		return models.User{}, c, err
	}

	return user, c, nil
}

func GetUserBySessionId(c *http.Cookie) (models.User, error) {
	user := models.User{}
	user.Session.SessionId = c.Value

	// err := user.GetUser(bson.M{"sessionData.sessionId": c.Value}, &user)
	err := user.GetUser(bson.M{"sessionData.sessionId": user.Session.SessionId}, &user)
	return user, err
}

func UpdateLastActivity(sessionId string) error {
	err := models.UpdateUserSessionBySessionId(sessionId, time.Now().Format(time.Layout))

	return err
}
