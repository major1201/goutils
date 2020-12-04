package goutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestToInt(t *testing.T) {
	ta := assert.New(t)
	ta.Zero(ToInt("abc"))
	ta.Equal(0, ToInt("0"))
	ta.Equal(123, ToInt("123"))
	ta.Equal(123, ToInt("0123"))
	ta.Equal(-1, ToInt("-1"))
}

func TestToIntDv(t *testing.T) {
	ta := assert.New(t)
	ta.Equal(7, ToIntDv("abc", 7))
	ta.Equal(0, ToIntDv("0", 7))
	ta.Equal(123, ToIntDv("123", 7))
	ta.Equal(123, ToIntDv("0123", 7))
	ta.Equal(-1, ToIntDv("-1", 7))
}

func TestFileSize(t *testing.T) {
	ta := assert.New(t)
	ta.Equal("0 B", FileSize(0))
	ta.Equal("1023 B", FileSize(1023))
	ta.Equal("1.0 KB", FileSize(1024))
	ta.Equal("1.0 MB", FileSize(1024*1024))
}

func TestMakeRange(t *testing.T) {
	ta := assert.New(t)

	ta.Nil(MakeRange(2, 1))
	ta.Equal([]int{1, 2, 3, 4, 5}, MakeRange(1, 5))
	ta.Equal([]int{2}, MakeRange(2, 2))
	ta.Equal([]int{-3, -2, -1, 0, 1, 2}, MakeRange(-3, 2))
}
