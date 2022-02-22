package identity

import (
	"strings"
)

var IsRegulator bool

func Identity(userStatus, RegulatorString string) {
	if strings.ToLower(userStatus) == RegulatorString {
		IsRegulator = true
	} else {
		IsRegulator = false
	}
}
