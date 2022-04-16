package general

import (
	"fmt"
	"net/http"
	"tel-note/lib/validator"
)

type checkIranNationalCode struct{}

var CheckIranNational checkIranNationalCode

func (cic *checkIranNationalCode) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//TODO::: get in body
	NationalCode := r.Header.Values("NID")
	fmt.Fprintf(w, "%v", validator.CheckNationalID(NationalCode[0]))
}
