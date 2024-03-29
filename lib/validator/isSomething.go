package validator

import (
	"github.com/dlclark/regexp2"
	"reflect"
	"regexp"
)

var (
	regexNationalID = regexp.MustCompile("^[\\d]{10}$")
	regexNumber     = regexp.MustCompile("[0-9]+")
	complexRexexp   = `^(?=.*[a-z])(?=.*[A-Z])(?=.*\d)(?=.*[@$!%*?&])[A-Za-z\d@$!%*?&]{8,}$`
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

// Returns true if input matches the passed pattern
func RegexRule(pattern string) func(interface{}) bool {
	regex := regexp2.MustCompile(pattern, regexp2.None)
	return func(input interface{}) bool {
		inputValue := reflect.ValueOf(input)
		switch inputValue.Kind() {
		case reflect.String:
			output, _ := regex.MatchString(input.(string))
			return output
		default:
			return false
		}
	}
}

func CheckIsComplex(input interface{}) bool {
	return RegexRule(complexRexexp)(input)
}
