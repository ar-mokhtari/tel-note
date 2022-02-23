package cli

import (
	"tel-note/internal/services"
	"tel-note/internal/services/globalVars"
)

func RunApp() {
	//serve port 1212
	//cli.Serv()

	//To know level of user (regulator | normal user)
	Identity()

	//Show greeting note and description
	ShowGreeting()

	//create|allocate global variables
	globalVars.Init()

	//create global services
	services.Init()

	//show contact menu
	ShowMenu()
}
