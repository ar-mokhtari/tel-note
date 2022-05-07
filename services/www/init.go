package www

import (
	"github.com/ar-mokhtari/tel-note/config"
	"net/http"
)

func Init(mux *http.ServeMux) {
	fs := http.FileServer(http.Dir(config.MainPath + "/UI"))
	mux.Handle("/UI/", http.StripPrefix("/UI/", fs))
	mux.HandleFunc("/", WWW.serveTemplate)
}
