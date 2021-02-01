package models

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// LabCollection represent a mongodb session with a lab data model
type LabCollection struct {
	coll *mongo.Collection
}

// NewLabCollection creates a new LabCollection
func NewLabCollection(coll *mongo.Collection) *LabCollection {
	return &LabCollection{coll}
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
func (labs *LabCollection) CreateLab(lab Lab) (*mongo.InsertOneResult, error) {
	return labs.coll.InsertOne(context.TODO(), lab)
}

// UpdateLabName updates the name of a lab
func (labs *LabCollection) UpdateLabName(id string, newName string) (*bson.M, error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	update := bson.D{{"$set", bson.D{{"name", newName}}}}
	var updatedDoc bson.M
	err = labs.coll.FindOneAndUpdate(context.TODO(), bson.M{"_id": ID}, update).Decode(&updatedDoc)
	if err != nil {
		return nil, err
	}
	return &updatedDoc, err
}

// DeleteLab deletes a lab from database
func (labs *LabCollection) DeleteLab(id string) (*mongo.DeleteResult, error) {
	ID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	return labs.coll.DeleteOne(context.TODO(), bson.M{"_id": ID})
}
