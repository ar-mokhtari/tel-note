package cli

import (
	"fmt"
	"github.com/ar-mokhtari/tel-note/env"
	"github.com/ar-mokhtari/tel-note/services/identity"
)

func Identity() {
	var userStatus string
	fmt.Println("user or regulator? (notInputAnyThing/reg)")
	fmt.Scanln(&userStatus)
	identity.Identity(userStatus, env.RegulatorString)
	//If user is regulator, maybe want to change greeting note
	if identity.IsRegulator {
		ChangeGreeting()
	}

}
