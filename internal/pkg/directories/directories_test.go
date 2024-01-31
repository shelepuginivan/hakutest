package directories

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestConfig(t *testing.T) {
	assert.True(t, strings.HasSuffix(Config(), "hakutest"))
}

func TestData(t *testing.T) {
	dataDir := Data()
	assert.True(t, strings.HasSuffix(dataDir, "hakutest") || dataDir == "data")
}

func TestExecutable(t *testing.T) {
	exeDir := Executable()
	expected, err := os.Executable()

	if err != nil {
		assert.Equal(t, exeDir, ".")
		return
	}

	assert.Equal(t, exeDir, filepath.Dir(expected))
}
