package db

import (
	"context"
	"time"

	"github.com/digitalHanzo/Meganeura-Social-Network/models"
)

//InsertRelation - For NXM Connect
func InsertRelation(t models.Relation) (bool, error) {
	//avoid long timelapses
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	//database
	db := MongoCN.Database("meganeuradb")
	//collection
	col := db.Collection("relation")

	_, err := col.InsertOne(ctx, t)
	if err != nil {
		return false, err
	}

	return true, nil
}
