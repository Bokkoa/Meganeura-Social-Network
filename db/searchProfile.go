package db

import (
	"context"
	"fmt"
	"time"

	"github.com/digitalHanzo/Meganeura-Social-Network/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// SearchProfile - by ID
func SearchProfile(ID string) (models.User, error) {

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*15)
	defer cancel()

	db := MongoCN.Database("meganeuradb")
	col := db.Collection("users")

	var profile models.User

	//convert Param ID to ObjectID
	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id": objID,
	}

	err := col.FindOne(ctx, condition).Decode(&profile)
	//nullified pass for security
	profile.Password = ""

	if err != nil {
		fmt.Println("Error obteniendo al usuario" + err.Error())
		return profile, err
	}

	return profile, nil
}
