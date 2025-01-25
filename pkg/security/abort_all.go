package security

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// AbortAllMiddleware redirects all requests to the index route and aborts
// them. Protected route cannot be accessed.
func AbortAllMiddleware(c *gin.Context) {
	// If request is directly to /, it is allowed. This is required to avoid
	// infinite redirects.
	if c.Request.URL.Path == "/" {
		c.Next()
		return
	}

	c.Redirect(http.StatusSeeOther, "/")
	c.Abort()
}
