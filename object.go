package goutils

import "reflect"

// Contains tells an object is in a list or not
func Contains(obj interface{}, v ...interface{}) bool {
	for _, o := range v {
		if obj == o {
			return true
		}
	}
	return false
}

// DeepContains tells an object is in a list or not, but uses reflect.DeepEqual to determine the equality
func DeepContains(obj interface{}, v ...interface{}) bool {
	for _, o := range v {
		if reflect.DeepEqual(o, obj) {
			return true
		}
	}
	return false
}

// FilterString filters a string slice with the function returns false
func FilterString(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// FilterEmptyString removes empty strings in a string slice
func FilterEmptyString(vs []string) []string {
	return FilterString(vs, IsNotEmpty)
}

// FilterBlankString removes blank strings in a string slice
func FilterBlankString(vs []string) []string {
	return FilterString(vs, IsNotBlank)
}

// FilterInt filters an int slice with the function returns false
func FilterInt(vs []int, f func(int) bool) []int {
	vsf := make([]int, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// Map maps a string slice with a function
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}
