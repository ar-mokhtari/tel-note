package www

import (
	"github.com/ar-mokhtari/tel-note/config"
	"net/http"
)

func Init() {
	fs := http.FileServer(http.Dir(config.MainPath + "/UI"))
	http.Handle("/UI/", http.StripPrefix("/UI/", fs))
	http.HandleFunc("/", WWW.serveTemplate)
}
