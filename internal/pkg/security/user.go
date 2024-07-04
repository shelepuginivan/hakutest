package security

import (
	"encoding/json"
	"os"
	"path/filepath"

	"github.com/shelepuginivan/hakutest/internal/pkg/fsutil"
	"github.com/shelepuginivan/hakutest/internal/pkg/paths"
)

// User represents user data stored locally.
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// NewUser returns a new instance of User.
// Provided password is hashed using HMAC-SHA256 algorithm.
func NewUser(username, password string, role string) *User {
	return &User{
		Username: username,
		Password: HashPassword(password),
		Role:     role,
	}
}

// GetUser returns a locally stored user, if it exists.
func GetUser(username string) (*User, error) {
	var u User

	userFile := filepath.Join(paths.Users, u.Username)
	data, err := os.ReadFile(userFile)
	if err != nil {
		return nil, err
	}

	if err = json.Unmarshal(data, &u); err != nil {
		return nil, err
	}

	return &u, nil
}

// SaveUser saves user locally.
func SaveUser(u *User) error {
	userFile := filepath.Join(paths.Users, u.Username)
	data, err := json.Marshal(u)
	if err != nil {
		return err
	}

	return fsutil.WriteAll(userFile, data)
}

// DeleteUser deletes locally stored user.
func DeleteUser(username string) error {
	return os.RemoveAll(filepath.Join(paths.Users, username))
}
