package security

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"

	mapset "github.com/deckarep/golang-set/v2"
	"github.com/denisbrodbeck/machineid"
)

// Unique identifier for this machine.
var machineID string

func init() {
	machineID, _ = machineid.ProtectedID("hakutest")
}

// HashPassword hashes password using HMAC-SHA256 algorithm.
func HashPassword(password string) string {
	h := hmac.New(sha256.New, []byte(machineID))
	h.Write([]byte(password))
	return hex.EncodeToString(h.Sum(nil))
}

// ComparePasswords reports whether password matches with expected hash.
func ComparePasswords(password, expected string) bool {
	return HashPassword(password) == expected
}

// HasPermissions reports whether user has permission for the action based on
// roles.
func HasPermissions(c *Credentials, requiredRoles []string) bool {
	userRoles := mapset.NewSet(c.Roles...)
	wantRoles := mapset.NewSet(requiredRoles...)

	return userRoles.IsSuperset(wantRoles)
}
