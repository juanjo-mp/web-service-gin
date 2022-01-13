package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Painting struct {
	ID     primitive.ObjectID `bson:"_id"`
	Title  *string            `json:"title" validate:"required"`
	Artist *string            `json:"artist" validate:"required"`
	Image  *string            `json:"image" validate:"required"`
}
