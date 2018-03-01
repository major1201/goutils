package goutils

import (
	"testing"
	"github.com/stretchr/testify/assert"
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
