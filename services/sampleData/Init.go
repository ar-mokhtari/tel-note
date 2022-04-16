package sampleData

import (
	"net/http"
	"tel-note/env"
)

func Init() {
	http.Handle(env.InsertSomeSamplesDataR, &FillData)
	http.Handle(env.PrintAllDataR, &GetData)
}
