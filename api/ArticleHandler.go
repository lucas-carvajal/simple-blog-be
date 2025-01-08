package api

import (
	"net/http"
	"simple-blog-be/repository"

	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
	ArticlesRepository *repository.ArticlesRepository
}

func (h *ArticleHandler) GetArticleByID(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Article ID is required"})
		return
	}

	article, err := h.ArticlesRepository.GetArticleByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve article"})
		return
	}

	// Convert entity to DTO before returning
	articleDto := FromEntity(article)
	c.JSON(http.StatusOK, articleDto)
}

func (h *ArticleHandler) GetArticleComments(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "comments coming soon",
	})
}

func (h *ArticleHandler) AddComment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "comments coming soon",
	})
}
