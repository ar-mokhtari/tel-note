package cli

import (
	"fmt"
	"tel-note/config"
)

func RunApp() {
	if config.RunAppType == "serv" {
		//serve port 1212
		Serv()
	} else {

		//To know level of user (regulator | normal user)
		Identity()

		//Show greeting note and description
		ShowGreeting()

		//show contact menu
		fmt.Println(ShowMenu())

		//run menu
		RunMenu()
	}
}
