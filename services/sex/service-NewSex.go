package sex

import (
	"encoding/json"
	"fmt"
	"net/http"
	"tel-note/env"
	"tel-note/protocol"
)

type (
	addSex      struct{}
	AddRequest  protocol.Sex
	addResponse struct {
		State   uint
		Message string
	}
)

var AddSex addSex

func (ep *addSex) Do(NewSex AddRequest) error {
	if err := storage.NewSex(protocol.Sex(NewSex)); err != nil {
		return err
	}
	return nil
}

func (ep *addSex) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var (
		req AddRequest
		res addResponse
	)
	switch r.Method {
	case env.PostMethod:
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		if err := req.DecoderJson(*r); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if err := ep.Do(req); err != nil {
			res.EncoderJson(w, addResponse{400, err.Error()})
		} else {
			res.EncoderJson(w, addResponse{
				200,
				fmt.Sprintf("sex added"),
			})
		}
	default:
		res.EncoderJson(w, addResponse{400, "method not support"})
	}
}

func (epq *AddRequest) DecoderJson(r http.Request) error {
	return json.NewDecoder(r.Body).Decode(&epq)
}

func (eps *addResponse) EncoderJson(w http.ResponseWriter, output addResponse) {
	json.NewEncoder(w).Encode(output)
}
