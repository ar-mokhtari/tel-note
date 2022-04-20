package job

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tel-note/env"
	"tel-note/lib/convertor"
	"tel-note/protocol"
)

type (
	editJob      struct{}
	EditRequest  protocol.Job
	editResponse struct {
		State   uint
		Message string
	}
)

var EditJob editJob

func (ej *editJob) Do(ID uint, NewJob EditRequest) error {
	if err := storage.EditJob(ID, protocol.Job(NewJob)); err != nil {
		return err
	}
	return nil
}

func (ej *editJob) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		req EditRequest
		res editResponse
	)
	switch r.Method {
	case env.PatchMethod:
		if err := req.DecoderJson(*r); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err, id := convertor.StrToUint(r.FormValue("jobID")); err != nil {
			res.EncoderJson(w, editResponse{400, err.Error()})
		} else {
			if err := ej.Do(id, req); err != nil {
				res.EncoderJson(w, editResponse{400, err.Error()})
			} else {
				res.EncoderJson(w, editResponse{
					200,
					fmt.Sprintf("job edited"),
				})
			}
		}
	default:
		res.EncoderJson(w, editResponse{400, "method not support"})
	}
}

func (ejq *EditRequest) DecoderJson(r http.Request) error {
	return json.NewDecoder(r.Body).Decode(&ejq)
}

func (ejs *editResponse) EncoderJson(w http.ResponseWriter, output editResponse) {
	json.NewEncoder(w).Encode(output)
}
