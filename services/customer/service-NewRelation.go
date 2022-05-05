package customer

import (
	"encoding/json"
	"github.com/ar-mokhtari/tel-note/env"
	"github.com/ar-mokhtari/tel-note/lib/convertor"
	"net/http"
)

type (
	newGrpRelation struct{}
)

var NewGrpRelation newGrpRelation

func (ng *newGrpRelation) Do(customerID uint, groupID uint) {
	storageRelation.NewRelation(customerID, groupID)
}

func (ng *newGrpRelation) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case env.PostMethod:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		if err, customerID := convertor.StrToUint(r.FormValue("customerID")); err != nil {
			return
		} else {
			if err, groupID := convertor.StrToUint(r.FormValue("groupID")); err != nil {
				return
			} else {
				ng.Do(customerID, groupID)
				json.NewEncoder(w).Encode(struct {
					State uint
					Data  string
				}{200, "new group relation added"})
			}
		}
	default:
		json.NewEncoder(w).Encode(struct {
			State uint
			Data  string
		}{400, "method not support"})
	}
}
