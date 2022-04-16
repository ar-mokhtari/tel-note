package sampleData

import (
	"net/http"
	"tel-note/env"
)

func Init() {
	http.Handle(env.InsertDataR, &FillData)
	http.Handle(env.GetDataR, &GetData)
}
