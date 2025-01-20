package models

type Order struct {
	Id    string  `json:"id" bson:"id"`
	Product string `json:"product" bson:"product"`
	UserId  string `json:"userid" bson:"userid"`
}


