package db

import (
	"context"
	"time"

	"github.com/digitalHanzo/Meganeura-Social-Network/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//InsertNugget - User is creating a nugget
func InsertNugget(n models.RecordNugget) (string, bool, error) {
	//avoid long timelapses
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	//database
	db := MongoCN.Database("meganeuradb")
	//collection
	col := db.Collection("nugget")

	record := bson.M{
		"userid":  n.UserID,
		"message": n.Message,
		"date":    n.Date,
	}

	result, err := col.InsertOne(ctx, record)

	if err != nil {
		return "", false, err
	}

	objID, _ := result.InsertedID.(primitive.ObjectID)

	//same to HEX
	return objID.String(), true, nil

}
