package mongo

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m Mongo) DeleteNote(noteID string) error {
	// Convert string to primitive ID
	noteObjId, err := primitive.ObjectIDFromHex(noteID)

	if err != nil {
		return err
	}

	filter := bson.M{
		"_id": noteObjId,
	}

	_, err = m.DB.Collection("notes").DeleteOne(context.Background(), filter)

	if err != nil {
		return err
	}

	return nil
}
