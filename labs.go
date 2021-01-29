package main

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

// LabCollection represent a mongodb session with a lab data model
type LabCollection struct {
	coll *mongo.Collection
}

// GetAll retrieves all labs data
func (labs *LabCollection) GetAll() ([]Lab, error) {
	ctx := context.TODO()
	list := []Lab{}

	cursor, err := labs.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(ctx, &list)
	if err != nil {
		return nil, err
	}
	return list, err
}

// GetByID retrieves a single lab by id

// CreateLab inserts a lab into database

// UpdateLabName updates the name of a lab

// DeleteLab deletes a lab from database
