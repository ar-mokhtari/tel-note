package www

import (
	"net/http"
	"tel-note/config"
)

func Init() {
	fs := http.FileServer(http.Dir(config.MainPath + "/UI"))
	http.Handle("/UI/", http.StripPrefix("/UI/", fs))
	http.HandleFunc("/", WWW.serveTemplate)
}
