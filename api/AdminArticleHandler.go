package api

import (
	"net/http"
	"simple-blog-be/repository"

	"github.com/gin-gonic/gin"
)

type AdminArticleHandler struct {
	ArticlesRepository *repository.ArticlesRepository
}

// CreateArticle handles the creation of a new article
func (h *AdminArticleHandler) CreateArticle(c *gin.Context) {
	var articleDto ArticleDto
	if err := c.ShouldBindJSON(&articleDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Convert DTO to entity
	articleEntity, err := articleDto.ToEntity()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid article data"})
		return
	}

	// Create article using repository
	createdArticle, err := h.ArticlesRepository.CreateArticle(c.Request.Context(), *articleEntity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create article"})
		return
	}

	// Convert created entity back to DTO
	c.JSON(http.StatusCreated, FromEntity(createdArticle))
}

// UpdateArticle handles the update of an existing article
func (h *AdminArticleHandler) UpdateArticle(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Article ID is required"})
		return
	}

	var articleDto ArticleDto
	if err := c.ShouldBindJSON(&articleDto); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	// Convert DTO to entity
	articleEntity, err := articleDto.ToEntity()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid article data"})
		return
	}

	// Update article using repository
	err = h.ArticlesRepository.UpdateArticle(c.Request.Context(), id, *articleEntity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update article"})
		return
	}

	// Get updated article
	updatedArticle, err := h.ArticlesRepository.GetArticleByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve updated article"})
		return
	}

	c.JSON(http.StatusOK, FromEntity(updatedArticle))
}

// DeleteArticle handles the deletion of an article
func (h *AdminArticleHandler) DeleteArticle(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Article ID is required"})
		return
	}

	// Delete article using repository
	err := h.ArticlesRepository.DeleteArticle(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete article"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Article deleted successfully"})
}
