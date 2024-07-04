// Package security provides private security methods.
package security

import "fmt"

// Credentials represents authorization data.
// It contains Role and Username.
type Credentials struct {
	Role     string
	Username string
}

// Login performs a sign in operation for provided username and password.
func Login(username, password string) (*Credentials, error) {
	user, err := GetUser(username)
	if err != nil {
		return nil, err
	}

	if username != user.Username || !ComparePasswords(password, user.Password) {
		return nil, fmt.Errorf("invalid username or password")
	}

	return &Credentials{
		Username: user.Username,
		Role:     user.Role,
	}, nil
}
