package repository

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Article struct {
	ID        primitive.ObjectID `bson:"_id"`
	Title     string             `bson:"title"`
	Subheader string             `bson:"subheader"`
	Content   []Content          `bson:"content"`
	CreatedAt time.Time          `bson:"createdAt"`
	UpdatedAt time.Time          `bson:"updatedAt"`
}

type BaseContent struct {
	Order int `bson:"order"`
}

type Content interface{}

type Paragraph struct {
	Metadata BaseContent `bson:"metadata"`
	Text     string      `bson:"text"`
}
