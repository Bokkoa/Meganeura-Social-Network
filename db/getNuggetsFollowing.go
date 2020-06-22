package db

import (
	"context"
	"fmt"
	"time"

	"github.com/digitalHanzo/Meganeura-Social-Network/models"
	"go.mongodb.org/mongo-driver/bson"
)

//GetNuggetsFollowing - Getting nuggets from people i follow
func GetNuggetsFollowing(ID string, page int) ([]models.GetNuggetsFollow, bool) {
	//avoid long timelapses
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("meganeuradb")
	col := db.Collection("relation")

	skip := (page - 1) * 20

	conditions := make([]bson.M, 0)
	conditions = append(conditions, bson.M{"$match": bson.M{"userid": ID}})

	//like an inner join
	conditions = append(conditions, bson.M{
		"$lookup": bson.M{
			"from":         "nugget",
			"localField":   "userrelationid",
			"foreignField": "userid",
			"as":           "nugget",
		},
	})

	//with undwind let use information (not a master detail)
	conditions = append(conditions, bson.M{"$unwind": "$nugget"})
	//order desc
	conditions = append(conditions, bson.M{"$sort": bson.M{"nugget.date": -1}})
	//pagination
	conditions = append(conditions, bson.M{"$skip": skip})
	//limit
	conditions = append(conditions, bson.M{"$limit": 20})

	pointer, err := col.Aggregate(ctx, conditions)

	var result []models.GetNuggetsFollow

	err = pointer.All(ctx, &result)
	if err != nil {
		fmt.Println(err.Error())
		return result, false
	}

	return result, true

}
