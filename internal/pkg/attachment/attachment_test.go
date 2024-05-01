package attachment

import (
	"os"
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

func TestNew(t *testing.T) {
	bytes := []byte("some text")

	tmp, err := os.CreateTemp("", "*.txt")
	if err != nil {
		panic(err)
	}

	_, err = tmp.Write(bytes)
	if err != nil {
		panic(err)
	}

	defer func() {
		err = tmp.Close()
		if err != nil {
			panic(err)
		}

		err = os.Remove(tmp.Name())
		if err != nil {
			panic(err)
		}
	}()

	cases := []struct {
		act *Attachment
		exp *Attachment
	}{
		{
			act: New(
				"Some name",
				AttachmentFile,
				"https://example.com/file.pdf",
			),
			exp: &Attachment{
				Name: "Some name",
				Type: AttachmentFile,
				Src:  "https://example.com/file.pdf",
			},
		},
		{
			act: New(
				"Another",
				AttachmentImage,
				tmp.Name(),
			),
			exp: &Attachment{
				Name: "Another",
				Type: AttachmentImage,
				Src:  "data:text/plain; charset=utf-8;base64,c29tZSB0ZXh0",
			},
		},
	}

	for _, c := range cases {
		assert.Equal(t, c.exp.Name, c.act.Name)
		assert.Equal(t, c.exp.Type, c.act.Type)
		assert.Equal(t, c.exp.Src, c.act.Src)
	}
}
