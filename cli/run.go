package cli

import (
	"tel-note/internal/config"
	"tel-note/internal/services/identity"
)

func RunApp() {
	//serve port 1212
	//cli.Serv()

	//To know level of user (admin | normal user)
	Identity()

	//If user is admin, maybe want to change greeting note
	if identity.IsAdmin {
		ChangeGreeting()
	}

	//Show greeting note and description
	ShowGreeting()

	//create global MainData
	config.Init()

	//show contact menu
	ShowMenu()
}
