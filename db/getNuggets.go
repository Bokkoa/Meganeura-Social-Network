package db

import (
	"context"
	"log"
	"time"

	"github.com/digitalHanzo/Meganeura-Social-Network/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//GetNuggets - return an slice  of nuggets for one user
func GetNuggets(ID string, page int64) ([]*models.BringNuggets, bool) {
	//avoid long timelapses
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	//database
	db := MongoCN.Database("meganeuradb")
	//collection
	col := db.Collection("nugget")

	var results []*models.BringNuggets

	condition := bson.M{
		"userid": ID,
	}

	//setting propertys between find
	options := options.Find()
	//get 20 rows
	options.SetLimit(20)
	//order by date DESC (value -1)
	options.SetSort(bson.D{{Key: "date", Value: -1}})

	//pagination
	options.SetSkip((page - 1) * 20)

	pointer, err := col.Find(ctx, condition, options)

	if err != nil {
		log.Fatal(err.Error())
		return results, false
	}

	//iterating on docs
	for pointer.Next(context.TODO()) {

		//instance of doc
		var row models.BringNuggets

		err := pointer.Decode(&row)

		if err != nil {
			return results, false
		}

		//adding the doc
		results = append(results, &row)
	}

	return results, true
}
