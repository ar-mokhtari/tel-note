package main

import (
	"TelNote/internal/cli"
	"TelNote/internal/service/greeting"
	"TelNote/internal/service/identity"
	"TelNote/internal/storage"
)

func main() {
	var MainData = new(storage.AllContact)
	//To know level of user (admin | normal user)
	identity.Identity()
	//If user is admin, maybe want to change greeting note
	if identity.IsAdmin {
		greeting.ChangeGreeting()
	}
	//Show greeting note and description
	greeting.Greeting()
	cli.ShowMenu(MainData)
}
