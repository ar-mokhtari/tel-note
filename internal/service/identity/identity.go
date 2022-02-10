package identity

import (
	"strings"
	"tel-note/internal/config"
)

var IsAdmin bool

func Identity(userStatus string) {
	if strings.ToLower(userStatus) == string(config.AdminString) {
		IsAdmin = true
	}
}
