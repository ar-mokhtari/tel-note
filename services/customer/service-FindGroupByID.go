package customer

import (
	"encoding/json"
	"github.com/ar-mokhtari/tel-note/env"
	"github.com/ar-mokhtari/tel-note/lib/convertor"
	"github.com/ar-mokhtari/tel-note/protocol"
	"net/http"
)

type (
	findGroupByID struct{}
)

var FindGroupByID findGroupByID

func (fgi *findGroupByID) Do(ID uint) protocol.CustomerGroup {
	return storageGroup.FindGroupByID(ID)
}

func (fgi *findGroupByID) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case env.GetMethod:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if err, id := convertor.StrToUint(r.FormValue("groupID")); err != nil {
			return
		} else {
			result := fgi.Do(id)
			json.NewEncoder(w).Encode(struct {
				State uint
				Data  protocol.CustomerGroup
			}{200, result})
		}
	default:
		json.NewEncoder(w).Encode(struct {
			State uint
			Data  string
		}{400, "method not support"})
	}
}
