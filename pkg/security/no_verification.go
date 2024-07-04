package security

import "github.com/gin-gonic/gin"

// NoVerificationMiddleware is a gin middleware that sets route security policy.
// This middleware does not perform any checks and calls the next handler in
// the chain.
func NoVerificationMiddleware(c *gin.Context) {
	c.Next()
}
