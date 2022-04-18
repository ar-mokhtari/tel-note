package www

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"tel-note/config"
)

type www struct{}

var WWW www

func (ww www) serveTemplate(w http.ResponseWriter, r *http.Request) {
	lp := filepath.Join(config.MainPath, "UI", "layout.html")
	fp := filepath.Join(config.MainPath, "UI", filepath.Clean(r.URL.Path))

	// Return a 404 if the template doesn't exist
	info, err := os.Stat(fp)
	if err != nil {
		if os.IsNotExist(err) {
			http.NotFound(w, r)
			return
		}
	}

	// Return a 404 if the request is for a directory
	if info.IsDir() {
		http.NotFound(w, r)
		return
	}

	tmpl, err := template.ParseFiles(lp, fp)
	if err != nil {
		// Log the detailed error
		log.Println(err.Error())
		// Return a generic "Internal Server Error" message
		http.Error(w, http.StatusText(500), 500)
		return
	}

	err = tmpl.ExecuteTemplate(w, "layout", nil)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, http.StatusText(500), 500)
	}
}
