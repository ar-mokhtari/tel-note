package customer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tel-note/env"
	"tel-note/protocol"
)

type (
	addCustomer struct{}
	NewRequest  protocol.Customer
	newResponse struct {
		State   uint
		Message string
	}
)

var AddCustomer addCustomer

func (ac *addCustomer) Do(newCustomer ...NewRequest) (err error) {
	for _, data := range newCustomer {
		err = storage.NewCustomer(protocol.Customer(data))
	}
	return err
}

func (ac *addCustomer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		req NewRequest
		res newResponse
	)
	switch r.Method {
	case env.PostMethod:
		if err := req.DecoderJson(*r); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := ac.Do(req); err != nil {
			res.EncoderJson(w, newResponse{400, err.Error()})
		} else {
			res.EncoderJson(w, newResponse{
				200,
				fmt.Sprintf("customer added"),
			})
		}
	default:
		res.EncoderJson(w, newResponse{400, "method not support"})
	}
}

func (ecq *NewRequest) DecoderJson(r http.Request) error {
	return json.NewDecoder(r.Body).Decode(&ecq)
}

func (ecs *newResponse) EncoderJson(w http.ResponseWriter, output newResponse) {
	json.NewEncoder(w).Encode(output)
}
