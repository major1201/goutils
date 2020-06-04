package goutils

// PanicError panic if err is not nil
func PanicError(err error) {
	if err != nil {
		panic(err)
	}
}
