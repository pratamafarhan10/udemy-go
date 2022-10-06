package models

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type User struct {
	Id             primitive.ObjectID `bson:"_id" json:"_id"`
	Username       string             `bson:"username" json:"username"`
	Password       string             `bson:"password" json:"password"`
	FirstName      string             `bson:"firstname" json:"firstname"`
	LastName       string             `bson:"lastname" json:"lastname"`
	Role           string             `bson:"role" json:"role"`
	Age            string             `bson:"age" json:"age"`
	ProfilePicture string             `bson:"profilePicture" json:"profilePicture"`
	Session        SessionData        `bson:"sessionData" json:"sessionData"`
}

func InsertUser(user User) error {
	_, err := userCollection.InsertOne(context.Background(), user)
	return err
}

func GetUserByUsername(user User, container *User) error {
	res := userCollection.FindOne(context.Background(), bson.M{"username": user.Username})
	return DecodeData(res, container)
}

func GetUserBySessionId(user User, container *User) error {
	res := userCollection.FindOne(context.Background(), bson.M{"sessionData.sessionId": user.Session.SessionId})
	return DecodeData(res, container)
}

func UpdateUser(user User) error {
	_, err := userCollection.UpdateOne(context.Background(), bson.M{"_id": user.Id}, bson.M{"$set": bson.M{"username": user.Username, "password": user.Password, "firstname": user.FirstName, "lastname": user.LastName, "role": user.Role, "age": user.Age, "profilePicture": user.ProfilePicture}})
	return err
}

func UpdateUserSession(user User) error {
	_, err := userCollection.UpdateOne(context.Background(), bson.M{"_id": user.Id}, bson.M{"$set": bson.M{"sessionData": user.Session}})
	return err
}

func UpdateUserSessionBySessionId(sessionId, lastActivity string) error {
	_, err := userCollection.UpdateOne(context.Background(), bson.M{"sessionData.sessionId": sessionId}, bson.M{"$set": bson.M{"sessionData.sessionId": sessionId, "sessionData.lastActivity": lastActivity}})
	return err
}

func DecodeData(res *mongo.SingleResult, container *User) error {
	err := res.Decode(&container)
	if err != nil {
		log.Println(err == mongo.ErrNoDocuments)
		log.Println("ERROR DECODE GETUSER")
		return err
	}
	return res.Err()
}
