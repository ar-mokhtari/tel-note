package customer

import (
	"encoding/json"
	"net/http"
	"tel-note/env"
	"tel-note/lib/convertor"
	"tel-note/protocol"
)

type (
	findCustomerGroupID struct{}
	findResponse        struct {
		State   uint
		Message string
	}
)

var FindCustomerGroupID findCustomerGroupID

func (fcg *findCustomerGroupID) Do(ID uint) protocol.CustomerStorage {
	return storageRelation.FindCustomerByGroupID(ID)
}

func (fcg findCustomerGroupID) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var res findResponse
	switch r.Method {
	case env.GetMethod:
		if err, id := convertor.StrToUint(r.FormValue("custID")); err != nil {
			res.EncoderJson(w, editResponse{400, "customer id not found"})
		} else {
			result := fcg.Do(id)
			json.NewEncoder(w).Encode(struct {
				State uint
				Data  protocol.CustomerStorage
			}{State: 200, Data: result})
		}
	default:
		res.EncoderJson(w, editResponse{400, "method not support"})
	}
}

func (fcg *findResponse) EncoderJson(w http.ResponseWriter, output editResponse) {
	json.NewEncoder(w).Encode(output)
}
