package job

import (
	"encoding/json"
	"fmt"
	"github.com/ar-mokhtari/tel-note/env"
	"github.com/ar-mokhtari/tel-note/protocol"
	"net/http"
)

type (
	newJob      struct{}
	NewRequest  protocol.Job
	newResponse struct {
		State   uint
		Message string
	}
)

var NewJob newJob

func (nj *newJob) Do(NewJob NewRequest) error {
	if err := storage.NewJob(protocol.Job(NewJob)); err != nil {
		return err
	}
	return nil
}

func (nj *newJob) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		req NewRequest
		res newResponse
	)
	switch r.Method {
	case env.PostMethod:
		if err := req.DecoderJson(*r); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := nj.Do(req); err != nil {
			res.EncoderJson(w, newResponse{400, err.Error()})
		} else {
			res.EncoderJson(w, newResponse{
				200,
				fmt.Sprintf("job added"),
			})
		}

	default:
		res.EncoderJson(w, newResponse{400, "method not support"})
	}
}

func (njq *NewRequest) DecoderJson(r http.Request) error {
	return json.NewDecoder(r.Body).Decode(&njq)
}

func (njs *newResponse) EncoderJson(w http.ResponseWriter, output newResponse) {
	json.NewEncoder(w).Encode(output)
}
