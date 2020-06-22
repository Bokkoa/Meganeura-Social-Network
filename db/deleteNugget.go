package db

import (
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//DeleteNugget - For db deletion of nuggets by ID of nugget and User ID
func DeleteNugget(ID string, UserID string) error {

	//avoid long timelapses
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("meganeuradb")
	col := db.Collection("nugget")

	//getting ID as ObjID
	objID, _ := primitive.ObjectIDFromHex(ID)

	condition := bson.M{
		"_id":    objID,
		"userid": UserID,
	}

	_, err := col.DeleteOne(ctx, condition)

	//doesnt matter if err != nil
	return err

}
