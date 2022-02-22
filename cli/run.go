package cli

import (
	"tel-note/internal/config"
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

	//create global MainData
	config.Init()

	//create global services
	globalVars.Init()
	services.Init()

	//show contact menu
	ShowMenu()
}
