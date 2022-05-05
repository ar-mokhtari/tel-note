package customer

import (
	"encoding/json"
	"github.com/ar-mokhtari/tel-note/env"
	"github.com/ar-mokhtari/tel-note/protocol"
	"net/http"
)

type getCustomerGroupRelation struct{}

var GetCustomerGroupRelation getCustomerGroupRelation

func (cgr *getCustomerGroupRelation) Do() protocol.CustomerGRelationStorage {
	return storageRelation.GetCustomerGroupRelation()
}

func (cgr *getCustomerGroupRelation) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case env.GetMethod:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		result := cgr.Do()
		json.NewEncoder(w).Encode(struct {
			State uint
			Data  protocol.CustomerGRelationStorage
		}{200, result})
	default:
		json.NewEncoder(w).Encode(struct {
			State uint
			Data  string
		}{400, "method not support"})
	}
}
