package main

import (
	"tel-note/cli"
	"tel-note/internal/config"
	"tel-note/internal/service/identity"
)

func main() {
	//serve port 1212
	//cli.Serv()

	//To know level of user (admin | normal user)
	cli.Identity()

	//If user is admin, maybe want to change greeting note
	if identity.IsAdmin {
		cli.ChangeGreeting()
	}

	//Show greeting note and description
	cli.ShowGreeting()

	//create global MainData
	config.Init()

	//show contact menu
	cli.ShowMenu()
}
