package main

import (
	"simple-blog-be/api"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Initialize all handlers
	allArticlesHandler := api.AllArticlesHandler{}
	articleHandler := api.ArticleHandler{}
	adminArticleHandler := api.AdminArticleHandler{}

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
