package security

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shelepuginivan/hakutest/internal/pkg/i18n"
	"github.com/shelepuginivan/hakutest/internal/pkg/security"
)

// LoginForm represents data sent in authorization form.
type LoginForm struct {
	Username string `form:"username"`
	Password string `form:"password"`
	To       string `form:"to"`
}

// CredentialsMiddleware is a gin middleware that sets route security policy.
// Protected route can only be accessed by an authorized user.
// If user is unauthorized, they are redirected to the authorization page.
// It must be used with the `CredentialsRegister` method.
func CredentialsMiddleware(roles []string) gin.HandlerFunc {
	return func(c *gin.Context) {
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

		if !security.HasPermissions(credentials, roles) {
			c.Redirect(http.StatusSeeOther, location)
			c.Abort()
			return
		}

		c.Next()
	}
}

// CredentialsRegister registers additional routes for the `CredentialsMiddleware`.
func CredentialsRegister(router gin.IRouter) {
	router.GET("/auth", func(c *gin.Context) {
		c.HTML(http.StatusOK, "auth.gohtml", gin.H{
			"Invalid": false,
			"To":      c.Query("to"),
		})
	})

	router.POST("/auth", func(c *gin.Context) {
		var form LoginForm

		if err := c.Bind(&form); err != nil {
			c.HTML(http.StatusUnprocessableEntity, "error.gohtml", gin.H{
				"Title":   i18n.Get("err.unprocessable.title"),
				"Text":    i18n.Get("err.unprocessable.text"),
				"Code":    http.StatusUnprocessableEntity,
				"Message": "failed to parse authorization form",
				"Error":   err.Error(),
			})
			return
		}

		credentials, err := security.Login(form.Username, form.Password)
		if err != nil {
			c.HTML(http.StatusUnauthorized, "auth.gohtml", gin.H{
				"Invalid": true,
				"To":      c.Query("to"),
			})
			return
		}

		jwt, err := security.GenerateJWT(credentials)
		if err != nil {
			c.HTML(http.StatusUnprocessableEntity, "error.gohtml", gin.H{
				"Title":   i18n.Get("err.jwt.title"),
				"Text":    i18n.Get("err.jwt.text"),
				"Code":    http.StatusUnprocessableEntity,
				"Message": "failed to generate jwt",
				"Error":   err.Error(),
			})
			return
		}

		c.SetCookie("jwt", jwt, 60*60*24, "/", "", true, false)
		c.Redirect(http.StatusSeeOther, form.To)
	})
}
