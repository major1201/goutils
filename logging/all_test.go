package logging

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

var logger = New("TEST")

func addStdout() {
	if len(Writers) == 0 {
		AddStdout(0)
	}
}

func TestLogger_Debug(t *testing.T) {
	addStdout()
	assert.NotPanics(t, func() {
		logger.Debug("golang log debug")
	})
}

func TestLogger_Info(t *testing.T) {
	addStdout()
	assert.NotPanics(t, func() {
		logger.Info("golang log info")
	})
}

func TestLogger_Warning(t *testing.T) {
	addStdout()
	assert.NotPanics(t, func() {
		logger.Warning("golang log warning")
	})
}

func TestLogger_Error(t *testing.T) {
	addStdout()
	assert.NotPanics(t, func() {
		logger.Error("golang log error")
	})
}
