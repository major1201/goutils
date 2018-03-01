package goutils

import (
	"testing"
	"os"
	"path/filepath"
	"github.com/stretchr/testify/assert"
)

func TestDownload(t *testing.T) {
	url := "https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_150x54dp.png"
	dest := filepath.Join(os.TempDir(), filepath.Base(url))
	ta := assert.New(t)
	ta.NotPanics(func() {
		Download(url, dest)
	})
	ta.FileExists(dest)
}
