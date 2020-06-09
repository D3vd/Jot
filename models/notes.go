package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Notes struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Content string             `json:"content"`
	Created time.Time          `json:"created"`
}
