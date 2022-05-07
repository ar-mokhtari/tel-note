package sampleData

import (
	"github.com/ar-mokhtari/tel-note/env"
	"net/http"
)

func Init(mux *http.ServeMux) {
	mux.Handle(env.InsertDataR, &FillData)
	mux.Handle(env.GetDataR, &GetData)
}
