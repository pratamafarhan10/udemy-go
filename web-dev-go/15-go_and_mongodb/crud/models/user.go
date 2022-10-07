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

func (user User) InsertUser() error {
	_, err := userCollection.InsertOne(context.Background(), user)
	return err
}

func (user User) GetUser(filter bson.M, dst *User) error {
	res := userCollection.FindOne(context.Background(), filter)
	return user.DecodeData(res, dst)
}

func (user User) GetUserByUsername(dst *User) error {
	res := userCollection.FindOne(context.Background(), bson.M{"username": user.Username})
	return user.DecodeData(res, dst)
}

func (user User) GetUserBySessionId(dst *User) error {
	res := userCollection.FindOne(context.Background(), bson.M{"sessionData.sessionId": user.Session.SessionId})
	return user.DecodeData(res, dst)
}

func (user User) UpdateUser() error {
	_, err := userCollection.UpdateOne(context.Background(), bson.M{"_id": user.Id}, bson.M{"$set": bson.M{"username": user.Username, "password": user.Password, "firstname": user.FirstName, "lastname": user.LastName, "role": user.Role, "age": user.Age, "profilePicture": user.ProfilePicture}})
	return err
}

func (user User) UsernameAlreadyTaken() (bool, error) {
	dst := User{
		Username: user.Username,
	}

	err := user.GetUserByUsername(&dst)
	return err != mongo.ErrNoDocuments, err
}

func (user User) DecodeData(res *mongo.SingleResult, dst *User) error {
	err := res.Decode(&dst)
	if err != nil {
		log.Println(err == mongo.ErrNoDocuments)
		log.Println("ERROR DECODE GETUSER")
		return err
	}
	return res.Err()
}
