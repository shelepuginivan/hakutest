package attachment

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsURL(t *testing.T) {
	cases := []struct {
		src string
		exp bool
	}{
		{src: "/etc/somefile", exp: false},
		{src: "https://example.com/", exp: true},
		{src: "https://example.com/some/path/to/image.jpeg", exp: true},
		{src: "data:image/png;base64,somebase64stuff", exp: false},
		{src: "", exp: false},
	}

	for _, c := range cases {
		assert.Equal(t, c.exp, IsURL(c.src))
	}
}
