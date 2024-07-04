package security

import (
	mapset "github.com/deckarep/golang-set/v2"
	"github.com/gin-gonic/gin"
)

// Security policies.
const (
	// PolicyNoVerification does not require any verification. Default for student.
	PolicyNoVerification = "no"

	// PolicyCredentials requires unauthorized user to enter their credentials. User data is stored locally.
	PolicyCredentials = "credentials"

	// PolicyHostOnly blocks any request from machines other than host.
	PolicyHostOnly = "hostonly"
)

// Middleware returns a gin middleware to apply policy for a route.
func Middleware(role, policy string) gin.HandlerFunc {
	switch policy {
	case PolicyHostOnly:
		return HostOnlyMiddleware
	case PolicyCredentials:
		return CredentialsMiddleware(role)
	default:
		return NoVerificationMiddleware
	}
}

// Register registers policy-specific routes.
func Register(e *gin.Engine, policies ...string) {
	policySet := mapset.NewSet[string]()

	for _, policy := range policies {
		policySet.Add(policy)
	}

	if policySet.Contains(PolicyCredentials) {
		CredentialsRegister(e)
	}
}
