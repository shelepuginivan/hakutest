package security

import (
	"sync"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/gin-gonic/gin"
)

// SecurityFields represents configuration for security policies.
type Config struct {
	// Teacher security policy.
	Teacher string `yaml:"teacher"`

	// Student security policy.
	Student string `yaml:"student"`
}

var (
	mu sync.Mutex
)

func Init(cfg Config) {
	mu.Lock()
	defer mu.Unlock()
}

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
// For `credentials` policy required roles can be specified.
func Middleware(policy string, roles ...string) gin.HandlerFunc {
	switch policy {
	case PolicyHostOnly:
		return HostOnlyMiddleware
	case PolicyCredentials:
		return CredentialsMiddleware(roles)
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
