package protocol

import (
	"encoding/json"
	"net/http"
)

type (
	General          struct{}
	EditCityRequest  City
	EditCityResponse struct {
		State   uint
		Message string
	}
	GeneralServices interface {
		DecoderJson(r http.Request) error
		EncoderJson(w http.ResponseWriter, output EditCityResponse)
	}
)

var (
	req EditCityRequest
	res EditCityResponse
)

func (g *General) DecoderJson(r http.Request) error {
	return json.NewDecoder(r.Body).Decode(&req)
}

func (g *General) EncoderJson(w http.ResponseWriter, output EditCityResponse) {
	json.NewEncoder(w).Encode(output)
}
