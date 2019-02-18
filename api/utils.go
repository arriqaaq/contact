package api

import (
	"strconv"
)

func StringToUInt(in string) (uint, error) {
	// to convert a string number to a uint
	val, err := strconv.Atoi(in)
	return uint(val), err
}

func isEmptyStr(in string) bool {
	if in == "" {
		return true
	}
	return false
}
