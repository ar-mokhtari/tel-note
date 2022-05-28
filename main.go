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
	mux := http.NewServeMux()
	//create global services
	services.Init(mux)
	switch config.RunAppType {
	case "serv":
		srv := &http.Server{
			Handler:      mux,
			Addr:         ":1212",
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

func main() {

}
