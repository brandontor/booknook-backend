package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type BookNook struct {
	BookNookName string             `bson:"booknookname,omitempty" validate:"required"`
	BookTitle    string             `bson:"booktitle,omitempty" validate:"required"`
	Genre        string             `bson:"genre,omitempty" validate:"required"`
	Schedule     string             `bson:"schedule,omitempty" validate:"required"`
	Description  string             `bson:"description,omitempty" validate:"required"`
	BookNook_ID  primitive.ObjectID `bson:"booknook_id,omitempty"`
}
