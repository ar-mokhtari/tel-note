package main

import (
	"log"
	"net/http"
	"tel-note/cli"
	"tel-note/config"
	"tel-note/services"
	"time"
)

func init() {
	//create global services
	services.Init()
}

func main() {
	switch config.RunAppType {
	case "serv":
		srv := &http.Server{
			Handler:      nil,
			Addr:         "127.0.0.1:1212",
			WriteTimeout: 30 * time.Second,
			ReadTimeout:  30 * time.Second,
		}
		//serve port 1212
		//listen for request:
		log.Fatalln(srv.ListenAndServe())
	default:
		cli.RunCli()
	}
}
