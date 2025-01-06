package repository

import (
	"go.mongodb.org/mongo-driver/mongo"
)

type ArticlesRepository struct {
	articles *mongo.Collection
}

func NewArticlesRepository(client *mongo.Client) *ArticlesRepository {
	return &ArticlesRepository{
		articles: client.Database("simple-blog").Collection("articles"),
	}
}
