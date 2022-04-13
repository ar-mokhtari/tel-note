package menu

import (
	"encoding/json"
	"net/http"
	"tel-note/env"
)

type menuDuty struct{}

var Duty menuDuty

func (md *menuDuty) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET")
	w.Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
	json.NewEncoder(w).Encode(env.GlobalMenu.Group)
}
