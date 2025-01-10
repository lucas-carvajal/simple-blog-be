package api

import (
	"net/http"
	"simple-blog-be/repository"

	"github.com/gin-gonic/gin"
)

type AllArticlesHandler struct {
	ArticlesRepository *repository.ArticlesRepository
}

func (h *AllArticlesHandler) GetAllArticles(c *gin.Context) {
	articles, err := h.ArticlesRepository.GetAllArticles(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve articles"})
		return
	}

	c.JSON(http.StatusOK, FromEntitiesWithFirstContent(articles))
}

func (h *AllArticlesHandler) SearchArticles(c *gin.Context) {
	query := c.Query("q")
	if query == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Search query is required"})
		return
	}

	articles, err := h.ArticlesRepository.SearchArticles(c.Request.Context(), query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to search articles"})
		return
	}

	c.JSON(http.StatusOK, FromEntitiesWithFirstContent(articles))
}
