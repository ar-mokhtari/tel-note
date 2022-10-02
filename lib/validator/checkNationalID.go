package validator

import (
	"strconv"
)

func CheckIsComplex(input interface{}) bool {
	return RegexRule(`^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`)(input)
}

func CheckNationalID(inputChar string) bool {
	IsLuhnAlgorithm(inputChar)
	var (
		counter, uintChar uint
		controlNumber, _  = strconv.Atoi(string(inputChar[9]))
	)
	for index := 1; index < len(inputChar)-1; index++ {
		internalUintChar, _ := strconv.Atoi(string(inputChar[index]))
		uintChar = uint(internalUintChar)
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
