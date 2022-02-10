package main

import (
	"tel-note/cli"
	"tel-note/internal/service/identity"
	"tel-note/internal/storage"
)

func main() {
	//serve port 1212
	cli.Serv()
	//To know level of user (admin | normal user)
	cli.Identity()
	//If user is admin, maybe want to change greeting note
	if identity.IsAdmin {
		cli.ChangeGreeting()
	}
	//Show greeting note and description
	cli.ShowGreeting()
	//show contact menu
	cli.ShowMenu(new(storage.AllContact), new(storage.AllCities))
}
