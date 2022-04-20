package customer

import (
	"encoding/json"
	"net/http"
	"tel-note/env"
	"tel-note/lib/convertor"
	"tel-note/protocol"
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
