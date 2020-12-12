package goutils

import (
	"github.com/stretchr/testify/assert"
	"os"
	"path/filepath"
	"testing"
)

func TestDownload(t *testing.T) {
	url := "https://www.google.com/images/branding/googlelogo/1x/googlelogo_color_150x54dp.png"
	dest := filepath.Join(os.TempDir(), filepath.Base(url))
	ta := assert.New(t)
	ta.NotPanics(func() {
		ta.NoError(Download(url, dest))
	})
	ta.FileExists(dest)
}
