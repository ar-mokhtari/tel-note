package sex

import (
	"encoding/json"
	"net/http"
	"tel-note/env"
	"tel-note/protocol"
)

type getSex struct{}

var GetSex getSex

func (gs *getSex) Do() []*protocol.Sex {
	return storage.GetSex()
}

func (gs *getSex) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case env.GetMethod:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		var result = gs.Do()
		json.NewEncoder(w).Encode(struct {
			State uint
			Data  []*protocol.Sex
		}{200, result})
	default:
		json.NewEncoder(w).Encode(struct {
			State   uint
			Message string
		}{400, "method not support"})
	}
}
