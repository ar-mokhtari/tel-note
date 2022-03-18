package fillSampleData

import "net/http"

func Init() {
	http.Handle("/get-data", http.HandlerFunc(FillDataStruct.ServeGetDataHandle))
	http.Handle("/fill-data", http.HandlerFunc(FillDataStruct.ServeFillData))
}
