package repository

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ArticleEntity struct {
	ID        primitive.ObjectID `bson:"_id"`
	Title     string             `bson:"title"`
	Subheader string             `bson:"subheader"`
	Content   []ContentEntity    `bson:"content"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

type BaseContentEntity struct {
	Order int `bson:"order"`
}

type ContentEntity interface{}

type ParagraphEntity struct {
	Metadata BaseContentEntity `bson:"metadata"`
	Text     string            `bson:"text"`
}

// Custom unmarshal method for ArticleEntity
func (a *ArticleEntity) UnmarshalBSON(data []byte) error {
	// Create an auxiliary struct with same fields but Content as []map[string]interface{}
	type Aux struct {
		ID        primitive.ObjectID       `bson:"_id"`
		Title     string                   `bson:"title"`
		Subheader string                   `bson:"subheader"`
		Content   []map[string]interface{} `bson:"content"`
		CreatedAt time.Time                `bson:"createdAt"`
		UpdatedAt time.Time                `bson:"updatedAt"`
	}

	// First unmarshal into auxiliary struct
	var aux Aux
	if err := bson.Unmarshal(data, &aux); err != nil {
		return fmt.Errorf("failed to unmarshal article: %v", err)
	}

	// Copy the simple fields
	a.ID = aux.ID
	a.Title = aux.Title
	a.Subheader = aux.Subheader
	a.CreatedAt = aux.CreatedAt
	a.UpdatedAt = aux.UpdatedAt

	// Convert the content items
	a.Content = make([]ContentEntity, len(aux.Content))
	for i, c := range aux.Content {
		// Extract metadata
		metadataMap, ok := c["metadata"].(map[string]interface{})
		if !ok {
			return fmt.Errorf("invalid metadata format at content index %d", i)
		}

		// Convert order to int - in BSON it should come as int32 or int64
		var order int
		switch v := metadataMap["order"].(type) {
		case int32:
			order = int(v)
		case int64:
			order = int(v)
		default:
			return fmt.Errorf("invalid order format at content index %d", i)
		}

		// Extract text
		text, ok := c["text"].(string)
		if !ok {
			return fmt.Errorf("invalid text format at content index %d", i)
		}

		// Create ParagraphEntity
		a.Content[i] = ParagraphEntity{
			Metadata: BaseContentEntity{
				Order: order,
			},
			Text: text,
		}
	}

	return nil
}
