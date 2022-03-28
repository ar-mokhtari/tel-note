//go:build memory

package city

import (
	"encoding/json"
	"net/http"
	"tel-note/lib/convertor"
)

type deleteCityPool struct{}

var DeleteCityPool deleteCityPool

func (allData deleteCityPool) Do(IDS []uint) []uint {
	return storage.DeleteCityByID(IDS)
}

func (allData deleteCityPool) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	idsCollection := convertor.StrToSlice(r.Header.Get("ids"))
	if err, data := convertor.StrSliceToUintSlice(idsCollection); err == nil {
		result := allData.Do(data)
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
}
