package identity

import (
	"strings"
)

var IsAdmin bool

func Identity(userStatus,AdminString string) {
	if strings.ToLower(userStatus) == AdminString {
		IsAdmin = true
	}
}
