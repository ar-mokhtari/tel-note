package job

import (
	"encoding/json"
	"github.com/ar-mokhtari/tel-note/env"
	"github.com/ar-mokhtari/tel-note/protocol"
	"net/http"
)

type getJob struct{}

var GetJob getJob

func (gj *getJob) Do() []*protocol.Job {
	return storage.GetJobs()
}

func (gj *getJob) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case env.GetMethod:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("Access-Control-Allow-Origin", "*")
		result := gj.Do()
		json.NewEncoder(w).Encode(struct {
			State uint
			Data  []*protocol.Job
		}{200, result})
	default:
		json.NewEncoder(w).Encode(struct {
			State   uint
			Message string
		}{400, "method not support"})

	}
}
