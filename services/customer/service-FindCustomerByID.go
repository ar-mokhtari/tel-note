package customer

import (
	"encoding/json"
	"net/http"
	"tel-note/env"
	"tel-note/lib/convertor"
	"tel-note/protocol"
)

type (
	findCustomerByID struct{}
)

var FindCustomerByID findCustomerByID

func (fci *findCustomerByID) Do(ID uint) protocol.Customer {
	return storage.FindCustomerByID(ID)
}

func (fci *findCustomerByID) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case env.GetMethod:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if err, id := convertor.StrToUint(r.FormValue("custID")); err != nil {
			return
		} else {
			result := fci.Do(id)
			json.NewEncoder(w).Encode(struct {
				State uint
				Data  protocol.Customer
			}{200, result})
		}
	default:
		json.NewEncoder(w).Encode(struct {
			State uint
			Data  string
		}{400, "method not support"})
	}
}
