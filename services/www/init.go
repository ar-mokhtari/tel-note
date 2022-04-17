package www

import "net/http"

func Init() {
	fs := http.FileServer(http.Dir("./UI"))
	http.Handle("/UI/", http.StripPrefix("/UI/", fs))
	http.HandleFunc("/", WWW.serveTemplate)
}
