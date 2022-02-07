package identity

import (
	"fmt"
	"strings"
	"tel-note/internal/config"
)

var IsAdmin bool

func Identity() {
	var userStatus string
	fmt.Println("user or admin?")
	fmt.Scanln(&userStatus)
	if strings.ToLower(userStatus) == string(config.AdminString) {
		IsAdmin = true
	}
}
