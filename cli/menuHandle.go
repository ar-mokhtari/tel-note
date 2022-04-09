package cli

import (
	"encoding/json"
	"net/http"
)

func menuHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	json.NewEncoder(w).Encode(GlobalMenu.Group)
}

func MenuHTTPServe() {
	http.Handle("/menu-list", http.HandlerFunc(menuHandle))
}
