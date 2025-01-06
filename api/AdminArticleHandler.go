package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type AdminArticleHandler struct {
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
