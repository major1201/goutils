package goutils

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPanicError(t *testing.T) {
	assert.Panics(t, func() {
		PanicError(errors.New("some errors"))
	})

	assert.NotPanics(t, func() {
		PanicError(nil)
	})
}
