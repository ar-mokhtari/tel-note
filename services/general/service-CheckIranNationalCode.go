package general

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
)

type checkIranNationalCode struct{}

var CheckIranNational *checkIranNationalCode

var regexGatewayID = regexp.MustCompile("^[\\d]{10}$")

func CheckIranNationalCode(inputChar string) bool {
	//TODO::: cleaning
	if !regexGatewayID.MatchString(inputChar) {
		return false
	}
	var (
		counter, uintChar uint
		controlNumber, _  = strconv.Atoi(string(inputChar[9]))
	)
	for index := 1; index < len(inputChar)-1; index++ {
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

func (iranNationalCode *checkIranNationalCode) ServeCheckIranNationalCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	NationalCode := r.Header.Values("NID")
	fmt.Fprintf(w, "%v", CheckIranNationalCode(NationalCode[0]))
}
