package cli

import (
	"tel-note/internal/services"
	"tel-note/internal/services/globalVars"
	"tel-note/internal/services/identity"
)

func RunApp() {
	//serve port 1212
	//cli.Serv()

	//To know level of user (regulator | normal user)
	Identity()

	//If user is regulator, maybe want to change greeting note
	if identity.IsRegulator {
		ChangeGreeting()
	}

	//Show greeting note and description
	ShowGreeting()

	//create global variables
	globalVars.Init()

	//create global services
	services.Init()

	//show contact menu
	ShowMenu()
}
