package goutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestContainsString(t *testing.T) {
	ta := assert.New(t)
	ta.True(ContainsString("golang", "go", "lang", "golang"))
	ta.False(ContainsString("golang", "go", "lang", "golang1"))
}

func TestContainsInt(t *testing.T) {
	ta := assert.New(t)
	ta.True(ContainsInt(1, 1, 2, 3))
	ta.False(ContainsInt(4, 1, 2, 3))
}

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

func TestFilterString(t *testing.T) {
	ta := assert.New(t)
	ta.Equal([]string{"golang", "go"}, FilterString([]string{"golang", "lang", "go"}, func(s string) bool {
		return s != "lang"
	}))
}

func TestFilterEmptyString(t *testing.T) {
	ta := assert.New(t)
	vs := []string{"golang", " ", "  ", ""}
	ta.Equal([]string{"golang", " ", "  "}, FilterEmptyString(vs))
}

func TestFilterBlankString(t *testing.T) {
	ta := assert.New(t)
	vs := []string{"golang", " ", "  ", ""}
	ta.Equal([]string{"golang"}, FilterBlankString(vs))
}

func TestFilterInt(t *testing.T) {
	ta := assert.New(t)
	ta.Equal([]int{1, 3}, FilterInt([]int{1, 2, 3}, func(i int) bool {
		return i != 2
	}))
}

func TestMap(t *testing.T) {
	ta := assert.New(t)
	ta.Equal([]string{"hellogolang", "hellogo", "hellolang", "hello"}, Map([]string{"golang", "go", "lang", ""}, func(s string) string {
		return "hello" + s
	}))
}

func TestTernary(t *testing.T) {
	ta := assert.New(t)
	ta.Equal(1, Ternary(true, 1, 2))
	ta.Equal(2, Ternary(false, 1, 2))
}

func TestDefaultIfNil(t *testing.T) {
	ta := assert.New(t)

	ta.Equal(1, DefaultIfNil(nil, 1))
	ta.Equal(2, DefaultIfNil(2, 1))
	ta.Nil(DefaultIfNil(nil, nil))

	type tt struct{}
	var obj *tt

	ta.Equal(1, DefaultIfNil(obj, 1))
}
