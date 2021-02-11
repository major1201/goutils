package goutils

import (
	"reflect"
	"testing"
	"unsafe"

	"github.com/stretchr/testify/assert"
)

func TestIsValueNil(t *testing.T) {
	ta := assert.New(t)

	var mpNil map[int]int
	ta.PanicsWithValue("value is invalid", func() { IsValueNil(reflect.ValueOf(nil)) })
	ta.True(IsValueNil(reflect.ValueOf(mpNil)))
	ta.False(IsValueNil(reflect.ValueOf(123)))
}

type testNilStruct struct {
}

func (s testNilStruct) Add(a, b int) int {
	return a + b
}

type testNilPtrStruct struct {
}

func (s *testNilPtrStruct) Add(a, b int) int {
	return a + b
}

type testNilInterface interface {
	Add(a, b int) int
}

func TestIsNil(t *testing.T) {
	// basic value
	assert.False(t, IsNil((int8)(123)))
	assert.False(t, IsNil("hello"))

	// interface
	var iFace testNilInterface
	assert.True(t, IsNil(iFace))

	var sp *testNilPtrStruct
	assert.True(t, IsNil(sp))
	iFace = sp
	assert.True(t, IsNil(iFace))

	sp = &testNilPtrStruct{}
	iFace = sp
	assert.False(t, IsNil(sp))
	assert.False(t, IsNil(iFace))

	// struct
	s := testNilStruct{}
	assert.False(t, IsNil(s))
	iFace = s
	assert.False(t, IsNil(iFace))

	// Ptr
	assert.True(t, IsNil((*int8)(nil)))

	// map
	var mm map[int]struct{}
	assert.True(t, IsNil(mm))
	mm = make(map[int]struct{})
	assert.False(t, IsNil(mm))

	// slice
	var ss []int
	assert.True(t, IsNil(ss))
	ss = make([]int, 0)
	assert.False(t, IsNil(ss))

	// chain
	var cc chan int
	assert.True(t, IsNil(cc))
	cc = make(chan int)
	assert.False(t, IsNil(cc))

	// func
	var ff func(a, b int) int
	assert.True(t, IsNil(ff))
	ff = func(a, b int) int {
		return a + b
	}
	assert.False(t, IsNil(ff))

	// unsafe.Pointer
	up := unsafe.Pointer(nil)
	assert.True(t, IsNil(up))
}
