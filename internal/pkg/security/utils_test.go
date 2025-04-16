package security_test

import (
	"testing"

	"github.com/shelepuginivan/hakutest/internal/pkg/security"
	"github.com/stretchr/testify/assert"
)

func TestHasPermissions(t *testing.T) {
	tests := []struct {
		c             *security.Credentials
		requiredRoles []string
		want          bool
	}{
		{
			c: &security.Credentials{
				Username: "John Doe",
				Roles:    []string{"teacher", "admin", "user"},
			},
			requiredRoles: []string{"user", "admin"},
			want:          true,
		},
		{
			c: &security.Credentials{
				Username: "Jane Smith",
				Roles:    []string{"user"},
			},
			requiredRoles: []string{"admin"},
			want:          false,
		},
		{
			c: &security.Credentials{
				Username: "Alice Johnson",
				Roles:    []string{"editor", "user"},
			},
			requiredRoles: []string{"editor"},
			want:          true,
		},
		{
			c: &security.Credentials{
				Username: "Bob Brown",
				Roles:    []string{},
			},
			requiredRoles: []string{"user"},
			want:          false,
		},
		{
			c: &security.Credentials{
				Username: "Charlie Black",
				Roles:    []string{"admin", "user"},
			},
			requiredRoles: []string{"admin", "editor"},
			want:          false,
		},
		{
			c: &security.Credentials{
				Username: "Diana White",
				Roles:    []string{"user", "guest"},
			},
			requiredRoles: []string{"guest"},
			want:          true,
		},
		{
			c: &security.Credentials{
				Username: "Eve Green",
				Roles:    []string{"user"},
			},
			requiredRoles: []string{"user", "admin"},
			want:          false,
		},
		{
			c: &security.Credentials{
				Username: "Frank Blue",
				Roles:    []string{"admin", "superuser"},
			},
			requiredRoles: []string{"superuser"},
			want:          true,
		},
		{
			c: &security.Credentials{
				Username: "Grace Yellow",
				Roles:    []string{"user"},
			},
			requiredRoles: []string{"user", "admin", "editor"},
			want:          false,
		},
		{
			c: &security.Credentials{
				Username: "Hank Red",
				Roles:    []string{"admin", "user", "editor"},
			},
			requiredRoles: []string{"admin", "user", "editor"},
			want:          true,
		},
		{
			c: &security.Credentials{
				Username: "Ivy Purple",
				Roles:    []string{"user"},
			},
			requiredRoles: []string{},
			want:          true,
		},
	}

	for _, tt := range tests {
		got := security.HasPermissions(tt.c, tt.requiredRoles)
		assert.Equal(t, tt.want, got)
	}
}
