//go:build memory

package city

import (
	"encoding/json"
	"net/http"
	"tel-note/env"
	"tel-note/lib/convertor"
)

type deleteCity struct{}

var DeleteCity deleteCity

func (dc deleteCity) Do(IDS []uint) []uint {
	return storage.DeleteCityByID(IDS)
}

func (dc *deleteCity) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	switch r.Method {
	case env.DeleteMethod:
		idsCollection := convertor.StrToSlice(r.FormValue("ids"))
		if err, data := convertor.StrSliceToUintSlice(idsCollection); err == nil {
			result := dc.Do(data)
			json.NewEncoder(w).Encode(struct {
				State      uint
				DeletedIDs []uint
			}{200, result})
		} else {
			json.NewEncoder(w).Encode(struct {
				State     uint
				ErrString string
			}{400, err.Error()})
		}
	default:
		json.NewEncoder(w).Encode(struct {
			State   uint
			Message string
		}{400, "method not support"})
	}
}
