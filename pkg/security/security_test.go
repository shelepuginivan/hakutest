package security_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/shelepuginivan/hakutest/pkg/security"
	"github.com/stretchr/testify/assert"
)

func TestMain(m *testing.M) {
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

