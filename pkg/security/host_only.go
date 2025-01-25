package security

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/shelepuginivan/hakutest/internal/pkg/network"
)

// HostOnlyMiddleware is a gin middleware that sets route security policy.
// Protected route can only be accessed from the host machine.
// Requests from IPs other than the host IP are redirected to the index route.
func HostOnlyMiddleware(c *gin.Context) {
	// If request is directly to /, it is allowed. This is required to avoid
	// infinite redirects.
	if network.IsLocalIP(c.ClientIP()) || c.Request.URL.Path == "/" {
		c.Next()
		return
	}

	c.Redirect(http.StatusSeeOther, "/")
	c.Abort()
}
