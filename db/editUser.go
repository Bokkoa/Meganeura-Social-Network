package db

import (
	"context"
	"time"

	"github.com/digitalHanzo/Meganeura-Social-Network/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//EditUser allows the profile edition
func EditUser(u models.User, ID string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("meganeuradb")
	col := db.Collection("users")

	//map of interface with string index
	//for validation reasons that
	//cant be done on json format
	userMap := make(map[string]interface{})
	if len(u.Name) > 0 {
		userMap["name"] = u.Name
	}
	if len(u.LastName) > 0 {
		userMap["lastName"] = u.LastName
	}

	userMap["birthDate"] = u.BirthDate

	if len(u.Bio) > 0 {
		userMap["bio"] = u.Bio
	}
	if len(u.Banner) > 0 {
		userMap["banner"] = u.Banner
	}
	if len(u.Avatar) > 0 {
		userMap["avatar"] = u.Avatar
	}
	if len(u.Location) > 0 {
		userMap["location"] = u.Location
	}
	if len(u.Web) > 0 {
		userMap["web"] = u.Web
	}

	//updating  string for mongoDB format
	updtString := bson.M{
		"$set": userMap,
	}

	//id of param to ObjID
	objID, _ := primitive.ObjectIDFromHex(ID)

	//get the doc OF the ID WITH "EQUAL"
	filter := bson.M{"_id": bson.M{"$eq": objID}}

	//updating on mongodb
	_, err := col.UpdateOne(ctx, filter, updtString)

	if err != nil {
		return false, err
	}

	return true, nil
}
