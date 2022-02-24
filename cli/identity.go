package cli

import (
	"fmt"
	"tel-note/services/identity"
)

func Identity() {
	var userStatus string
	fmt.Println("user or regulator? (notInputAnyThing/reg)")
	fmt.Scanln(&userStatus)
	identity.Identity(userStatus, RegulatorString)
	//If user is regulator, maybe want to change greeting note
	if identity.IsRegulator {
		ChangeGreeting()
	}

}
