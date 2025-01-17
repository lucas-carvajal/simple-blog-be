package main

import (
	"context"
	"fmt"
	"net/http"
	"simple-blog-be/api"
	"simple-blog-be/repository"
	"simple-blog-be/utils"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var store = sessions.NewCookieStore([]byte(utils.SESSION_ENCRYPTION_KEY))

func main() {
	r := gin.Default()

	// Initialize repository
	client := setUpMongoDbConnection()
	articlesRepository := repository.NewArticlesRepository(client)

	// Initialize all handlers
	authHandler := api.AuthHandler{}
	allArticlesHandler := api.AllArticlesHandler{ArticlesRepository: articlesRepository}
	articleHandler := api.ArticleHandler{ArticlesRepository: articlesRepository}
	adminArticleHandler := api.AdminArticleHandler{ArticlesRepository: articlesRepository}

	// Set up CORS
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	})

	// Auth routes
	auth := r.Group("/auth")
	{
		auth.POST("/login", authHandler.Login)
		auth.POST("/logout", authHandler.Logout)
		auth.GET("/is-authenticated", authHandler.IsAuthenticated)
	}

	// Public routes for articles
	r.GET("/articles", allArticlesHandler.GetAllArticles)
	r.GET("/articles/search", allArticlesHandler.SearchArticles)

	// Single article routes
	r.GET("/article/:id", articleHandler.GetArticleByID)
	r.GET("/article/:id/comments", articleHandler.GetArticleComments)
	r.POST("/article/:id/comments", articleHandler.AddComment)

	// Admin routes
	admin := r.Group("/admin")
	admin.Use(authMiddleware())
	{
		admin.POST("/article", adminArticleHandler.CreateArticle)
		admin.PUT("/article/:id", adminArticleHandler.UpdateArticle)
		admin.DELETE("/article/:id", adminArticleHandler.DeleteArticle)
	}

	r.Run(":8080")
}

func authMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		session, _ := store.Get(c.Request, "session-name")
		if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			c.Abort()
			return
		}
		c.Next()
	}
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
