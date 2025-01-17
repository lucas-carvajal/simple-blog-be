package api

import (
	"net/http"
	"simple-blog-be/utils"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
)

type AuthHandler struct {
	CookieStore *sessions.CookieStore
}

func (h *AuthHandler) Login(c *gin.Context) {
	password := c.PostForm("password")

	if password != utils.USER_PASSWORD {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	// Create a session
	session, err := h.CookieStore.Get(c.Request, "session-name")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create session"})
		return
	}
	session.Values["authenticated"] = true
	session.Save(c.Request, c.Writer)

	c.Redirect(http.StatusSeeOther, "/admin")
}

func (h *AuthHandler) Logout(c *gin.Context) {
	session, _ := h.CookieStore.Get(c.Request, "session-name")
	// Delete the session or mark it as invalid
	session.Options.MaxAge = -1
	session.Save(c.Request, c.Writer)
	c.Redirect(http.StatusSeeOther, "/")
}

func (h *AuthHandler) IsAuthenticated(c *gin.Context) {

}
