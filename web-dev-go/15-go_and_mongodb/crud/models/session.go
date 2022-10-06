package models

type SessionData struct {
	SessionId    string `bson:"sessionId" json:"sessionId"`
	LastActivity string `bson:"lastActivity" json:"lastActivity"`
}
