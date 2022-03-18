package cli

import (
	"encoding/json"
	"net/http"
)

func menuHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(GlobalMenu.Group)
}

func MenuHTTPServe() {
	http.Handle("/menu-list", http.HandlerFunc(menuHandle))
}
