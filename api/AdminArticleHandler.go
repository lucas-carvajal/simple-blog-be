package api

import (
	"net/http"
	"simple-blog-be/repository"

	"github.com/gin-gonic/gin"
)

type AdminArticleHandler struct {
	ArticlesRepository *repository.ArticlesRepository
}

func (h *AdminArticleHandler) CreateArticle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Creating new article",
	})
}

func (h *AdminArticleHandler) UpdateArticle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Updating article",
	})
}

func (h *AdminArticleHandler) DeleteArticle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Deleting article",
	})
}
