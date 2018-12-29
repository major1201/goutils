package goutils

import (
	"fmt"
	"math"
	"strconv"
)

// ToInt safely parse a string to a integer with a default value 0 when parse fails
func ToInt(s string) int {
	return ToIntDv(s, 0)
}

// ToIntDv safely parse a string to a integer with a specified default value
func ToIntDv(s string, dv int) int {
	i, e := strconv.Atoi(s)
	if e != nil {
		return dv
	}
	return i
}

func humanReadableBytes(s uint64, base float64, units []string) string {
	if s < 10 {
		return fmt.Sprintf("%d %s", s, units[0])
	}
	e := math.Floor(math.Log(float64(s)) / math.Log(base))
	suffix := units[int(e)]
	val := float64(s) / math.Pow(base, math.Floor(e))
	f := "%.0f"
	if val < 10 {
		f = "%.1f"
	}

	return fmt.Sprintf(f+" %s", val, suffix)
}

// FileSize translate bytes number into human-readable size
func FileSize(s uint64) string {
	units := []string{"B", "KB", "MB", "GB", "TB", "PB", "EB"}
	return humanReadableBytes(s, 1024, units)
}
