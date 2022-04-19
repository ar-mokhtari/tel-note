package customer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tel-note/env"
	"tel-note/lib/convertor"
	"tel-note/protocol"
)

type (
	editCustomer struct{}
	EditRequest  protocol.Customer
	editResponse struct {
		State   uint
		Message string
	}
)

var EditCustomer editCustomer

func (ac *editCustomer) Do(id uint, editCustomer EditRequest) (err error) {
	err = storage.EditCustomer(id, protocol.Customer(editCustomer))
	return err
}

func (ac *editCustomer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		req EditRequest
		res editResponse
	)
	switch r.Method {
	case env.PatchMethod:
		if err := req.DecoderJson(*r); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err, id := convertor.StrToUint(r.FormValue("custID")); err != nil {
			res.EncoderJson(w, editResponse{400, err.Error()})
		} else {
			if err := ac.Do(id, req); err != nil {
				res.EncoderJson(w, editResponse{400, err.Error()})
			} else {
				res.EncoderJson(w, editResponse{
					200,
					fmt.Sprintf("customer edited"),
				})
			}
		}
	default:
		res.EncoderJson(w, editResponse{400, "method not support"})
	}
}

func (ecq *EditRequest) DecoderJson(r http.Request) error {
	return json.NewDecoder(r.Body).Decode(&ecq)
}

func (ecs *editResponse) EncoderJson(w http.ResponseWriter, output editResponse) {
	json.NewEncoder(w).Encode(output)
}
