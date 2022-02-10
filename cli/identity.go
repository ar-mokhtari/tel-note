package cli

import (
	"fmt"
	"tel-note/internal/service/identity"
)

func Identity() {
	var userStatus string
	fmt.Println("user or admin?")
	fmt.Scanln(&userStatus)
	identity.Identity(userStatus)
}
