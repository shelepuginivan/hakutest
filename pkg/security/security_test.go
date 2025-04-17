package security_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	isecurity "github.com/shelepuginivan/hakutest/internal/pkg/security"
	"github.com/shelepuginivan/hakutest/pkg/security"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
)

func TestMain(m *testing.M) {
	isecurity.InitDBWithDialector(sqlite.Open(":memory:"))
	gin.SetMode(gin.TestMode)
	m.Run()
}

func TestAbortAllMiddleware(t *testing.T) {
	t.Run("should redirect to / if request path is not /", func(t *testing.T) {
		router := gin.Default()
		router.GET("/path", security.AbortAllMiddleware)

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/path", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusSeeOther, w.Code)

		loc, err := w.Result().Location()
		assert.NoError(t, err)
		assert.Equal(t, "/", loc.Path)
	})

	t.Run("should handle requests to / to avoid infinite redirects", func(t *testing.T) {
		router := gin.Default()
		router.GET("/path", security.AbortAllMiddleware)
		router.Use(func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"some": "value",
			})
		})

		w := httptest.NewRecorder()
		req, _ := http.NewRequest("GET", "/", nil)
		router.ServeHTTP(w, req)

		assert.Equal(t, http.StatusOK, w.Code)
		assert.JSONEq(t, `{"some":"value"}`, w.Body.String())
	})
}

func TestCredentialsMiddleware(t *testing.T) {
	isecurity.CreateUser("John Doe", "123456", []string{security.RoleStudent})
	isecurity.CreateUser("Lorem Ipsum", "letmein", []string{security.RoleTeacher})

	router := gin.Default()
	router.GET("/protected", security.CredentialsMiddleware([]string{security.RoleTeacher}), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"some_secure_data": 42,
		})
	})

	// No JWT cookie is set.
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/protected", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusSeeOther, w.Code)
	loc, err := w.Result().Location()
	assert.NoError(t, err)
	assert.Equal(t, "/auth", loc.Path)

	// JWT cookie is set but user does not have enough permissions.
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/protected", nil)

	jwtCookie, _ := isecurity.GenerateJWT(&isecurity.Credentials{
		Username: "John Doe",
		Roles:    []string{security.RoleStudent},
	})

	req.AddCookie(&http.Cookie{
		Name:  "jwt",
		Value: jwtCookie,
	})

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusSeeOther, w.Code)
	loc, err = w.Result().Location()
	assert.NoError(t, err)
	assert.Equal(t, "/auth", loc.Path)

	// JWT is valid and has enough permissions to access the page.
	w = httptest.NewRecorder()
	req, _ = http.NewRequest("GET", "/protected", nil)

	jwtCookie, _ = isecurity.GenerateJWT(&isecurity.Credentials{
		Username: "Lorem Ipsum",
		Roles:    []string{security.RoleTeacher},
	})

	req.AddCookie(&http.Cookie{
		Name:  "jwt",
		Value: jwtCookie,
	})

	router.ServeHTTP(w, req)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"some_secure_data":42}`, w.Body.String())
}

func TestNoVerificationMiddleware(t *testing.T) {
	router := gin.Default()
	router.GET("/well", security.NoVerificationMiddleware, func(c *gin.Context) {
		c.Data(http.StatusNoContent, gin.MIMEHTML, nil)
	})

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/well", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusNoContent, w.Code)
}
