package security

import (
	"sync"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/gin-gonic/gin"
	"github.com/shelepuginivan/hakutest/internal/pkg/security"
)

// SecurityFields represents configuration for security policies.
type Config struct {
	// DSN of the database used to store user data.
	//
	// For SQLite dialect you can use `file::memory:?cache=shared`, in this
	// case user data will be stored in memory.
	DSN string `json:"dsn,omitempty" yaml:"dsn,omitempty"`

	// Dialect of the database used to store user data.
	//
	// List of supported dialects:
	//   - `postgres` for PostgreSQL,
	//   - `mysql` for MySQL,
	//   - `sqlite` for SQLite (default).
	Dialect string `json:"dialect,omitempty" yaml:"dialect,omitempty"`

	// Teacher security policy.
	Teacher string `json:"teacher,omitempty" yaml:"teacher,omitempty"`

	// Student security policy.
	Student string `json:"student,omitempty" yaml:"student,omitempty"`
}

var (
	mu sync.Mutex
)

func Init(cfg Config) {
	mu.Lock()
	defer mu.Unlock()

	security.InitDB(cfg.DSN, cfg.Dialect)
}

// Security policies.
const (
	// PolicyAbortAll aborts all requests.
	PolicyAbortAll = "abort_all"

	// PolicyNoVerification does not require any verification. Default for
	// student.
	PolicyNoVerification = "no_verification"

	// PolicyCredentials requires unauthorized user to enter their credentials.
	// User data is stored locally.
	PolicyCredentials = "credentials"

	// PolicyHostOnly blocks any request from machines other than host.
	PolicyHostOnly = "hostonly"
)

// Middleware returns a gin middleware to apply policy for a route.
// For `credentials` policy required roles can be specified.
func Middleware(policy string, roles ...string) gin.HandlerFunc {
	switch policy {
	case PolicyAbortAll:
		return AbortAllMiddleware
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
