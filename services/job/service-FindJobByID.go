package job

import (
	"encoding/json"
	"github.com/ar-mokhtari/tel-note/env"
	"github.com/ar-mokhtari/tel-note/lib/convertor"
	"github.com/ar-mokhtari/tel-note/protocol"
	"net/http"
)

type findJobID struct{}

var FindJobID findJobID

func (fji *findJobID) Do(inputID uint) (err error, data protocol.Job) {
	if err, data = storage.FindJobByID(inputID); err == nil {
		return nil, data
	}
	return err, protocol.Job{}
}

func (fji *findJobID) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case env.GetMethod:
		w.Header().Set("Content-Type", "application/json")
		_, uintID := convertor.StrToUint(r.FormValue("jobID"))
		if err, result := fji.Do(uintID); err == nil {
			json.NewEncoder(w).Encode(struct {
				Status  int
				JobData protocol.Job
			}{200, result})
		} else {
			json.NewEncoder(w).Encode(struct {
				Status  int
				JobData string
			}{400, err.Error()})
		}
	default:
		json.NewEncoder(w).Encode(struct {
			State   uint
			Message string
		}{400, "method not support"})
	}
}
