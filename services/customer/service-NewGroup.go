package customer

import (
	"encoding/json"
	"github.com/ar-mokhtari/tel-note/env"
	"net/http"
)

type (
	newGroup struct{}
)

var NewGroup newGroup

func (ng *newGroup) Do(groupName string) {
	storageGroup.NewGroup(groupName)
}

func (ng *newGroup) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case env.PostMethod:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		ng.Do(r.FormValue("groupName"))
		json.NewEncoder(w).Encode(struct {
			State uint
			Data  string
		}{200, "new group added"})
	default:
		json.NewEncoder(w).Encode(struct {
			State uint
			Data  string
		}{400, "method not support"})
	}
}
