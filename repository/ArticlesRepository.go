package repository

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ArticlesRepository struct {
	articles *mongo.Collection
}

func NewArticlesRepository(client *mongo.Client) *ArticlesRepository {
	return &ArticlesRepository{
		articles: client.Database("simple-blog").Collection("articles"),
	}
}

// GetAllArticles retrieves all articles ordered by updatedAt desc
func (r *ArticlesRepository) GetAllArticles(ctx context.Context) ([]ArticleEntity, error) {
	opts := options.Find().SetSort(bson.D{{Key: "updatedAt", Value: -1}})
	cursor, err := r.articles.Find(ctx, bson.M{}, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var articles []ArticleEntity
	if err = cursor.All(ctx, &articles); err != nil {
		return nil, err
	}
	return articles, nil
}

// SearchArticles searches articles by title or subheader ordered by updatedAt desc
func (r *ArticlesRepository) SearchArticles(ctx context.Context, query string) ([]ArticleEntity, error) {
	filter := bson.M{
		"$or": []bson.M{
			{"title": bson.M{"$regex": query, "$options": "i"}},
			{"subheader": bson.M{"$regex": query, "$options": "i"}},
		},
	}

	opts := options.Find().SetSort(bson.D{{Key: "updatedAt", Value: -1}})
	cursor, err := r.articles.Find(ctx, filter, opts)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	var articles []ArticleEntity
	if err = cursor.All(ctx, &articles); err != nil {
		return nil, err
	}
	return articles, nil
}

// GetArticleByID retrieves a single article by ID
func (r *ArticlesRepository) GetArticleByID(ctx context.Context, id string) (*ArticleEntity, error) {
	fmt.Println(id)
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	fmt.Println(objectID)

	var article ArticleEntity
	err = r.articles.FindOne(ctx, bson.M{"_id": objectID}).Decode(&article)
	fmt.Println(err)
	if err != nil {
		return nil, err
	}
	return &article, nil
}

// CreateArticle creates a new article
func (r *ArticlesRepository) CreateArticle(ctx context.Context, article ArticleEntity) (*ArticleEntity, error) {
	article.ID = primitive.NewObjectID()
	article.CreatedAt = time.Now()
	article.UpdatedAt = time.Now()

	_, err := r.articles.InsertOne(ctx, article)
	if err != nil {
		return nil, err
	}

	return &article, nil
}

// UpdateArticle updates an existing article
func (r *ArticlesRepository) UpdateArticle(ctx context.Context, id string, article ArticleEntity) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	article.UpdatedAt = time.Now()
	update := bson.M{
		"$set": bson.M{
			"title":     article.Title,
			"subheader": article.Subheader,
			"content":   article.Content,
			"updatedAt": time.Now(),
		},
	}

	_, err = r.articles.UpdateOne(ctx, bson.M{"_id": objectID}, update)
	return err
}

// DeleteArticle deletes an article
func (r *ArticlesRepository) DeleteArticle(ctx context.Context, id string) error {
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	_, err = r.articles.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}
