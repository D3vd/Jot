package mongo

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (m Mongo) UpdateNote(noteID string, content string) error {

	// Convert string to primitive ID
	noteObjId, err := primitive.ObjectIDFromHex(noteID)

	if err != nil {
		return err
	}

	filter := bson.M{
		"_id": noteObjId,
	}

	// Update the Note title and content
	update := bson.D{
		{"$set", bson.D{
			{"content", content},
		}},
	}

	updateResult, err := m.DB.Collection("notes").UpdateOne(context.Background(), filter, update)

	if err != nil {
		return err
	}

	if updateResult.ModifiedCount == 0 {
		return errors.New("no document was updated")
	}

	return nil
}
