package general

import (
	"strconv"
)

func CheckIranNationalCode(inputChar string) bool {
	var lenInputChar int
	if lenInputChar = len(inputChar); lenInputChar != 10 {
		return false
	}
	var (
		counter, uintChar uint
		controlNumber, _  = strconv.Atoi(string(inputChar[9]))
	)
	for index := 1; index < lenInputChar-1; index++ {
		if internalUintChar, err := strconv.Atoi(string(inputChar[index])); err != nil {
			return false
		} else {
			uintChar = uint(internalUintChar)
		}
		counter += uint(10-index) * uintChar
	}
	switch remaining := counter % 11; remaining < 2 {
	case true:
		if remaining == uint(controlNumber) {
			return true
		}
	case false:
		if 11-remaining == uint(controlNumber) {
			return true
		}
	}
	return false
}
