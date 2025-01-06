package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ArticleHandler struct {
}

func (h *ArticleHandler) GetArticleByID(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Getting article by ID",
	})
}

func (h *ArticleHandler) GetArticleComments(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Getting article comments",
	})
}

func (h *ArticleHandler) AddComment(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Adding comment to article",
	})
}
