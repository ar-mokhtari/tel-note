package general

import "net/http"

func Init() {
	http.Handle("/check-national-id", http.HandlerFunc(CheckIranNational.ServeCheckIranNationalCode))
}
