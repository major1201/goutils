package goutils

import "reflect"

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
