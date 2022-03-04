package cli

import (
	"tel-note/services"
)

func RunApp() {
	//serve port 1212
	//cli.Serv()

	//To know level of user (regulator | normal user)
	Identity()

	//Show greeting note and description
	ShowGreeting()

	//create global services
	services.Init()

	//show contact menu
	ShowMenu()
}
