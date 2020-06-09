package mongo

import (
	"context"
	"github.com/R3l3ntl3ss/Jot/models"
	"go.mongodb.org/mongo-driver/bson"
)

func (m Mongo) GetAllNotes() ([]models.Notes, error) {

	cur, err := m.DB.Collection("notes").Find(context.Background(), bson.D{})

	if err != nil {
		return nil, err
	}

	var results []models.Notes

	// Iterate through the cursor
	for cur.Next(context.TODO()) {
		// Create a value into which the single document can be decoded
		var elem models.Notes
		err := cur.Decode(&elem)
		if err != nil {
			return nil, err
		}

		results = append(results, elem)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	// Close the cursor once finished
	if err := cur.Close(context.TODO()); err != nil {
		return nil, err
	}

	return results, nil
}
