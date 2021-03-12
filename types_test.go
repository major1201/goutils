package goutils

import (
	"context"
	"reflect"
	"testing"
	"time"
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

type tmpCtx int

func (*tmpCtx) Deadline() (deadline time.Time, ok bool) {
	return
}

func (*tmpCtx) Done() <-chan struct{} {
	return nil
}

func (*tmpCtx) Err() error {
	return nil
}

func (*tmpCtx) Value(key interface{}) interface{} {
	return nil
}

func (e *tmpCtx) String() string {
	return "tmpCtx"
}

type tmpStruct struct {
	msg string
}

type tmpError struct {
	msg string
}

func (e *tmpError) Error() string {
	return e.msg
}

func TestUnsafe(t *testing.T) {
	is := assert.New(t)

	ss := []string{"a", "b", "c"}
	mp := map[int]int{1: 2, 3: 4}

	var mpNil map[int]int
	is.PanicsWithValue("value is invalid", func() { IsValueNil(reflect.ValueOf(nil)) })
	is.True(IsValueNil(reflect.ValueOf(mpNil)))
	is.False(IsValueNil(reflect.ValueOf(123)))

	fn := func(a, b int) int {
		return a + b
	}
	is.True(IsTypeFunction(reflect.TypeOf(fn)))
	is.False(IsTypeFunction(reflect.TypeOf(mpNil)))
	is.False(IsTypeFunction(reflect.TypeOf(ss)))

	is.True(IsTypeCollection(reflect.TypeOf(ss)))
	is.False(IsTypeCollection(reflect.TypeOf(mp)))

	is.True(IsTypeMap(reflect.TypeOf(mp)))
	is.False(IsTypeMap(reflect.TypeOf(ss)))

	is.True(IsTypeImplementsError(reflect.TypeOf(&tmpError{"hello"})))
	is.False(IsTypeImplementsError(reflect.TypeOf(tmpError{"hello"})))
	is.False(IsTypeImplementsError(reflect.TypeOf(&tmpStruct{"hello"})))
	is.False(IsTypeImplementsError(reflect.TypeOf(tmpStruct{"hello"})))

	ctx := context.Background()
	is.True(IsTypeImplementsContext(reflect.TypeOf(ctx)))

	tCtx := new(tmpCtx)
	is.True(IsTypeImplementsContext(reflect.TypeOf(tCtx)))
	is.True(IsTypeImplementsContext(reflect.TypeOf(ctx)))
	is.False(IsTypeImplementsContext(reflect.TypeOf(&tmpStruct{"hello"})))
	is.False(IsTypeImplementsContext(reflect.TypeOf(tmpStruct{"hello"})))
}

type testStruct struct {
	A string
}

func TestValuesToPtrs(t *testing.T) {
	is := assert.New(t)

	is.PanicsWithValue(
		"s is not slice",
		func() {
			ValuesToPtrs(1)
		},
	)
	is.PanicsWithValue(
		"elem is pointer type",
		func() {
			modelPtrs := []*testStruct{{"1"}, {"2"}}
			ValuesToPtrs(modelPtrs)
		},
	)

	intPtrs := ValuesToPtrs([]int{}).([]*int)
	is.Empty(intPtrs)

	intVals := []int{1, 2, 3, 4, 5}
	intPtrs = ValuesToPtrs([]int{1, 2, 3, 4, 5}).([]*int)
	for index, intPtr := range intPtrs {
		is.Equal(intVals[index], *intPtr)
	}

	models := []testStruct{
		{"1"},
		{"2"},
		{"3"},
	}
	modelPtrs := ValuesToPtrs(models).([]*testStruct)
	for index, modelPtr := range modelPtrs {
		is.Equal(models[index], *modelPtr)
	}
}

var (
	testStr   = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	testBytes = []byte("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func TestUnsafeBytesToStr(t *testing.T) {
	s := UnsafeBytesToStr(testBytes)
	assert.Equal(t, testStr, s)
}

func TestUnsafeStrToBytes(t *testing.T) {
	bs := UnsafeStrToBytes(testStr)
	assert.Equal(t, true, reflect.DeepEqual(bs, testBytes))
}

func BenchmarkBytesToStr(b *testing.B) {
	b.Run(
		"BytesToStr-UnSafe",
		func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_ = UnsafeBytesToStr(testBytes)
			}
		},
	)

	b.Run(
		"BytesToStr-Safe",
		func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_ = string(testBytes)
			}
		},
	)
}

func BenchmarkStrToBytes(b *testing.B) {
	b.Run(
		"StrToBytes-UnSafe",
		func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_ = UnsafeStrToBytes(testStr)
			}
		},
	)

	b.Run(
		"StrToBytes-Safe",
		func(b *testing.B) {
			b.ResetTimer()
			b.ReportAllocs()
			for i := 0; i < b.N; i++ {
				_ = []byte(testStr)
			}
		},
	)
}
