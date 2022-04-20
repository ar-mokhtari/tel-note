package customer

import (
	"encoding/json"
	"net/http"
	"tel-note/env"
	"tel-note/protocol"
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
