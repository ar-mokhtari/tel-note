package customer

import (
	"encoding/json"
	"net/http"
	"tel-note/env"
	"tel-note/lib/convertor"
	"tel-note/protocol"
)

type (
	findCustGrpRelationByGrpID struct{}
	findResponse               struct {
		State   uint
		Message string
	}
)

var FindCustGrpRelationByGrpID findCustGrpRelationByGrpID

func (fcg *findCustGrpRelationByGrpID) Do(ID uint) protocol.CustomerStorage {
	return storageRelation.FindCustomerGroupRelationByGroupID(ID)
}

func (fcg findCustGrpRelationByGrpID) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var res findResponse
	switch r.Method {
	case env.GetMethod:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if err, id := convertor.StrToUint(r.FormValue("groupID")); err != nil {
			res.EncoderJson(w, editResponse{400, "cust id format wrong"})
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
