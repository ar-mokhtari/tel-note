package cli

import (
	"fmt"
	"tel-note/internal/services/identity"
)

func Identity() {
	var userStatus string
	fmt.Println("user or regulator? (notInputAnyThing/reg)")
	fmt.Scanln(&userStatus)
	identity.Identity(userStatus, RegulatorString)
}
