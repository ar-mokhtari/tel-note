package general

import (
	"fmt"
	"github.com/ar-mokhtari/tel-note/lib/validator"
	"net/http"
)

type checkIranNationalCode struct{}

var CheckIranNational checkIranNationalCode

func (cic *checkIranNationalCode) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//TODO::: get in body
	NationalCode := r.Header.Values("NID")
	fmt.Fprintf(w, "%v", validator.CheckIsComplex(NationalCode[0]))
}
