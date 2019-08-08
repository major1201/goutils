package goutils

// ToBool returns the boolean value represented by the string, default: false
func ToBool(s string) bool {
	return ToBoolDv(s, false)
}

// ToBoolDv returns the boolean value represented by the string with a default bool value
func ToBoolDv(s string, dv bool) bool {
	switch s {
	case "1", "t", "T", "true", "TRUE", "True", "on", "ON", "On":
		return true
	case "0", "f", "F", "false", "FALSE", "False", "off", "OFF", "Off":
		return false
	default:
		return dv
	}
}
