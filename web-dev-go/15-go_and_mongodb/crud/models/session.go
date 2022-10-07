package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
)

type SessionData struct {
	SessionId    string `bson:"sessionId" json:"sessionId"`
	LastActivity string `bson:"lastActivity" json:"lastActivity"`
}

func UpdateUserSession(user User) error {
	_, err := userCollection.UpdateOne(context.Background(), bson.M{"_id": user.Id}, bson.M{"$set": bson.M{"sessionData": user.Session}})
	return err
}

func UpdateUserSessionBySessionId(sessionId, lastActivity string) error {
	_, err := userCollection.UpdateOne(context.Background(), bson.M{"sessionData.sessionId": sessionId}, bson.M{"$set": bson.M{"sessionData.sessionId": sessionId, "sessionData.lastActivity": lastActivity}})
	return err
}
