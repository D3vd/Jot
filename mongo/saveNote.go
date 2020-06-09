package mongo

import (
	"context"
	"github.com/R3l3ntl3ss/Jot/models"
	"time"
)

func (m Mongo) SaveNote(content string) error {

	note := models.Notes{
		Content: content,
		Created: time.Now(),
	}

	//	Insert Note into the DB
	insertResult, err := m.DB.Collection("notes").InsertOne(context.Background(), note)

	// Check if the note was successfully inserted into the DB
	if err != nil || insertResult.InsertedID == nil {
		return err
	}

	return nil
}
