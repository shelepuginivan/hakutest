// Package security provides private security methods.
package security

// Credentials represents authorization data.
// It contains Username and Password.
type Credentials struct {
	Username string
	Password string
}

func IsValidCredentials(c *Credentials) bool {
	// TODO: Add credentials validation.
	return false
}
