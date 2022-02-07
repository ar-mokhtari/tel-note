package identity

import (
	"fmt"
	"tel-note/internal/config"
)

var IsAdmin bool

func Identity() {
	var userStatus string
	fmt.Println("user or admin?")
	fmt.Scanln(&userStatus)
	if userStatus == string(config.Admin) {
		IsAdmin = true
	}
}
