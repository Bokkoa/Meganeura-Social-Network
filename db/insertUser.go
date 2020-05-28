package db

import (
	"context"
	"time"

	"github.com/digitalHanzo/Meganeura-Social-Network/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//InsertUser - Insert User on DB
func InsertUser(u models.User) (string, bool, error) {

	//avoid long timelapses
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	//database
	db := MongoCN.Database("meganeuradb")
	//collection
	col := db.Collection("users")

	u.Password, _ = EncryptPassword(u.Password)

	//inserting (ctx = context avoid timeout )
	result, err := col.InsertOne(ctx, u)

	if err != nil {
		return "", false, err
	}

	//getting mongodb ID
	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	//cast to string
	return ObjID.String(), true, nil

}
