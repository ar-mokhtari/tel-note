package main

import (
	"tel-note/cli"
	"tel-note/services"
)

func init() {
	//create global services
	services.Init()
}

func main() {
	cli.RunApp()
}
