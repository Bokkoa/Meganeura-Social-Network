package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

//omitempty = not null
//bson = how look in db
//json = how it is show

//must be not a space between comma and omitempty

// User - Model
type User struct {
	ID        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name,omitempty"`
	LastName  string             `bson:"lastName" json:"lastName,omitempty"`
	BirthDate time.Time          `bson:"birthDate" json:"birthDate,omitempty"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password,omitempty"`
	Avatar    string             `bson:"avatar" json:"avatar,omitempty"`
	Banner    string             `bson:"banner" json:"banner,omitempty"`
	Bio       string             `bson:"bio" json:"bio,omitempty"`
	Location  string             `bson:"location" json:"location,omitempty"`
	Web       string             `bson:"web" json:"web,omitempty"`
}
