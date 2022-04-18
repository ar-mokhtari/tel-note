package country

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tel-note/env"
)

type (
	deleteCountry  struct{}
	DeleteRequest  []uint
	deleteResponse struct {
		State   uint
		Counter uint
		Message string
	}
)

var (
	DeleteCountry deleteCountry
)

func (dc *deleteCountry) Do(IDS []uint) (uint, error) {
	return storage.DeleteCountry(IDS), nil
}

func (dc *deleteCountry) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		req DeleteRequest
		res deleteResponse
	)
	switch r.Method {
	case env.DeleteMethod:
		if err := req.DecoderJson(*r); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if counter, err := dc.Do(req); err != nil {
			res.EncoderJson(w, deleteResponse{400, 0, err.Error()})
		} else {
			res.EncoderJson(w, deleteResponse{
				200,
				counter,
				fmt.Sprintf("%v", "country deleted"),
			})
		}
	default:
		res.EncoderJson(w, deleteResponse{400, 0, "method not support"})
	}
}

func (ecq *DeleteRequest) DecoderJson(r http.Request) error {
	return json.NewDecoder(r.Body).Decode(&ecq)
}

func (ecs *deleteResponse) EncoderJson(w http.ResponseWriter, output deleteResponse) {
	json.NewEncoder(w).Encode(output)
}
