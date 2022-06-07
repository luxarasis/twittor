package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

/* User el el modelo de usuario en la base de datos MongoDB */
type User struct {
	ID        primitive.ObjectID `bson: "_id, omitempty" json: "id"`
	Name      string             `bson: "name, omitempty" json: "name"`
	LastName  string             `bson: "lastname, omitempty" json: "lastname"`
	Birthday  time.Time          `bson: "birthday, omitempty" json: "birthday"`
	Email     string             `bson: "email" json: "email"`
	Password  string             `bson: "password, omitempty" json: "password"`
	Avatar    string             `bson: "avatar, omitempty" json: "avatar"`
	Banner    string             `bson: "banner, omitempty" json: "banner"`
	Biography string             `bson: "biography, omitempty" json: "biography"`
	Location  string             `bson: "location, omitempty" json: "location"`
	WebSite   string             `bson: "website, omitempty" json: "website"`
}
