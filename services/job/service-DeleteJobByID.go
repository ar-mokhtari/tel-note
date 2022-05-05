package job

import (
	"encoding/json"
	"fmt"
	"github.com/ar-mokhtari/tel-note/env"
	"net/http"
)

type (
	deleteJob      struct{}
	DeleteRequest  []uint
	deleteResponse struct {
		State   uint
		Message string
	}
)

var DeleteJob deleteJob

func (dj *deleteJob) Do(ID []uint) []uint {
	return storage.DeleteJob(ID)
}

func (dj *deleteJob) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		req DeleteRequest
		res deleteResponse
	)
	switch r.Method {
	case env.DeleteMethod:
		if err := req.DecoderJson(*r); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		dj.Do(req)
		res.EncoderJson(w, deleteResponse{
			200,
			fmt.Sprintf("job(s) deleted"),
		})

	default:
		res.EncoderJson(w, deleteResponse{400, "method not support"})
	}
}

func (ecq *DeleteRequest) DecoderJson(r http.Request) error {
	return json.NewDecoder(r.Body).Decode(&ecq)
}

func (ecs *deleteResponse) EncoderJson(w http.ResponseWriter, output deleteResponse) {
	json.NewEncoder(w).Encode(output)
}
