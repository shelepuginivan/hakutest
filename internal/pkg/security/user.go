package security

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"

	"github.com/shelepuginivan/hakutest/internal/pkg/fsutil"
	"github.com/shelepuginivan/hakutest/internal/pkg/paths"
)

// User represents user data stored locally.
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

// UserFile returns full path to the file with user data.
func UserFile(username string) string {
	return filepath.Join(paths.Users, strings.ToLower(username))
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

	data, err := os.ReadFile(UserFile(username))
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
	data, err := json.Marshal(u)
	if err != nil {
		return err
	}

	return fsutil.WriteAll(UserFile(u.Username), data)
}

// DeleteUser deletes locally stored user.
func DeleteUser(username string) error {
	return os.RemoveAll(UserFile(username))
}
