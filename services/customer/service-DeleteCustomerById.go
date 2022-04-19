package customer

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tel-note/env"
)

type (
	deleteCustomer struct{}
	DeleteRequest  []uint
	deleteResponse struct {
		State   uint
		Message string
	}
)

var DeleteCustomer deleteCustomer

func (ac *deleteCustomer) Do(ids DeleteRequest) (errs []error) {
	var err error
	for _, data := range ids {
		if err = storage.DeleteCustomerById(data); err != nil {
			errs = append(errs, err)
		}
	}
	return errs
}

func (ac *deleteCustomer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		req DeleteRequest
		res deleteResponse
	)
	switch r.Method {
	case env.DeleteMethod:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if err := req.DecoderJson(*r); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if errs := ac.Do(req); errs != nil {
			for _, err := range errs {
				res.EncoderJson(w, deleteResponse{400, "some values have error\n" + err.Error()})
			}
		} else {
			res.EncoderJson(w, deleteResponse{
				200,
				fmt.Sprintf("customer deleteed"),
			})
		}
	default:
		res.EncoderJson(w, deleteResponse{400, "method not support"})
	}
}

func (ecq *DeleteRequest) DecoderJson(r http.Request) error {
	return json.NewDecoder(r.Body).Decode(&ecq)
}

func (ecs *deleteResponse) EncoderJson(w http.ResponseWriter, output deleteResponse) {
	json.NewEncoder(w).Encode(output)
}
