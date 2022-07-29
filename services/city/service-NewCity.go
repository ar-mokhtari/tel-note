//go:build memory

package city

import (
	"encoding/json"
	"fmt"
	"github.com/ar-mokhtari/tel-note/env"
	"net/http"
)

type (
	newCity        struct{}
	NewCityRequest struct {
		DTO
	}
	newCityResponse struct {
		State   uint
		Message string
	}
)

var NewCity newCity

func (nc *newCity) Do(req NewCityRequest) (err error) {
	err = storage.NewCity(&req)
	return err
}

func (nc *newCity) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		req NewCityRequest
		res newCityResponse
	)
	switch r.Method {
	case env.PostMethod:
		if err := req.DecoderJson(*r); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := nc.Do(req); err != nil {
			res.EncoderJson(w, newCityResponse{400, err.Error()})
		} else {
			res.EncoderJson(w, newCityResponse{
				200,
				fmt.Sprintf("City added"),
			})
		}
	default:
		res.EncoderJson(w, newCityResponse{400, "method not support"})

	}
}

func (ecq *NewCityRequest) DecoderJson(r http.Request) error {
	return json.NewDecoder(r.Body).Decode(&ecq)
}

func (ecs *newCityResponse) EncoderJson(w http.ResponseWriter, output newCityResponse) {
	json.NewEncoder(w).Encode(output)
}
