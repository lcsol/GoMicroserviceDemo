package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// A Lab represents lab metadata
type Lab struct {
	ID          primitive.ObjectID `bson:"_id,omitempty"`
	Name        string             `bson:"name,omitempty"`
	Description string             `bson:"description,omitempty"`
	CreatedOn   time.Time          `bson:"createdon,omitempty"`
}
