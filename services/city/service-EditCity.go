//go:build memory

package city

import (
	"encoding/json"
	"fmt"
	"github.com/ar-mokhtari/tel-note/env"
	"net/http"
)

type (
	editCity        struct{}
	EditCityRequest struct {
		city
	}
	editCityResponse struct {
		State   uint
		Message string
	}
)

var (
	EditCity editCity
)

func (ec *editCity) Do(req EditCityRequest) (err error) {
	if err = storage.EditCity(&req); err != nil {
		return err
	}
	return nil
}

func (ec *editCity) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		req EditCityRequest
		res editCityResponse
	)
	switch r.Method {
	case env.PatchMethod:
		if err := req.DecoderJson(*r); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := ec.Do(req); err != nil {
			res.EncoderJson(w, editCityResponse{400, err.Error()})
		} else {
			res.EncoderJson(w, editCityResponse{
				200,
				fmt.Sprintf("City #%v edited", req.id),
			})
		}
	default:
		res.EncoderJson(w, editCityResponse{400, "method not support"})

	}
}

func (ecq *EditCityRequest) DecoderJson(r http.Request) error {
	return json.NewDecoder(r.Body).Decode(&ecq)
}

func (ecs *editCityResponse) EncoderJson(w http.ResponseWriter, output editCityResponse) {
	json.NewEncoder(w).Encode(output)
}
