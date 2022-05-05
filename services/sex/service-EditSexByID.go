package sex

import (
	"encoding/json"
	"fmt"
	"github.com/ar-mokhtari/tel-note/env"
	"github.com/ar-mokhtari/tel-note/protocol"
	"net/http"
)

type (
	editSex      struct{}
	EditRequest  protocol.Sex
	editResponse struct {
		State   uint
		Message string
	}
)

var EditSex editSex

func (ep *editSex) Do(NewSex EditRequest) error {
	if err := storage.EditSex(protocol.Sex(NewSex)); err != nil {
		return err
	}
	return nil
}

func (ep *editSex) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		req EditRequest
		res editResponse
	)
	switch r.Method {
	case env.PatchMethod:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if err := req.DecoderJson(*r); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := ep.Do(req); err != nil {
			res.EncoderJson(w, editResponse{400, err.Error()})
		} else {
			res.EncoderJson(w, editResponse{
				200,
				fmt.Sprintf("sex edited"),
			})
		}
	default:
		res.EncoderJson(w, editResponse{400, "method not support"})
	}
}

func (epq *EditRequest) DecoderJson(r http.Request) error {
	return json.NewDecoder(r.Body).Decode(&epq)
}

func (eps *editResponse) EncoderJson(w http.ResponseWriter, output editResponse) {
	json.NewEncoder(w).Encode(output)
}
