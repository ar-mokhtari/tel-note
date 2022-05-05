package customer

import (
	"encoding/json"
	"github.com/ar-mokhtari/tel-note/env"
	"github.com/ar-mokhtari/tel-note/protocol"
	"net/http"
)

type getCustomerGroup struct{}

var GetCustomerGroup getCustomerGroup

func (gcg *getCustomerGroup) Do() protocol.CustomerGroupStorage {
	return storageGroup.GetCustomerGroup()
}

func (gcg *getCustomerGroup) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case env.GetMethod:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		result := gcg.Do()
		json.NewEncoder(w).Encode(struct {
			State uint
			Data  protocol.CustomerGroupStorage
		}{200, result})
	default:
		json.NewEncoder(w).Encode(struct {
			State uint
			Data  string
		}{400, "method not support"})
	}
}
