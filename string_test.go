package goutils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	ta := assert.New(t)
	ta.True(IsEmpty(""))
	ta.False(IsEmpty(" "))
	ta.False(IsEmpty("golang"))
}

func TestIsNotEmpty(t *testing.T) {
	ta := assert.New(t)
	ta.False(IsNotEmpty(""))
	ta.True(IsNotEmpty(" "))
	ta.True(IsNotEmpty("golang"))
}

func TestIsBlank(t *testing.T) {
	ta := assert.New(t)
	ta.True(IsBlank(""))
	ta.True(IsBlank(" "))
	ta.True(IsBlank("\n"))
	ta.True(IsBlank("\t"))
	ta.False(IsBlank("golang"))
}

func TestIsNotBlank(t *testing.T) {
	ta := assert.New(t)
	ta.False(IsNotBlank(""))
	ta.False(IsNotBlank(" "))
	ta.False(IsNotBlank("\n"))
	ta.False(IsNotBlank("\t"))
	ta.True(IsNotBlank("golang"))
}

func TestTrim(t *testing.T) {
	ta := assert.New(t)
	ta.Equal("", Trim(""))
	ta.Equal("", Trim(" "))
	ta.Equal("", Trim("  "))
	ta.Equal("golang", Trim("golang "))
	ta.Equal("golang", Trim(" golang"))
	ta.Equal("golang", Trim(" golang "))
	ta.Equal("golang", Trim("  golang  "))
	ta.NotEqual("golang", Trim("  golang \n"))
}

func TestTrimLeft(t *testing.T) {
	ta := assert.New(t)
	ta.Equal("", TrimLeft(""))
	ta.Equal("", TrimLeft(" "))
	ta.Equal("", TrimLeft("  "))
	ta.Equal("golang ", TrimLeft("golang "))
	ta.Equal("golang", TrimLeft(" golang"))
	ta.Equal("golang ", TrimLeft(" golang "))
	ta.Equal("golang  ", TrimLeft("  golang  "))
}

func TestTrimRight(t *testing.T) {
	ta := assert.New(t)
	ta.Equal("", TrimRight(""))
	ta.Equal("", TrimRight(" "))
	ta.Equal("", TrimRight("  "))
	ta.Equal("golang", TrimRight("golang "))
	ta.Equal(" golang", TrimRight(" golang"))
	ta.Equal(" golang", TrimRight(" golang "))
	ta.Equal("  golang", TrimRight("  golang  "))
}

func TestLeftPad(t *testing.T) {
	ta := assert.New(t)
	ta.Equal("golang", LeftPad("golang", "-", 0))
	ta.Equal("golang", LeftPad("golang", "-", 6))
	ta.Equal("----golang", LeftPad("golang", "-", 10))
	ta.Equal("", LeftPad("", "-", 0))
	ta.Equal("-", LeftPad("", "-", 1))
}

func TestRightPad(t *testing.T) {
	ta := assert.New(t)
	ta.Equal("golang", RightPad("golang", "-", 0))
	ta.Equal("golang", RightPad("golang", "-", 6))
	ta.Equal("golang----", RightPad("golang", "-", 10))
	ta.Equal("", RightPad("", "-", 0))
	ta.Equal("-", RightPad("", "-", 1))
}

func TestZeroFill(t *testing.T) {
	ta := assert.New(t)
	ta.Equal("00", ZeroFill("", 2))
	ta.Equal("00", ZeroFill("0", 2))
	ta.Equal("01", ZeroFill("1", 2))
	ta.Equal("11", ZeroFill("11", 2))
	ta.Equal("111", ZeroFill("111", 2))
}

func TestLen(t *testing.T) {
	ta := assert.New(t)
	ta.Equal(0, Len(""))
	ta.Equal(5, Len("Hello"))
	ta.Equal(2, Len("你好"))
	ta.Equal(9, Len("Hello, 世界"))
}

func TestIndex(t *testing.T) {
	ta := assert.New(t)
	ta.Equal(0, Index("Hello, world!", ""))
	ta.Equal(0, Index("Hello, world!", "He"))
	ta.Equal(2, Index("Hello, world!", "llo"))
	ta.Equal(-1, Index("Hello, world!", "not exist"))
	ta.Equal(0, Index("Hello, 世界!", "He"))
	ta.Equal(7, Index("Hello, 世界!", "世"))
	ta.Equal(7, Index("Hello, 世界!", "世界"))
	ta.Equal(-1, Index("Hello, 世界!", "不存在"))
}

func TestBetween(t *testing.T) {
	ta := assert.New(t)
	ta.Equal("", Between("abcdefgh", "not", "exist"))
	ta.Equal("def", Between("abcdefgh", "bc", "gh"))
	ta.Equal("bcde", Between("abcdefgh", "a", "f"))
	ta.Equal("的世", Between("你好我的世界", "好我", "界"))
}
