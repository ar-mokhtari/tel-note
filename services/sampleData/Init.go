package sampleData

import (
	"github.com/ar-mokhtari/tel-note/env"
	"net/http"
)

func Init() {
	http.Handle(env.InsertDataR, &FillData)
	http.Handle(env.GetDataR, &GetData)
}
