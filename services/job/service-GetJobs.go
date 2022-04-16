package job

import (
	"encoding/json"
	"net/http"
	"tel-note/protocol"
)

type getJob struct{}

var GetJob getJob

func (gj *getJob) Do() []*protocol.Job {
	return storage.GetJobs()
}

func (gj *getJob) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	result := gj.Do()
	json.NewEncoder(w).Encode(struct {
		State uint
		Data  []*protocol.Job
	}{200, result})
}
