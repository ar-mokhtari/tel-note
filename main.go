package main

import (
	"log"
	"net/http"
	"tel-note/cli"
	"tel-note/config"
	"tel-note/services"
)

func init() {
	//create global services
	services.Init()
}

func main() {
	switch config.RunAppType {
	case "serv":
		//serve port 1212
		//listen for request:
		log.Fatalln(http.ListenAndServe(":1212", nil))
	default:
		cli.RunCli()
	}
}
