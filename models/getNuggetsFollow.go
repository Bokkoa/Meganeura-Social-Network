package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//GetNuggetsFollowing - Model for  get nuggets of persons that i follow
type GetNuggetsFollowing struct {
	ID             primitive.ObjectID `bson:"_id" json"_id,omitempty"`
	UserID         string             `bson:"userid" json"userId,omitempty"`
	UserRelationID string             `bson:"userrelationid" json"userRelationId,omitempty"`
	Nugget         struct {
		Message string    `bson:"message" json"message,omitempty"`
		Date    time.Time `bson:"date" json"date,omitempty"`
		ID      string    `bson:"_id" json"_id,omitempty"`
	},
}
