package security

import "github.com/gin-gonic/gin"

const (
	PolicyNoVerification = "no"
	PolicyCredentials    = "credentials"
	PolicyHostOnly       = "hostonly"
)

// Middleware returns a gin middleware to apply policy for a route.
func Middleware(policy string) gin.HandlerFunc {
	switch policy {
	case PolicyHostOnly:
		return HostOnlyMiddleware
	case PolicyCredentials:
		return CredentialsMiddleware
	default:
		return NoVerificationMiddleware
	}
}

// Register registers policy-specific routes.
func Register(e *gin.Engine, policy string) {
	if policy == PolicyCredentials {
		CredentialsRegister(e)
	}
}
