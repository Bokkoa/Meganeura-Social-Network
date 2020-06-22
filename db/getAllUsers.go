package db

import (
	"context"
	"fmt"
	"time"

	"github.com/digitalHanzo/Meganeura-Social-Network/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//GetAllUsers - ID -> I, Page -> pagination, search -> filter, typeOfSearch -> itself
func GetAllUsers(ID string, page int64, search string, typeOfSearch string) ([]*models.User, bool) {

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	db := MongoCN.Database("meganeuradb")
	col := db.Collection("users")

	var results []*models.User

	findOptions := options.Find()
	findOptions.SetSkip((page - 1) * 20)
	findOptions.SetLimit(20)

	//No importara si son mayusculas o minusculas con (?!)
	query := bson.M{
		"name": bson.M{"$regex": `(?i)` + search},
	}

	cur, err := col.Find(ctx, query, findOptions)

	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	var found, including bool

	for cur.Next(ctx) {
		//user instance
		var s models.User

		//setting on instance
		err := cur.Decode(&s)

		fmt.Println(s)

		if err != nil {
			fmt.Println(err.Error())
			return results, false
		}

		var r models.Relation
		r.UserID = ID
		//getting string from model ID
		r.UserRelationID = s.ID.Hex()

		//for each user for relation
		including = false

		found, err = GetRelation(r)

		//getting only no relation users
		if typeOfSearch == "new" && found == false {
			including = true
		}

		//getting only following users
		if typeOfSearch == "follow" && found == true {
			including = true
		}

		//exclude self follow
		if r.UserRelationID == ID {
			including = false
		}

		if including == true {
			//clean not important field for get
			s.Password = ""
			s.Bio = ""
			s.Location = ""
			s.Banner = ""
			s.Email = ""
			s.Web = ""

			//adding the user to slice
			results = append(results, &s)
		}
	}

	err = cur.Err()
	if err != nil {
		fmt.Println(err.Error())
		return results, false
	}

	//close cursor
	cur.Close(ctx)
	fmt.Println(results)
	return results, true

}
