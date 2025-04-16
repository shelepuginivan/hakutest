package security_test

import (
	"testing"

	"github.com/shelepuginivan/hakutest/internal/pkg/security"
	"github.com/stretchr/testify/assert"
)

func TestGenerateJWT(t *testing.T) {
	tests := []*security.Credentials{
		{
			Roles:    []string{"user"},
			Username: "testuser",
		},
		{
			Roles:    []string{"teacher", "student"},
			Username: "John Doe",
		},
	}

	for _, tt := range tests {
		t.Run(tt.Username, func(t *testing.T) {
			credentials := tt

			tokenString, err := security.GenerateJWT(credentials)
			assert.NoError(t, err)
			assert.NotEmpty(t, tokenString)

			parsedCredentials, err := security.ParseJWT(tokenString)
			assert.NoError(t, err)
			assert.Equal(t, credentials.Username, parsedCredentials.Username)
			assert.ElementsMatch(t, credentials.Roles, parsedCredentials.Roles)
		})
	}

}

func TestParseJWT_InvalidToken(t *testing.T) {
	invalidToken := "invalid.token.string"
	_, err := security.ParseJWT(invalidToken)
	assert.Error(t, err, "Expected error while parsing invalid JWT")
}
