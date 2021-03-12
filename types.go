package goutils

import (
	"context"
	"reflect"
	"unsafe"
)

var (
	reflectTypeSetForNilValidation = func() map[reflect.Kind]struct{} {
		emptyStruct := struct{}{}
		kindSet := make(map[reflect.Kind]struct{})
		kindSet[reflect.Chan] = emptyStruct
		kindSet[reflect.Func] = emptyStruct
		kindSet[reflect.Interface] = emptyStruct
		kindSet[reflect.Map] = emptyStruct
		kindSet[reflect.Slice] = emptyStruct
		kindSet[reflect.Ptr] = emptyStruct
		kindSet[reflect.UnsafePointer] = emptyStruct
		return kindSet
	}()
)

// IsValueNil returns if reflect.Value is nil.
// panics when value is invalid.
func IsValueNil(value reflect.Value) bool {
	if !value.IsValid() {
		panic("value is invalid")
	}
	kind := value.Kind()
	if _, exists := reflectTypeSetForNilValidation[kind]; exists {
		return value.IsNil()
	}
	return false
}

// IsNil checks if a specified object is nil or not, without Failing.
func IsNil(object interface{}) bool {
	if object == nil {
		// object is interface
		// its type and data are both nil
		return true
	}
	// object has type, so we check its value
	return IsValueNil(reflect.ValueOf(object))
}

// IsTypeFunction returns if the argument is reflect.Func.
func IsTypeFunction(typ reflect.Type) bool {
	return typ.Kind() == reflect.Func
}

// IsTypeIteratee returns if the argument is an iteratee.
func IsTypeIteratee(typ reflect.Type) bool {
	kind := typ.Kind()
	return kind == reflect.Array || kind == reflect.Slice || kind == reflect.Map
}

// IsTypeCollection returns if the argument is a slice/array.
func IsTypeCollection(typ reflect.Type) bool {
	kind := typ.Kind()
	return kind == reflect.Array || kind == reflect.Slice
}

// IsTypeSlice returns if argument is slice type
func IsTypeSlice(typ reflect.Type) bool {
	return typ.Kind() == reflect.Slice
}

// IsTypeArray returns if argument is array type
func IsTypeArray(typ reflect.Type) bool {
	return typ.Kind() == reflect.Array
}

// IsTypeMap returns if the argument is a map
func IsTypeMap(typ reflect.Type) bool {
	return typ.Kind() == reflect.Map
}

var errorType = reflect.TypeOf((*error)(nil)).Elem()

// IsTypeImplementsError returns if the argument is impl of error
func IsTypeImplementsError(typ reflect.Type) bool {
	return typ.Implements(errorType)
}

var ctxType = reflect.TypeOf((*context.Context)(nil)).Elem()

// IsTypeImplementsContext returns if the argument is impl of context.Context
func IsTypeImplementsContext(typ reflect.Type) bool {
	return typ.Implements(ctxType)
}

// ValuesToPtrs transfer []Type to []*Type.
// Panics When:
//  1. s is not slice (array is not allowed);
//  2. elem is pointer type;
//  3. elem is unaddressable value
func ValuesToPtrs(s interface{}) interface{} {
	sVal := reflect.ValueOf(s)
	sKind := sVal.Kind()
	if sKind != reflect.Slice {
		panic("s is not slice")
	}

	elemType := sVal.Type().Elem()
	if elemType.Kind() == reflect.Ptr {
		panic("elem is pointer type")
	}

	sLen := sVal.Len()
	rType := reflect.SliceOf(reflect.PtrTo(elemType))
	if sLen == 0 {
		return reflect.Zero(rType).Interface()
	}

	result := reflect.MakeSlice(rType, 0, sLen)
	for index := 0; index < sLen; index++ {
		result = reflect.Append(result, sVal.Index(index).Addr())
	}
	return result.Interface()
}

// UnsafeBytesToStr returns the byte slice as a volatile string
// THIS IS EVIL CODE.
// YOU HAVE BEEN WARNED.
func UnsafeBytesToStr(b []byte) string {
	// same as strings.Builder::String()
	return *(*string)(unsafe.Pointer(&b))
}

// UnsafeStrToBytes returns the string as a byte slice
// THIS IS EVIL CODE.
// YOU HAVE BEEN WARNED.
func UnsafeStrToBytes(s string) []byte {
	return *(*[]byte)(unsafe.Pointer(&reflect.SliceHeader{
		Len:  len(s),
		Cap:  len(s),
		Data: (*reflect.StringHeader)(unsafe.Pointer(&s)).Data,
	}))
}
