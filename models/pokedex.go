package models

import "gopkg.in/mgo.v2/bson"

// Pokedex is define the object of pokedex
type Pokedex struct {
	ID        bson.ObjectId `bson:"_id" json:"id"`
	Name      string        `bson:"name" json:"name"`
	Height    float64       `bson:"height" json:"height"`
	Abilities []string      `bson:"ability" json:"ability"`
}
