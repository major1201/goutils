package goutils

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestContains(t *testing.T) {
	ta := assert.New(t)
	ta.True(Contains("golang", "go", "lang", "golang"))
	ta.False(Contains("golang", "go", "lang", "golang1"))
	ta.Panics(func() {
		Contains([]string{"golang", "go"}, []string{"go"}, []string{"golang", "go"})
	})
}

func TestDeepContains(t *testing.T) {
	ta := assert.New(t)
	ta.True(DeepContains("golang", "go", "lang", "golang"))
	ta.False(DeepContains("golang", "go", "lang", "golang1"))
	ta.True(DeepContains([]string{"golang", "go"}, []string{"go"}, []string{"golang", "go"}))
}

