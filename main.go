package main

import (
	"github.com/ar-mokhtari/tel-note/cli"
	"github.com/ar-mokhtari/tel-note/config"
	"github.com/ar-mokhtari/tel-note/services"
	"log"
	"net/http"
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
			WriteTimeout: 5 * time.Second,
			ReadTimeout:  10 * time.Second,
		}
		//serve port 1212
		//listen for request:
		log.Fatalln(srv.ListenAndServe())
	default:
		cli.RunCli()
	}
}
