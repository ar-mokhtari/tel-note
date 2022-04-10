package sampleData

import (
	"net/http"
	"tel-note/env"
)

func Init() {
	http.Handle(env.InsertSomeSamplesDataR, &FillDataStruct)
	http.Handle(env.PrintAllDataR, &GetDataStruct)
}
