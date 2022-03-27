package fillSampleData

import "net/http"

func Init() {
	http.Handle("/fill-data", &FillDataStruct)
	http.Handle("/get-data", &GetDataStruct)
}
