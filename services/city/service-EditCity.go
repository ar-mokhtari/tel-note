//go:build memory

package city

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tel-note/env"
	"tel-note/protocol"
)

var EditCityRequest editCityRequest

type (
	editCityRequest protocol.City
	output          struct {
		State   uint
		Message string
	}
)

func (ec *editCityRequest) Do() (err error) {
	if err := storage.EditCity(protocol.City(*ec)); err != nil {
		return err
	}
	return nil
}
func (ec *editCityRequest) Decoder(r http.Request) error {
	return json.NewDecoder(r.Body).Decode(&ec)
}
func (ec *editCityRequest) Encoder(w http.ResponseWriter, output output) {
	json.NewEncoder(w).Encode(output)
}
func (ec *editCityRequest) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.Method != env.PatchMethod {
		ec.Encoder(w, output{400, "method not support"})
	} else {
		if err := ec.Decoder(*r); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := ec.Do(); err != nil {
			ec.Encoder(w, output{400, err.Error()})
		} else {
			ec.Encoder(w, output{
				200,
				fmt.Sprintf("City #%v edited", ec.Id),
			})
		}
	}
}
