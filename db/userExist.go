package db

import (
	"context"
	"time"

	"github.com/digitalHanzo/Meganeura-Social-Network/models"
	"go.mongodb.org/mongo-driver/bson"
)

//CheckUserExist - Get email for coincidences
func CheckUserExist(email string) (models.User, bool, string) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	//stop the context when finish
	defer cancel()

	db := MongoCN.Database("meganeuradb")
	col := db.Collection("users")

	//maping condition on bson format
	condition := bson.M{"email": email}

	var result models.User

	//receive context and bson condition
	err := col.FindOne(ctx, condition).Decode(&result)

	//storing ID as hexadecimal on result
	//for the third return param
	ID := result.ID.Hex()

	//IF ID is empty it show a empty string
	if err != nil {
		return result, false, ID
	}

	return result, true, ID
}
