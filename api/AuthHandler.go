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

	//session.Options.Secure = true                      // Only send the cookie over HTTPS
	session.Options.HttpOnly = true                    // Prevent client-side JavaScript from reading the cookie
	session.Options.SameSite = http.SameSiteStrictMode // Helps prevent CSRF attacks

	session.Save(c.Request, c.Writer)

	c.Status(http.StatusOK)
}

func (h *AuthHandler) Logout(c *gin.Context) {
	session, _ := h.CookieStore.Get(c.Request, "session-name")
	// Delete the session or mark it as invalid
	session.Options.MaxAge = -1
	session.Save(c.Request, c.Writer)
	c.Redirect(http.StatusSeeOther, "/")
}

func (h *AuthHandler) IsAuthenticated(c *gin.Context) {
	session, _ := h.CookieStore.Get(c.Request, "session-name")
	if auth, ok := session.Values["authenticated"].(bool); !ok || !auth {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
		return
	}
	c.Status(http.StatusOK)
}
