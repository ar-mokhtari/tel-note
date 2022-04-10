package cli

import (
	"fmt"
)

func RunCli() {
	//To know level of user (regulator | normal user)
	Identity()

	//Show greeting note and description
	ShowGreeting()

	//show contact menu
	fmt.Println(ShowMenu())

	//run menu
	RunMenu()
}
