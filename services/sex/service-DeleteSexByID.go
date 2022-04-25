package sex

import (
	"encoding/json"
	"net/http"
	"tel-note/env"
	"tel-note/lib/convertor"
)

type deleteSexByID struct{}

var DeleteSexByID deleteSexByID

func (ds *deleteSexByID) Do(ID byte) (err error) {
	err = storage.DeleteSex(ID)
	return err
}

func (ds deleteSexByID) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case env.DeleteMethod:
		if err, sexID := convertor.StrToUint(r.FormValue("sexID")); err == nil {
			if err = ds.Do(byte(sexID)); err == nil {
				json.NewEncoder(w).Encode(struct {
					State   uint
					Message string
				}{200, "success"})
			}
		} else {
			json.NewEncoder(w).Encode(struct {
				State   uint
				Message string
			}{400, err.Error()})
		}
	default:
		json.NewEncoder(w).Encode(struct {
			State   uint
			Message string
		}{400, "method not support"})
	}
}
