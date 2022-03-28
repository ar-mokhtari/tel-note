package validator

import "regexp"

var (
	regexNationalID = regexp.MustCompile("^[\\d]{10}$")
	regexNumber     = regexp.MustCompile("[0-9]+")
)

func IsLuhnAlgorithm(input string) bool {
	if regexNationalID.MatchString(input) {
		return true
	}
	return false
}

func IsNumber(input string) bool {
	if regexNumber.MatchString(input) {
		return true
	}
	return false
}

func ISNil(input []string) []string {
	if input == nil {
		return []string{""}
	}
	return input
}
