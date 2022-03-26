package general

import (
	"fmt"
	"net/http"
	"tel-note/lib"
)

type checkIranNationalCode struct{}

var CheckIranNational *checkIranNationalCode

func (iranNationalCode *checkIranNationalCode) ServeCheckIranNationalCode(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	NationalCode := r.Header.Values("NID")
	fmt.Fprintf(w, "%v", lib.CheckNationalID(NationalCode[0]))
}
