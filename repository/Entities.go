package repository

import (
	"time"

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
