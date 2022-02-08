package main

import (
	"tel-note/internal/cli"
	"tel-note/internal/service/greeting"
	"tel-note/internal/service/identity"
	"tel-note/internal/storage"
)

func main() {
	//To know level of user (admin | normal user)
	identity.Identity()
	//If user is admin, maybe want to change greeting note
	if identity.IsAdmin {
		greeting.ChangeGreeting()
	}
	//Show greeting note and description
	greeting.Greeting()
	cli.ShowMenu(new(storage.AllContact))
}
