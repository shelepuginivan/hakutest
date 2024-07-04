package security

import (
	"crypto/rand"
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

// JWT claims constants.
const (
	claimRole = "role"
	claimUser = "user"
)

// JWT key constants.
var (
	jwtKey    = make([]byte, 32)
	jwtKeyErr error
)

func init() {
	// Generate random JWT key.
	_, jwtKeyErr = rand.Read(jwtKey)
}

// GenerateJWT generates JSON Web Token (JWT) for the provided credentials.
// It uses secret key generated once at runtime.
func GenerateJWT(credentials *Credentials) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims[claimRole] = credentials.Role
	claims[claimUser] = credentials.Username

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", fmt.Errorf("failed to generate JWT token: %v", err)
	}

	return tokenString, nil
}

// ParseJWT returns credentials associated with provided JSON Web Token (JWT).
// It uses secret key generated once at runtime.
func ParseJWT(tokenString string) (*Credentials, error) {
	token, err := jwt.Parse(tokenString, func(_ *jwt.Token) (interface{}, error) {
		return jwtKey, jwtKeyErr
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("failed to parse JWT token: %v", err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("invalid JWT token")
	}

	return &Credentials{
		Role:     claims[claimRole].(string),
		Username: claims[claimUser].(string),
	}, nil
}
