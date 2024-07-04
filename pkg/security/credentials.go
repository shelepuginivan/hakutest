package security

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shelepuginivan/hakutest/internal/pkg/i18n"
	"github.com/shelepuginivan/hakutest/internal/pkg/security"
)

// CredentialsMiddleware is a gin middleware that sets route security policy.
// Protected route can only be accessed by an authorized user.
// If user is unauthorized, they are redirected to the authorization page.
// It must be used with the `CredentialsRegister` method.
func CredentialsMiddleware(c *gin.Context) {
	location := fmt.Sprintf("/auth?to=%s", c.Request.RequestURI)

	jwt, err := c.Cookie("jwt")
	if err != nil {
		c.Redirect(http.StatusSeeOther, location)
		c.Abort()
		return
	}

	credentials, err := security.ParseJWT(jwt)
	if err != nil {
		c.Redirect(http.StatusSeeOther, location)
		c.Abort()
		return
	}

	if security.IsValidCredentials(credentials) {
		c.Next()
		return
	}

	c.Redirect(http.StatusSeeOther, "/auth")
	c.Abort()
}

// CredentialsRegister registers additional routes for the `CredentialsMiddleware`.
func CredentialsRegister(e *gin.Engine) {
	e.GET("/auth", func(c *gin.Context) {
		c.HTML(http.StatusOK, "auth.gohtml", gin.H{
			"Invalid": false,
			"To":      c.Query("to"),
		})
	})

	e.POST("/auth", func(c *gin.Context) {
		err := c.Request.ParseForm()
		if err != nil {
			c.HTML(http.StatusUnprocessableEntity, "error.gohtml", gin.H{
				"Title":   i18n.Get("auth.unprocessable.title"),
				"Text":    i18n.Get("auth.unprocessable.text"),
				"Code":    http.StatusUnprocessableEntity,
				"Message": "failed to parse form",
				"Error":   err.Error(),
			})
			return
		}

		credentials := &security.Credentials{
			Username: c.PostForm("username"),
			Password: c.PostForm("password"),
		}

		to := c.PostForm("to")

		if security.IsValidCredentials(credentials) {
			c.HTML(http.StatusUnauthorized, "auth.gohtml", gin.H{
				"Invalid": true,
				"To":      to,
			})
			return
		}

		jwt, err := security.GenerateJWT(credentials)
		if err != nil {
			c.HTML(http.StatusUnprocessableEntity, "error.gohtml", gin.H{
				"Title":   i18n.Get("auth.jwt_generation_err.title"),
				"Text":    i18n.Get("auth.jwt_generation_err.text"),
				"Code":    http.StatusUnprocessableEntity,
				"Message": "failed to parse form",
				"Error":   err.Error(),
			})
			return
		}

		c.SetCookie("jwt", jwt, 60*60*24, "/", "", true, false)
		c.Redirect(http.StatusSeeOther, to)
	})
}
