package models

type User struct {
	Name   string `bson:"name,omitempty" json:"name"`
	Gender string `bson:"gender,omitempty" json:"gender"`
	Age    int    `bson:"age" json:"age"`
}
