package main

import (
	"context"
	"fmt"
	"simple-blog-be/api"
	"simple-blog-be/repository"
	"simple-blog-be/utils"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	r := gin.Default()

	// Initialize repository
	client := setUpMongoDbConnection()
	articlesRepository := repository.NewArticlesRepository(client)

	// Initialize all handlers
	allArticlesHandler := api.AllArticlesHandler{ArticlesRepository: articlesRepository}
	articleHandler := api.ArticleHandler{ArticlesRepository: articlesRepository}
	adminArticleHandler := api.AdminArticleHandler{ArticlesRepository: articlesRepository}

	// Public routes for articles
	r.GET("/articles", allArticlesHandler.GetAllArticles)
	r.GET("/articles/search", allArticlesHandler.SearchArticles)

	// Single article routes
	r.GET("/article/:id", articleHandler.GetArticleByID)
	r.GET("/article/:id/comments", articleHandler.GetArticleComments)
	r.POST("/article/:id/comments", articleHandler.AddComment)

	// Admin routes
	admin := r.Group("/admin")
	{
		admin.POST("/article", adminArticleHandler.CreateArticle)
		admin.PUT("/article/:id", adminArticleHandler.UpdateArticle)
		admin.DELETE("/article/:id", adminArticleHandler.DeleteArticle)
	}

	r.Run(":8080")
}

func setUpMongoDbConnection() *mongo.Client {
	mongoPassword := utils.MONGO_PASSWORD
	uri := fmt.Sprintf("mongodb+srv://lcscarvajal:%s@cluster0.d1pwr.mongodb.net/?retryWrites=true&w=majority", mongoPassword)

	// Set client options
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	clientOptions := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)

	// Connect to MongoDB
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic("Could not connect to MongoDB: " + err.Error())
	}

	// Ping the database to verify connection
	err = client.Ping(ctx, nil)
	if err != nil {
		panic("Could not ping MongoDB: " + err.Error())
	}

	fmt.Println("Connected to MongoDB!")
	return client
}
