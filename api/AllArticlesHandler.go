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
	c.JSON(http.StatusOK, gin.H{
		"message": "Getting all articles",
	})
}

func (h *AllArticlesHandler) SearchArticles(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Searching articles",
	})
}
